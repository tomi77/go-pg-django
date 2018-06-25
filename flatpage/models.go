// Package flatpage contains models from `django.contrib.flatpage` package
package flatpage

import (
	"fmt"

	"github.com/tomi77/go-pg-django/site"
)

// Flatpage represents django_flatpage table
type Flatpage struct {
	TableName string `sql:"django_flatpage"`

	ID                   uint16
	URL                  string       `sql:"type:varchar(100),notnull"`
	Title                string       `sql:"type:varchar(200),notnull"`
	Content              string       `sql:",notnull"`
	EnableComments       bool         `sql:",notnull,default:false"`
	TemplateName         string       `sql:"type:varchar(70),notnull"`
	RegistrationRequired bool         `sql:",notnull,default:false"`
	Sites                []*site.Site `pg:",many2many:flatpage_sites"`
}

func (fp Flatpage) String() string {
	return fmt.Sprintf("%s -- %s", fp.URL, fp.Title)
}

// Site represents django_flatpage_sites table
type Site struct {
	TableName string `sql:"django_flatpage_sites"`

	ID         uint16
	FlatpageID uint16 `sql:",notnull"`
	SiteID     uint16 `sql:",notnull"`
}
