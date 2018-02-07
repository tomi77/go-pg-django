package site

import (
  "fmt"
)

func ExampleSite_String() {
  site := Site{
    Id: 1,
    Domain: "www.example.com",
    Name: "Main page",
  }

  fmt.Println(site)
  // Output:
  // www.example.com
}
