# interview

API URL:http://newproject-dev.ap-southeast-1.elasticbeanstalk.com


## API Routes:
1. POST /signup
2. POST /login
3. GET /users
4. DELETE /users/delete
5. PUT /users/update
6. PUT /admin/roleUpdate
7. DELETE /admin/delete

## POST /signup
## Arguments

- FirstName        
- LastName
- Password
- Email
- Role : Can only be ADMIN/TECHNICIAN/MEMBER
- Company: (Optional Field)
- Designation (Optional Field)

Example Body:
```
{
	"FirstName":"Chee",        
	"LastName":"Jer En", 
	"Password":"0000",  
	"Email":"special email",
	"Role":"ADMIN",
  "Company":"new company"
}
```





