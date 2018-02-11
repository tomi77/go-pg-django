/*
Models from `django.contrib.flatpage` package
*/
package flatpage

import (
	"fmt"

	"github.com/tomi77/go-pg-django/site"
)

type Flatpage struct {
	TableName string `sql:"django_flatpage"`

	Id                   uint16
	Url                  string       `sql:"type:varchar(100),notnull"`
	Title                string       `sql:"type:varchar(200),notnull"`
	Content              string       `sql:",notnull"`
	EnableComments       bool         `sql:",notnull,default:false"`
	TemplateName         string       `sql:"type:varchar(70),notnull"`
	RegistrationRequired bool         `sql:",notnull,default:false"`
	Sites                []*site.Site `pg:",many2many:flatpage_sites"`
}

func (fp Flatpage) String() string {
	return fmt.Sprintf("%s -- %s", fp.Url, fp.Title)
}

type FlatpageSite struct {
	TableName string `sql:"django_flatpage_sites"`

	Id         uint16
	FlatpageId uint16 `sql:",notnull"`
	SiteId     uint16 `sql:",notnull"`
}
