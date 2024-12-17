package main

import (
	"net/http"
	"github.com/Romashka641/akademServer/routes"
)

func main()  {
	routes()
	http.ListenAndServe(":8081", nil)
}