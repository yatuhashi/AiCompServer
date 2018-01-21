# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run Base

### Go to http://localhost:9000/ and you'll see:

```
// ApiUser   Base/app/controllers/api/v1/user.go
GET     /api/v1/user                            ApiUser.Index
GET     /api/v1/user/:id                        ApiUser.Show
POST    /api/v1/user                            ApiUser.Create
PUT     /api/v1/user/:id                        ApiUser.Update
DELETE  /api/v1/user/:id                        ApiUser.Delete


// ApiAuth   Base/app/controllers/api/v1/auth.go
GET     /api/v1/signin                          ApiAuth.GetSessionID
POST    /api/v1/signin                          ApiAuth.SignIn
GET     /api/v1/signout                         ApiAuth.SignOut
```
