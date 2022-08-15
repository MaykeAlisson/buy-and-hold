package main

import (
	"fmt"

	"github.com/maykealisson/buy-and-hold/routes"
)

func main() {
	// database.ConectaComBancoDeDados()
	fmt.Println("run server port 3000")
	routes.HandlerRequest()
}
