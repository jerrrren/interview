package auth

import (
	"database/sql"

	"fmt"
	"net/http"

	"interviewProject/models"

	
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	
)





type loginCredentials struct {
	Password      string `json:"password" form:"password"`
	Email         string `json:"email" form:"email"`
}

type loginResponse struct {
	Token         string `json:"token"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func verifyPassword(loginPassword string, userPasswordHash string) (error) {

	err := bcrypt.CompareHashAndPassword([]byte(userPasswordHash), []byte(loginPassword))
	

	return err
}


func Signup(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.Bind(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}

		if(newUser.FirstName == ""||newUser.LastName==""||newUser.Password==""||newUser.Email=="") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing field"})
			return 
		}


		if(newUser.Role != "ADMIN"&&newUser.Role != "TECHNICIAN"&&newUser.Role != "MEMBER") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Role"})
			return
		}

		//hash password


		password := HashPassword(newUser.Password)
		newUser.Password = password
		
		result, err := db.Exec("INSERT INTO users ( firstname , lastname , password ,  email , role, company , designation ) VALUES ( $1, $2, $3, $4, $5, $6, $7)",
		 newUser.FirstName , newUser.LastName ,newUser.Password ,  newUser.Email , newUser.Role, newUser.Company , newUser.Designation)


		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error in user fields,maybe there is repeated email"})
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}
}


func findUserInDB(db *sql.DB,foundUser *models.User,loginCredentials loginCredentials) (error){
	row := db.QueryRow("SELECT id,password,email,role FROM users WHERE (email = $1)", loginCredentials.Email)


	err := row.Scan(&foundUser.ID, &foundUser.Password, &foundUser.Email, &foundUser.Role);
	
	return err;
}


func haveFindUserErr (c *gin.Context,findUserErr error) (bool){
	if (findUserErr != nil &&  findUserErr == sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password or email incorrect"})
		return true
	} 

	if(findUserErr != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": findUserErr.Error()})
		return true	
	}

	return false

}

func havePasswordError (checkPasswordErr error, c *gin.Context) (bool) {
	if (checkPasswordErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message":"wrong password"})
		return true
	} 
	return false;
}

func haveGenerateTokenError (generateTokenError error, c *gin.Context) (bool) {

	if generateTokenError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": generateTokenError.Error()})
		return true
	}

	return false
}


func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginCredentials loginCredentials
		var foundUser models.User
		if err := c.BindJSON(&loginCredentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		findUserErr := findUserInDB(db, &foundUser, loginCredentials)

		if(haveFindUserErr(c,findUserErr)) {
			return
		}
		
		checkPasswordErr := verifyPassword(loginCredentials.Password, foundUser.Password)
		
		if(havePasswordError(checkPasswordErr,c)) {
			return 
		}
		
		token,generateTokenError := GenerateAccessToken(foundUser.ID,foundUser.Role)

		if(haveGenerateTokenError(generateTokenError,c)) {
			return
		}
		var loginResponse loginResponse
		loginResponse.Token = token
		
		c.JSON(http.StatusOK, loginResponse)
	}

}



