### How to run
To start the app run:
```
docker-compose up
```
this will start go app and mongoDB instance

go http server is listening on port :8080

in postman or curl 

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

#### create user
```DELETE localhost:8080/user/{id}/delete```

### Run tests
`go test -v ./...`

### Improvements

* more test coverage.
* no endpoint for adding a mineral or removing.
* if action is not performed or fails, there is no message to the user.
* maybe factory shouldn't be coupled with Listener interface, and instead manager is passed to the factory.
* other ways to interact with the app.
* no check how app performs if Factory or Manager is not working.
* no user input validation.