// shortoftime main binary
package main

import (
	"github.com/CerealBoy/shortoftime/time"
	"github.com/CerealBoy/shortoftime/web"
)

func main() {
	t := time.New()
	w := web.New(t)

	w.Serve(":8080")
}
