package admin

import (
	"database/sql"
	"net/http"


	"github.com/gin-gonic/gin"
)




type userAndRoleToUpdate struct {
	Email         string `json:"email" form:"email"`
	NewRole       string `json:"role"  form:"role"`
}

type emailToDelete struct {
	Email         string `json:"email" form:"email"`
}


func isRoleValid (role string) (bool) {
	if((role == "ADMIN") || (role == "MEMBER") || (role == "TECHNICIAN")) {
		return true
	}

	return false
}


func isValidAdmin(db *sql.DB,c *gin.Context) (bool) {
	
	id := c.GetUint("id")
	var role string;

	row := db.QueryRow("Select role from users WHERE id = $1",id)
	err := row.Scan(&role)

	if(err != nil &&  err == sql.ErrNoRows ){
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user does not exist"})
		return false
	}

	if(role != "ADMIN") {
		print(role)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user not authorized to access resource"})
		return false
	}

	if(err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return false
	}

	return true
}



func updateRole(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userAndRoleToUpdate userAndRoleToUpdate

		if err := c.Bind(&userAndRoleToUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}

		if (!isRoleValid(userAndRoleToUpdate.NewRole)) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Role"})
			return
		}

		if(userAndRoleToUpdate.Email=="") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}

		if (!isValidAdmin(db,c)) {
			return 
		}

		res, _ := db.Exec("UPDATE users SET role = $1 WHERE email = $2",userAndRoleToUpdate.NewRole,userAndRoleToUpdate.Email)

		count, updateUserErr := res.RowsAffected()



		if count == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"user does not exists"})
			return
		}



		if updateUserErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":updateUserErr.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message":"role update is successful"})
	}
}


func deleteUserInDBWithEmail(db *sql.DB,email string) (sql.Result,error){

	res ,err := db.Exec("DELETE FROM users WHERE (email = $1)",email)
	
	return res,err;
}

func haveDeleteUserErr (c *gin.Context,deleteUserErr error,res sql.Result) (bool) {
	var count int64

	if(deleteUserErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": deleteUserErr.Error()})
		return true	
	}

	count, deleteUserErr = res.RowsAffected()
	if (count == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user does not exist"})
		return true
	}

	if(deleteUserErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"message": deleteUserErr.Error()})
		return true	
	}

	return false
}

func adminDeleteUser(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		
		if (!isValidAdmin(db,c)) {
			return 
		}

		var emailToDelete emailToDelete

		err := c.Bind(&emailToDelete); 
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return
		}

		if (emailToDelete.Email == "") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return;
		}

		res,deleteUserErr := deleteUserInDBWithEmail(db , emailToDelete.Email);
		

		if(haveDeleteUserErr(c,deleteUserErr,res)) {
			return
		}


		c.IndentedJSON(http.StatusOK, gin.H{"message": "the account has been deleted"})
	}

}