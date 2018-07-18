# go-pg-django

[![Go Report Card](https://goreportcard.com/badge/github.com/tomi77/go-pg-django)](https://goreportcard.com/report/github.com/tomi77/go-pg-django)

Django models for Golang [go-pg](https://github.com/go-pg/pg)

## Installation

~~~sh
go get -u github.com/tomi77/go-pg-django
~~~

## Usage

~~~go
import "github.com/tomi77/go-pg-django/auth"

type Tab struct {
  Id: uint32
  AuthUser *auth.User
}
~~~
