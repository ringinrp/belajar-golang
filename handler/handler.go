package handler

import (
	"belajar-golang/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello semua"))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "im learing golang",
	// 	"content": "iam still learning",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 3},
		{ID: 2, Name: "Xenia", Price: 24000000, Stock: 8},
		{ID: 3, Name: "Avanza", Price: 15000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello ini about"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Page ini tidak bisa ditampilkan", http.StatusBadRequest)

	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
			return
		}
		return

	}
	http.Error(w, "Page tidak bisa ditampilkan", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Page tidak bisa ditampilkan", http.StatusBadRequest)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Page tidak bisa ditampilkan", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Page tidak bisa ditampilkan", http.StatusBadRequest)

}
