### About
API for user management. Uses mongoDB for data storage. Application structure:
repository (for interactions with DB) -> service (business logic, at this app stage seems as an overkill but easier to manage dependencies in the future) -> handlers (http, also rpc can be added or other adapters)

### How to run
To start the app run:
```
docker-compose up
```
this will start go app and mongoDB instance

go http server is listening on port :8080

#### header
`Content-Type: application/json`

#### return users list
```
POST localhost:8080/user/findUsers

body {
    "filters": {
        "country": "UK",
    },
    "page": 1
}

```
If multiple filters -> OR query applied 

#### create user
```
POST localhost:8080/user/create

body {
    "first_name": "test1",
    "last_name": "test",
    "nickname": "test",
    "password": "test",
    "email": "test@gmail.com",
    "country": "test"
}
```
Creates new user with new mongoID and sends request to listner that user was created.

#### update user
```
PUT localhost:8080/user/{id}/update

body {
    "first_name": "test1",
    "last_name": "test",
    "nickname": "test",
    "password": "test",
    "email": "test@gmail.com",
    "country": "test"
}
```
Update user by ID given in url. Empty value is going to be set if field is not specified.

#### delete user
```DELETE localhost:8080/user/{id}/delete```

### Run tests
`go test -v ./...`

### Improvements

* password saved withouth encryption :O
* no user request validations
* generic error messages
* more tests and for different cases
* messy way of dealing with User struct (maping in repository etc.)
* documentation (swagger)
* logging implemantation
* no health check
* 