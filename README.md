# orderfaz-test-go
This Backend App / Web Service API is using Golang.

## Installation

Use the package manager [golang](https://go.dev/) to install application.
```
go install
```

## List API & Usage
First, run project with :
```bash
go run main.go
```
You can test the application with these routes :
- Register :
```
METHOD : POST
{{baseUrl}}/api/v1/register

# These route need raw json request body:
{
    "msisdn":"Your MSISDN / Phone number",
    "name":"Your Name",
    "username":"Your Username",
    "password":"Your Password"
}
```
NOTE : YOU CAN'T SIGN UP IF MSISDN OR USERNAME HAS ALREADY BEEN REGISTERED / TAKEN.

- Login:
```
METHOD : POST
{{baseUrl}}/api/v1/login

# These route need raw json request body:
{
    "msisdn": "Your registered msisdn",
    "password": "Your registered password"
}
```
NOTE : YOU CAN'T SIGN IN IF YOU NOT SIGN UP YET OR THE PASSWORD IS WRONG.
- Check Auth:
```
METHOD : GET
{{baseUrl}}/api/v1/check-auth

# These route only need Authorization with Type "Bearer Token"
```
NOTE : YOU HAVE TO LOGIN FIRST TO GET TOKEN FOR AUTHORIZATION.

- Add Logistic :
```
METHOD : POST
{{baseUrl}}/api/v1/logistic

# These route need raw json request body. For Example:
{
    "logistic_name":"JNT",
    "amount":15000,
    "destination_name":"BANDUNG",
    "origin_name":"JAKARTA",
    "duration": "1-3"
}
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.

- Get All Logistic :
```
METHOD : GET
{{baseUrl}}/api/v1/logistics
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.
- Get Logistic By Id:
```
METHOD : GET
{{baseUrl}}/api/v1/logistic/{id}
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.
- Get Logistic By Body:
```
METHOD : GET
{{baseUrl}}/api/v1/getlogistic

# These route need raw json request body:
{
    "origin_name":"OriginName",
    "destination_name":"DestinationName"
}
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.
- Edit Logistic:
```
METHOD : PATCH
{{baseUrl}}/api/v1/logistic/{id}

# These route need raw json request body. For example:
{
    "logistic_name":"JNE",
    "amount":20000,
    "destination_name":"JAKARTA",
    "origin_name":"BANDUNG",
    "duration": "2-4"
}
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.
- Delete Logistic:
```
METHOD : DELETE
{{baseUrl}}/api/v1/logistic/{id}
```
NOTE : YOU HAVE TO LOGIN FIRST AND INSERT THE TOKEN FOR USING THIS ENDPOINT.
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Project structure

The project structure of this application's code follows a certain convention. The project is divided into several folders, each with a specific purpose. Here is a tree structure representation of the project folder and its sub-folders:

```
orderfaz-test-go
├── database
│   └── migration.go
│
├── dto
│   ├── auth
│   │   ├──auth_request.go
│   │   └──auth_response.go
│   ├── logistic
│   │   ├──logistic_request.go
│   │   └──logistic_response.go
│   ├── result
│   │   └──result.go
│   └── users
│       ├──user_request.go
│       └──user_response.go
│
├── handlers
│   ├── auth.go
│   ├── logistic.go
│   └── users.go
│
├── models
│   ├── logistic.go
│   └── user.go
│
├── pkg
│   ├── bcrypt
│   │   └──hash_password.go
│   ├── jwt
│   │   └──jwt.go
│   ├── middleware
│   │   └──auth.go
│   └── mysql
│       └──mysql.go
│
├── repositories
│   ├── auth.go
│   ├── logistic.go
│   ├── repository.go
│   └── users.go
├── routes
│   ├── auth.go
│   ├── logistic.go
│   ├── routes.go
│   └── users.go
└──

```


## Swagger Documentation

[Link](https://app.swaggerhub.com/apis-docs/aditya-rizki/orderfaz-test_go/1.0.0#/)