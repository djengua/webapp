package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
	// files := []string{
	// 	"templates/layout.html",
	// 	"templates/navbar.html",
	// }
	// templates := template.Must(template.ParseFiles(files...))

	// // threads, err := data.Threads()
	// // if err == nil {
	// // 	templates.ExecuteTemplate(w, "layout", threads)
	// // }

	// threads := []models.Thread{
	// 	{
	// 		Id:        1,
	// 		Uuid:      "4fa6b458-12dd-4b9f-b7c3-48a720b3df06",
	// 		Topic:     "Hello first topic",
	// 		UserId:    1,
	// 		CreatedAt: time.Now(),
	// 	},
	// 	{
	// 		Id:        2,
	// 		Uuid:      "b4ff633d-84e0-4899-b3b7-b6a392ed442e",
	// 		Topic:     "Hello second topic",
	// 		UserId:    1,
	// 		CreatedAt: time.Now(),
	// 	},
	// }
	// templates.ExecuteTemplate(w, "layout", threads)
}

func main() {
	mux := http.NewServeMux()
	// files := http.FileServer(http.Dir("/public"))
	// mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println("Server in port 8080")
	server.ListenAndServe()
	// Serving through HTTPS
	// serve.ListenAndServeTLS("cert.pem", "key.pem")

}
