/*
Models from `django.contrib.site` package
*/
package site

type Site struct {
	TableName string `sql:"django_site"`

	Id     uint16
	Domain string `sql:",notnull"`
	Name   string `sql:",notnull"`
}

func (s Site) String() string {
	return s.Domain
}
