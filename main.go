package main

import (
	"interviewProject/admin"
	"interviewProject/auth"
	"interviewProject/db"
	"interviewProject/routerMiddleware"
	"interviewProject/user"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func main() {

	println("testing")
	var DB *sql.DB = db.SetupDatabase()
	router := gin.Default()
	router.Use(routerMiddleware.CORSMiddleware())

	auth.AuthRoutes(router,DB)
	user.UserRoutes(router,DB)
	admin.AdminRoutes(router,DB)

	router.Run(":8081")

	defer DB.Close()
}






