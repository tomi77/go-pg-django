package redirect

import (
  "fmt"

  "github.com/tomi77/go-pg-django/site"
)

func ExampleSite_String() {
  site := site.Site{
    Id: 1,
    Domain: "www.example.com",
    Name: "Main page",
  }
  redirect := Redirect{
    Id: 1,
    Site: &site,
    OldPath: "/old",
    NewPath: "/new",
  }

  fmt.Println(redirect)
  // Output:
  // /old --> /new
}
