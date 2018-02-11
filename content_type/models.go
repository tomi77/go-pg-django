/*
Models from `django.contrib.content_type` package
*/
package content_type

type ContentType struct {
	TableName string `sql:"django_content_type"`

	Id       uint16
	AppLabel string `sql:",notnull"`
	Model    string `sql:",notnull"`
}
