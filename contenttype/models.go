// Package contenttype contains models from `django.contrib.content_type` package
package contenttype

// ContentType represents django_content_type table
type ContentType struct {
	TableName string `sql:"django_content_type"`

	ID       uint16
	AppLabel string `sql:"type:varchar(100),notnull"`
	Model    string `sql:"type:varchar(100),notnull"`
}
