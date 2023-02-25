package user

import (
	"database/sql"
	"net/http"

	"interviewProject/auth"
	"interviewProject/models"

	"gopkg.in/guregu/null.v3"

	"github.com/gin-gonic/gin"
)


type userDetails struct {
	FirstName     string `json:"firstName" form:"firstName"`      
	LastName      string `json:"lastName" form:"lastName"`
	Role	      string `json:"role" form:"role"`
	Company       null.String `json:"company" form:"company"`
	Designation   null.String `json:"designation" form:"designation"`
}


type updateRequest struct {
	FirstName     models.JSONString `json:"firstName" form:"firstName"`      
	LastName      models.JSONString `json:"lastName" form:"lastName"`
	Password      models.JSONString `json:"password" form:"passowrd"`  
	Email         models.JSONString `json:"email"  form:"email"`
	Company       models.JSONString `json:"company" form:"company"`
	Designation   models.JSONString `json:"designation" form:"designation"`
}

func findUserInDB(db *sql.DB,foundUser *models.User,id uint) (error){
	row := db.QueryRow("SELECT firstname,lastname,role,company,designation FROM users WHERE (id = $1)", id)

	err := row.Scan(&foundUser.FirstName, &foundUser.LastName, &foundUser.Role,&foundUser.Company,&foundUser.Designation); 
	
	return err;
}

func haveFindUserErr (c *gin.Context,findUserErr error) (bool) {

	if (findUserErr != nil &&  findUserErr == sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user does not exist"})
		return true
	
	}

	if(findUserErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": findUserErr.Error()})
		return true	
	}

	return false

}


func haveDeleteUserErr (c *gin.Context,deleteUserErr error,res sql.Result) (bool) {
	var count int64

	if(deleteUserErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": deleteUserErr.Error()})
		return true	
	}

	count, deleteUserErr = res.RowsAffected()
	if (count == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "you are not a user"})
		return true
	}

	if(deleteUserErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": deleteUserErr.Error()})
		return true	
	}

	return false
}


func deleteUserInDB(db *sql.DB,id uint) (sql.Result,error){

	res ,err := db.Exec("DELETE FROM users WHERE (id = $1)",id)
	

	return res,err;
}


func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetUint("id")

		res,deleteUserErr := deleteUserInDB(db, id);
		
		if(haveDeleteUserErr(c,deleteUserErr,res)) {
			return
		}


		c.IndentedJSON(http.StatusOK, gin.H{"message": "your account has been deleted"})
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var foundUser models.User
		id := c.GetUint("id")

		findUserErr := findUserInDB(db, &foundUser, id);
		
		if(haveFindUserErr(c,findUserErr)) {
			return
		}

		var response userDetails

		response.Company = foundUser.Company
		response.LastName = foundUser.LastName

		response.FirstName = foundUser.FirstName

		response.Designation = foundUser.Designation

		response.Role = foundUser.Role

		c.IndentedJSON(http.StatusOK, response)
	}
}


func isUpdateFieldValid(updatedRequest updateRequest,c *gin.Context) (bool){

		if(updatedRequest.Email.Set && !updatedRequest.Email.Valid) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email is null"})
			return true
		}

		if(updatedRequest.Password.Set && !updatedRequest.Password.Valid) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Password is null"})
			return true
		}

		if(updatedRequest.FirstName.Set && !updatedRequest.FirstName.Valid) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "FirstName is null"})
			return true
		}

		
		if(updatedRequest.LastName.Set && !updatedRequest.LastName.Valid) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "LastName is null"})
			return true
		}

		return false
}

func isUpdateSuccessful(id uint,updatedRequest updateRequest ,c *gin.Context,db *sql.DB) (bool){

		if(updatedRequest.Email.Set && updatedRequest.Email.Valid) {
			_,err :=db.Exec("UPDATE users SET email = $2 WHERE (id = $1)",id,updatedRequest.Email.Value)	

			if(err != nil)	{
				c.JSON(http.StatusBadRequest, gin.H{"message": "email exists"})
				return true

			}
		}

		if(updatedRequest.FirstName.Set && updatedRequest.FirstName.Valid) {
			_,err := db.Exec("UPDATE users SET firstname = $2 WHERE (id = $1)",id,updatedRequest.FirstName.Value)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "firstname failed to update"})
				return true

			}	
		}

		if(updatedRequest.LastName.Set && updatedRequest.LastName.Valid) {
			_,err :=db.Exec("UPDATE users SET lastname = $2 WHERE (id = $1)",id,updatedRequest.LastName.Value)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "lastname failed to update"})
				return true

			}	
		}

		if(updatedRequest.Password.Set && updatedRequest.Password.Valid) {

			updatedPassword := auth.HashPassword(updatedRequest.Password.Value)
			_,err := db.Exec("UPDATE users SET password = $2 WHERE (id = $1)",id,updatedPassword)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "password failed to update"})
				return true

			}	
		}

		//set company
		if(updatedRequest.Company.Set && updatedRequest.Company.Valid) {
			_,err :=db.Exec("UPDATE users SET company = $2 WHERE (id = $1)",id,updatedRequest.Company.Value)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "company failed to update"})
				return true

			}	
		}

		if(updatedRequest.Company.Set && !updatedRequest.Company.Valid) {
			_,err := db.Exec("UPDATE users SET company = $2 WHERE (id = $1)",id,nil)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "company failed to update"})
				return true

			}	
		}

		//set desgination

		if(updatedRequest.Designation.Set && updatedRequest.Designation.Valid) {
			_,err := db.Exec("UPDATE users SET designation = $2 WHERE (id = $1)",id,updatedRequest.Designation.Value)	
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "designation failed to update"})
				return true
			}	
		}

		if(updatedRequest.Designation.Set && !updatedRequest.Designation.Valid) {
			_,err := db.Exec("UPDATE users SET designation = $2 WHERE (id = $1)",id,nil)		
			
			if(err != nil)	{
				c.JSON(http.StatusInternalServerError, gin.H{"message": "designation failed to update"})
				return true
			}

			
		}

	return false
}


func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedRequest updateRequest
		id := c.GetUint("id")

		if err := c.Bind(&updatedRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}

		//check if user in database
		row := db.QueryRow("SELECT id FROM users WHERE id = $1",id)
		var temp uint = 0
		err:=row.Scan(&temp) 

		if(err == sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User does not exists"})
			return 
		}


		if(isUpdateFieldValid(updatedRequest,c)) {
			return
		}

		if(isUpdateSuccessful(id,updatedRequest ,c,db)){
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Updated"})
		 
	}
}