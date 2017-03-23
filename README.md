## i-doit-go-api/goidoit 

is an I-doit API Client written in GO.

**__ This project is still under development and not fully featured __**

Its focused on simplicity to create easy writable code for your projects.

e.g:

```
package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.NewApi("https://example.com/src/jsonrpc.php", "yourapikey")

	// search for sht like test
	s, _ := a.Search("test")

  // thats it
  fmt.Println(s)
}
```

there are more advanced [examples documented](https://github.com/cseeger-epages/i-doit-go-api/tree/master/examples) in the repo

You find the full [code documentation here](https://godoc.org/github.com/cseeger-epages/i-doit-go-api)
