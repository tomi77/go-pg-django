// Package session contains models from `django.contrib.session` package
package session

import (
	"time"
)

// Session represents django_session table
type Session struct {
	TableName string `sql:"django_session"`

	SessionKey  string    `sql:"type:varchar(40),notnull"`
	SessionData string    `sql:",notnull"`
	ExpireDate  time.Time `sql:",notnull"`
}

func (s Session) String() string {
	return s.SessionKey
}
