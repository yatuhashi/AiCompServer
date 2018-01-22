# Ai Competion Server

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run 
 
### Build

   revel package AiCompServer

### This is Base System

```
├── README.md
├── app
│   ├── controllers
│   │   └── api
│   │       └── v1
│   │           ├── auth.go
│   │           ├── base.go
│   │           └── user.go
│   ├── db
│   │   └── gorm.go
│   ├── init.go
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   └── routes.go
│   ├── tmp
│   │   └── main.go
│   └── views
├── conf
│   ├── app.conf
│   └── routes
└── tests
    └── apptest.go
```



### Go to http://localhost:9000/ and you'll see:

```
// ApiUser   AiCompServer/app/controllers/api/v1/user.go
GET     /api/v1/user                            ApiUser.Index
GET     /api/v1/user/:id                        ApiUser.Show
POST    /api/v1/user                            ApiUser.Create
PUT     /api/v1/user/:id                        ApiUser.Update
DELETE  /api/v1/user/:id                        ApiUser.Delete


// ApiAuth   AiCompServer/app/controllers/api/v1/auth.go
GET     /api/v1/signin                          ApiAuth.GetSessionID
POST    /api/v1/signin                          ApiAuth.SignIn
GET     /api/v1/signout                         ApiAuth.SignOut
```
