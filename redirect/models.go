package redirect

import (
	"fmt"

	"github.com/tomi77/go-pg-django/site"
)

type Redirect struct {
	TableName string `sql:"django_redirect"`

	Id      uint16
	SiteId  uint16 `pg:",fk:Site" sql:",notnull"`
	Site    *site.Site
	OldPath string `sql:",notnull"`
	NewPath string `sql:",notnull"`
}

func (r Redirect) String() string {
	return fmt.Sprintf("%s --> %s", r.OldPath, r.NewPath)
}
