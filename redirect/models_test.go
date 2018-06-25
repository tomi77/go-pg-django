package redirect

import (
	"fmt"

	"github.com/tomi77/go-pg-django/site"
)

func ExampleRedirect_String() {
	site := site.Site{
		ID:     1,
		Domain: "www.example.com",
		Name:   "Main page",
	}
	redirect := Redirect{
		ID:      1,
		Site:    &site,
		OldPath: "/old",
		NewPath: "/new",
	}

	fmt.Println(redirect)
	// Output:
	// /old --> /new
}
