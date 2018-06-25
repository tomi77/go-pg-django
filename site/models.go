// Package site contains models from `django.contrib.site` package
package site

// Site represents django_site table
type Site struct {
	TableName string `sql:"django_site"`

	ID     uint16
	Domain string `sql:"type:varchar(100),notnull"`
	Name   string `sql:"type:varchar(50),notnull"`
}

func (s Site) String() string {
	return s.Domain
}
