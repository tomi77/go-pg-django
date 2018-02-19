/*
Models from `django.contrib.admin` package
*/
package admin

import (
	"fmt"
	"time"

	"github.com/tomi77/go-pg-django/auth"
	"github.com/tomi77/go-pg-django/content_type"
)

// Action flags
const (
	ADDITION = 1
	CHANGE   = 2
	DELETION = 3
)

type Log struct {
	TableName string `sql:"django_admin_log"`

	Id            uint16
	ActionTime    time.Time `sql:",notnull"`
	UserId        uint16    `pg:",fk:User" sql:",notnull"`
	User          *auth.User
	ContentTypeId uint16 `pg:",fk:ContentType"`
	ContentType   *content_type.ContentType
	ObjectId      string
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

func (l *Log) IsAddition() bool {
	return l.ActionFlag == ADDITION
}

func (l *Log) IsChange() bool {
	return l.ActionFlag == CHANGE
}

func (l *Log) IsDeletion() bool {
	return l.ActionFlag == DELETION
}
