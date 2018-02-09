package auth

import (
	"fmt"
	"time"

	"github.com/tomi77/go-pg-django/content_type"
)

func ExamplePermission_String() {
	contentType := content_type.ContentType{
		Id:       1,
		AppLabel: "auth",
		Model:    "user",
	}
	permission := Permission{
		Id:          1,
		Name:        "Can add user",
		ContentType: &contentType,
		Codename:    "add_user",
	}

	fmt.Println(permission)
	// Output:
	// auth.add_user
}

func ExampleGroup_String() {
	group := Group{
		Id:   1,
		Name: "Admin",
	}

	fmt.Println(group)
	// Output:
	// Admin
}

func ExampleUser_String() {
	user := User{
		Id:          1,
		Username:    "admin",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john@example.com",
		Password:    "password",
		IsStaff:     true,
		IsActive:    true,
		IsSuperuser: true,
		LastLogin:   time.Now(),
		DateJoined:  time.Now(),
	}

	fmt.Println(user)
	// Output:
	// admin
}
