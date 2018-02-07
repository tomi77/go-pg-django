package auth

import (
  "fmt"
  "time"

  "github.com/tomi77/go-pg-django/content_type"
)

type Permission struct {
  TableName     string       `sql:"auth_permission"`

  Id            uint16
  Name          string       `sql:"type:varchar(255),notnull"`
  ContentTypeId uint16       `pg:",fk:ContentType" sql:",notnull"`
  ContentType   *content_type.ContentType
  Codename      string       `sql:"type:varchar(100),notnull"`
}

func (p Permission) String() string {
  return fmt.Sprintf("%s.%s", p.ContentType.AppLabel, p.Codename)
}

type Group struct {
  TableName   string        `sql:"auth_group"`

  Id          uint16
  Name        string        `sql:"type:varchar(80),notnull"` // unique
  Permissions []*Permission `pg:",many2many:group_permissions"`
}

func (g Group) String() string {
  return g.Name
}

type GroupPermission struct {
  TableName    string `sql:"auth_group_permissions"`

  Id           uint16
  GroupId      uint16
  PermissionId uint16
}

type User struct {
  TableName   string        `sql:"auth_user"`

  Id          uint16
  Username    string        `sql:"type:varchar(150),notnull"`
  FirstName   string        `sql:"type:varchar(30),notnull"`
  LastName    string        `sql:"type:varchar(30),notnull"`
  Email       string        `sql:"varchar(254),notnull"`
  Password    string        `sql:"varchar(128),notnull"`
  IsStaff     bool          `sql:",notnull"`
  IsActive    bool          `sql:",notnull"`
  IsSuperuser bool          `sql:",notnull"`
  LastLogin   time.Time
  DateJoined  time.Time     `sql:",notnull"`
  Groups      []*Group      `pg:",many2many:user_groups"`
  Permissions []*Permission `pg:",many2many:user_permissions"`
}

func (u User) String() string {
  return u.Username
}

type UserGroup struct {
  TableName string `sql:"auth_user_groups"`

  Id        uint16
  UserId    uint16
  GroupId   uint16
}

type UserPermission struct {
  TableName    string `sql:"auth_user_user_permissions"`

  Id           uint16
  UserId       uint16
  PermissionId uint16
}
