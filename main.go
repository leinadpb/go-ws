package main

import (
	"net/http"

	"github.com/leinadpb/go-ws/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
