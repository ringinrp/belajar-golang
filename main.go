package main

import (
	"belajar-golang/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//untuk routing
	route := mux.NewRouter()

	route.HandleFunc("/", handler.HomeHandler).Methods("GET")
	route.HandleFunc("/about", handler.AboutHandler).Methods("GET")
	route.HandleFunc("/product", handler.ProductHandler).Methods("GET")
	route.HandleFunc("/post-get", handler.PostGet)
	route.HandleFunc("/form", handler.Form)
	route.HandleFunc("/process", handler.Process)

	//untuk manggil file css
	//huruf depan folder harus menggunakan huruf besar
	route.PathPrefix("/Assets/").Handler(http.StripPrefix("/Assets/", http.FileServer(http.Dir("./Assets"))))

	//untuk menjalankan server
	port := "8000"

	fmt.Println("Server sedang berjalan di port " + port)
	http.ListenAndServe("localhost:"+port, route)

}
