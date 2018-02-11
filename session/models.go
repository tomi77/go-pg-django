/*
Models from `django.contrib.session` package
*/
package session

import (
	"time"
)

type Session struct {
	TableName string `sql:"django_session"`

	Id          uint16
	SessionKey  string    `sql:",notnull"`
	SessionData string    `sql:",notnull"`
	ExpireDate  time.Time `sql:",notnull"`
}

func (s Session) String() string {
	return s.SessionKey
}
