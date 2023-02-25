package models
import(
	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID            uint   `json:"uid" form:"uid"` 
	FirstName     string `json:"firstName" form:"firstName"`      
	LastName      string `json:"lastName" form:"lastName"`
	Password      string `json:"password" form:"password"`
	Email         string `json:"email" form:"email"`
	Role	      string `json:"role" form:"role"`
	Company       null.String `json:"company" form:"company"`
	Designation   null.String `json:"designation" form:"designation"`
}