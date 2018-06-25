// Package admin contains models from `django.contrib.admin` package
package admin

import (
	"fmt"
	"time"

	"github.com/tomi77/go-pg-django/auth"
	"github.com/tomi77/go-pg-django/contenttype"
)

// Action flags
const (
	ADDITION = 1
	CHANGE   = 2
	DELETION = 3
)

// Log represents django_admin_log table
type Log struct {
	TableName string `sql:"django_admin_log"`

	ID            uint16
	ActionTime    time.Time `sql:",notnull"`
	UserID        uint16    `pg:",fk:User" sql:",notnull"`
	User          *auth.User
	ContentTypeID uint16 `pg:",fk:ContentType"`
	ContentType   *contenttype.ContentType
	ObjectID      string
	ObjectRepr    string `sql:"type:varchar(200),notnull"`
	ActionFlag    uint16 `sql:",notnull"`
	ChangeMessage string
}

func (l Log) String() string {
	switch l.ActionFlag {
	case ADDITION:
		return fmt.Sprintf("Added %s.", l.ObjectRepr)
	case CHANGE:
		return fmt.Sprintf("Changed %s - %s", l.ObjectRepr, l.ChangeMessage)
	case DELETION:
		return fmt.Sprintf("Deleted %s.", l.ObjectRepr)
	default:
		return "LogEntry Object"
	}
}

// IsAddition returns true if action is "add"
func (l *Log) IsAddition() bool {
	return l.ActionFlag == ADDITION
}

// IsChange returns true if action is "change"
func (l *Log) IsChange() bool {
	return l.ActionFlag == CHANGE
}

// IsDeletion returns true if action is "delete"
func (l *Log) IsDeletion() bool {
	return l.ActionFlag == DELETION
}
