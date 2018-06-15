# basicauth-transport

A simple transport to wrap other transport to add basic auth to request.


## Usage

```go
package main

import (
	"github.com/ArthurHlt/basicauth-transport"
	"net/http"
)

func main() {
	
	http.DefaultClient.Transport = transport.NewDefaultBasicAuthTransport("username", "password")
	
	// with  a custom transport
	http.DefaultClient.Transport = transport.NewBasicAuthTransport(
		"username", 
		"password",
		&http.Transport{})
}
```