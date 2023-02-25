package auth

import (
	"database/sql"


	"github.com/gin-gonic/gin"
)




func AuthRoutes(incomingRoutes *gin.Engine,db *sql.DB){
	incomingRoutes.POST("/signup",Signup(db))
	incomingRoutes.POST("/login",Login(db))
}