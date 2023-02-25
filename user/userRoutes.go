package user

import (
	"database/sql"

	"interviewProject/routerMiddleware"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine,db *sql.DB){
	incomingRoutes.GET("/users",routerMiddleware.Authenticate(),GetUser(db))
	incomingRoutes.DELETE("/users/delete",routerMiddleware.Authenticate(),DeleteUser(db))
	incomingRoutes.PUT("/users/update",routerMiddleware.Authenticate(),UpdateUser(db))
}				