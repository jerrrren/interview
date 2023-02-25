# interview

API URL:http://newproject-dev.ap-southeast-1.elasticbeanstalk.com


## API Routes:
(Non-Authenticated Endpoints)
1. POST /signup
2. POST /login
(Authenticated Endpoints): All authenticated endpoints require JWT token stored in authorization bearer header to be accessed
4. GET /users
5. DELETE /users/delete
6. PUT /users/update
7. PUT /admin/roleUpdate
8. DELETE /admin/delete

## POST /signup
Create new user

### Arguments

- FirstName        
- LastName
- Password
- Email
- Role : Can only be ADMIN/TECHNICIAN/MEMBER
- Company: (Optional Field)
- Designation (Optional Field)

Example Request:
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

## POST /login

Gets JWT token for endpoint authentication

### Arguments

- Email
- Password

Example Request:
```
{
  "password":"0000",
  "email":"special email"
}
```

Example Response:
```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6OSwiZXhwIjoxNjc3MjkwMTI0fQ.QrCG_duq1QIWTVJ5R13ERnlmxLLzvQZ1F3KVCfjTx7E"
}
```


## GET /users (Authenticated)
There is no argument required for this endpoint, just ensure to include the jwt token in the authorization bearer header for this endpoint to work

Gets user information

Example Response
```
{
  "firstName": "Chee",
  "lastName": "Jer En",
  "role": "ADMIN",
  "company": "new company",
  "designation": null
}
```

## GET /users/delete (Authenticated)
Deletes user
There is no argument required for this endpoint, just ensure to include the jwt token in the authorization bearer header for this endpoint to work

```
{
  "message": "your account has been deleted"
}
```

## PUT /users/update (Authenticated)
Updates user Details

### Arguments
- FirstName (Optional Field)       
- LastName (Optional Field)
- Password (Optional Field)
- Email (Optional Field)
- Company: (Optional Field)
- Designation (Optional Field)

Example Request:
```
{
	"FirstName":"Chee",        
	"LastName":"Jer En", 
	"Password":"0000",  
	"Email":"new email",
  	"Company":"new company",
  	"Designation":"Mr"
}

```

Example Response
```
{
  "message": "Updated"
}
```

(Authenticated)
## PUT /admin/roleUpdate
Updates user roles
### Arguments
- Email       
- Role 


```
{
  "role":"MEMBER",
  "email":"changed email"`
}
```

Example Response
```
{
  "message": "role update is successful"
}
```


## Delete /admin/delete (Authenticated)
Deletes any specific user
### Arguments
- Email       


```
{
  "email":"changed email"`
}
```

Example Response
```
{
  "message": "the account has been deleted"
}
```
 










