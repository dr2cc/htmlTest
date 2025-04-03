package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//mux := http.NewServeMux()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hi Go!")
	// })
	// http.HandleFunc("/echo/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, r.URL.Query().Get("message"))
	// })
	// http.ListenAndServe("localhost:8080", nil)
	router := gin.Default()
	router.POST("/employee")
	router.GET(" /employee/:id")
	router.PUT(" /employee/:id")
	router.DELETE(" /employee/:id")

	router.Run()
}
