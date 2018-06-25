package auth

import (
	"fmt"
	"time"

	"github.com/tomi77/go-pg-django/contenttype"
)

func ExamplePermission_String() {
	contentType := contenttype.ContentType{
		ID:       1,
		AppLabel: "auth",
		Model:    "user",
	}
	permission := Permission{
		ID:          1,
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
		ID:   1,
		Name: "Admin",
	}

	fmt.Println(group)
	// Output:
	// Admin
}

func ExampleUser_String() {
	user := User{
		ID:          1,
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
