package flatpage

import (
	"fmt"
)

func ExampleFlatpage_String() {
	flatpage := Flatpage{
		ID:                   1,
		URL:                  "https://www.example.com/",
		Title:                "Main page",
		Content:              "Starting page",
		EnableComments:       false,
		TemplateName:         "",
		RegistrationRequired: false,
	}

	fmt.Println(flatpage)
	// Output:
	// https://www.example.com/ -- Main page
}
