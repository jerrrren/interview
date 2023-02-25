package admin

import (
	"database/sql"

	"interviewProject/routerMiddleware"

	"github.com/gin-gonic/gin"
)


func AdminRoutes(incomingRoutes *gin.Engine,db *sql.DB){
	incomingRoutes.PUT("/admin/roleUpdate",routerMiddleware.Authenticate(),updateRole(db))
	incomingRoutes.DELETE("/admin/delete",routerMiddleware.Authenticate(),adminDeleteUser(db))
}				