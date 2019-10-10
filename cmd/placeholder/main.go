package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/wzshiming/placeholder/route"
)

var port = 8080

func main() {
	mux := route.Router()
	mux = handlers.RecoveryHandler()(mux)
	mux = handlers.CombinedLoggingHandler(os.Stdout, mux)
	p := fmt.Sprintf(":%d", port)

	fmt.Printf("Open http://127.0.0.1:%d/swagger/ with your browser.\n", port)
	err := http.ListenAndServe(p, mux)
	if err != nil {
		fmt.Println(err)
	}
	return
}
