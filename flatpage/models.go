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
	Url                  string `sql:",notnull"`
	Title                string `sql:",notnull"`
	Content              string `sql:",notnull"`
	EnableComments       bool   `sql:",notnull"` // default false
	TemplateName         string `sql:",notnull"`
	RegistrationRequired bool   `sql:",notnull"` // default false
	Sites                []*site.Site
}

func (fp Flatpage) String() string {
	return fmt.Sprintf("%s -- %s", fp.Url, fp.Title)
}
