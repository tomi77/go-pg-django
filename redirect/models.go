// Package redirect contains models from `django.contrib.redirect` package
package redirect

import (
	"fmt"

	"github.com/tomi77/go-pg-django/site"
)

// Redirect represents django_redirect table
type Redirect struct {
	TableName string `sql:"django_redirect"`

	ID      uint16
	SiteID  uint16 `pg:",fk:Site" sql:",notnull"`
	Site    *site.Site
	OldPath string `sql:"type:varchar(200),notnull"`
	NewPath string `sql:"type:varchar(200),notnull"`
}

func (r Redirect) String() string {
	return fmt.Sprintf("%s --> %s", r.OldPath, r.NewPath)
}
