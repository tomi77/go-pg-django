/*
Models from `django.contrib.site` package
*/
package site

type Site struct {
	TableName string `sql:"django_site"`

	Id     uint16
	Domain string `sql:"type:varchar(100),notnull"`
	Name   string `sql:"type:varchar(50),notnull"`
}

func (s Site) String() string {
	return s.Domain
}
