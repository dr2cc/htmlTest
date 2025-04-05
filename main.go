package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//создаем объект хранилища
	memoryStorage := NewMemoryStorage()
	//создаем обработчик, передавая хранилище в конструктор.
	handler := NewHandler(memoryStorage)
	//Это называется  внедрением зависимостей (Dependency Injection)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
