# Jopher

Easy JSON responses for Go on top of net/http

#### Install

```bash
go get github.com/emostafa/jopher
```


#### Usage

import jopher to your project

```go
import (
    ...
    "jopher"
)
```

There is no need to initialize any thing, the usage is very simple,
inside if your handlers use jopher to send a json response

e.g:

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    name := "Foo"

    jopher.Success(w, name)
}
```


#### Methods

The main function Jopher uses to send responses is Write(), you can use it for all of your responses,
but alongside it, Jopher provides two other types of methods, error responses, and success responses.


###### Write

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    name := "Foo"

    jopher.Write(w, http.StatusOK, name)
    // or
    jopher.Write(w, 200, name)
}
```

###### Success

Success() returns a 200 response with the supplied message as the body of the response

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    name := "Foo"

    jopher.Success(w, name)
}
```

###### Created

Success() returns a 201 response with the supplied message as the body of the response

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type post struct {
        Title string
        Body string
    }

    // e.g: save post to db: 
    db.Insert(&Post{"foo title", "bar body"})

    jopher.Created(w, "Successfully created the new post")
}
```

###### Error

Error() returns an error response with the supplied message as the body of the response
and the a status code.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type post struct {
        Title string
        Body string
    }

    // e.g: save post to db: 
    err := db.Insert(&Post{"foo title", "bar body"})
    if err != nil {
        jopher.Error(w, 500, err)
    }
    ...
}
```

###### BadRequest

BadRequest() uses Error() to returns an error response with the supplied message as the body of the response
and a status code of 400.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type user struct {
        Email string
        Password string
    }

    ...
    // e.g: validate request body parameters
    if req_body.Get("email") == nil {
        jopher.BadRequest(w, errors.New("Email field is required"))
    }
}
```

###### Unauthorized

Unauthorized() uses Error() to returns an error response with the supplied message as the body of the response
and a status code of 401.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type User struct {
        Email string
        Password string
    }

    // e.g: save post to db: 
    db.Insert(&Post{"foo title", "bar body"})
    if user == nil {
        jopher.Unauthorized(w, errors.New("You dont have permission to create a new post"))
    }
    ...
}
```

###### NotFound

NotFound() uses Error() to returns an error response with the supplied message as the body of the response
and a status code of 404.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type user struct {
        Email string
        Password string
    }

    // e.g: trying to find user in db
    c, _ := db.Count("users", &user{"foo@bar.com"})
    if c < 1 {
        jopher.NotFound(w, errors.New("User was not found"))
    }
    ...
}
```

###### InternalServerError

InternalServerError() uses Error() to returns an error response with the supplied message as the body of the response
and a status code of 500.

```go
func fooHandler(w http.ResponseWriter, r *http.Request) {
    type post struct {
        Title string
        Body string
    }

    // e.g: save post to db: 
    err := db.Insert(&Post{"foo title", "bar body"})
    if err != nil {
        jopher.InternalServerError(w, err)
    }
    ...
}
```