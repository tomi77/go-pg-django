// Package auth contains models from `django.contrib.auth` package
package auth

import (
	"fmt"
	"time"

	"github.com/tomi77/go-pg-django/contenttype"
)

// Permission represents auth_permission table
type Permission struct {
	TableName string `sql:"auth_permission"`

	ID            uint16
	Name          string `sql:"type:varchar(255),notnull"`
	ContentTypeID uint16 `pg:",fk:ContentType" sql:",notnull"`
	ContentType   *contenttype.ContentType
	Codename      string `sql:"type:varchar(100),notnull"`
}

func (p Permission) String() string {
	return fmt.Sprintf("%s.%s", p.ContentType.AppLabel, p.Codename)
}

// Group represents auth_group table
type Group struct {
	TableName string `sql:"auth_group"`

	ID          uint16
	Name        string        `sql:"type:varchar(80),notnull"` // unique
	Permissions []*Permission `pg:",many2many:group_permissions"`
}

func (g Group) String() string {
	return g.Name
}

// GroupPermission represents auth_group_permissions table
type GroupPermission struct {
	TableName string `sql:"auth_group_permissions"`

	ID           uint16
	GroupID      uint16 `sql:",notnull"`
	PermissionID uint16 `sql:",notnull"`
}

// User represents auth_user table
type User struct {
	TableName string `sql:"auth_user"`

	ID          uint16
	Username    string `sql:"type:varchar(150),notnull"`
	FirstName   string `sql:"type:varchar(30),notnull"`
	LastName    string `sql:"type:varchar(30),notnull"`
	Email       string `sql:"type:varchar(254),notnull"`
	Password    string `sql:"type:varchar(128),notnull"`
	IsStaff     bool   `sql:",notnull"`
	IsActive    bool   `sql:",notnull"`
	IsSuperuser bool   `sql:",notnull"`
	LastLogin   time.Time
	DateJoined  time.Time     `sql:",notnull"`
	Groups      []*Group      `pg:",many2many:user_groups"`
	Permissions []*Permission `pg:",many2many:user_permissions"`
}

func (u User) String() string {
	return u.Username
}

// UserGroup represents auth_user_groups table
type UserGroup struct {
	TableName string `sql:"auth_user_groups"`

	ID      uint16
	UserID  uint16 `sql:",notnull"`
	GroupID uint16 `sql:",notnull"`
}

// UserPermission represents auth_user_user_permissions table
type UserPermission struct {
	TableName string `sql:"auth_user_user_permissions"`

	ID           uint16
	UserID       uint16 `sql:",notnull"`
	PermissionID uint16 `sql:",notnull"`
}
