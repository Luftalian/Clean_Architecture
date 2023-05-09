package main

import (
	"github.com/Luftalian/Clean_Architecture/infrastructure"
)

func main() {
	infrastructure.Router.Logger.Fatal(infrastructure.Router.Start(":8080"))
}
