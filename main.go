package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//создаем объект хранилища
	memoryStorage := NewMemoryStorage()
	//создаем обработчик, передавая хранилище в конструктор.
	//в оригинале- handler
	//Но это просто Жашкевичу, а не мне..
	//Переименовал, что-бы не путаться кто и на ком стоял
	handlerLocal := NewHandler(memoryStorage)
	//Это называется  внедрением зависимостей (Dependency Injection)

	//здесь создаем объект типа *gin.Engine, "машина",
	//по сути- маршрутизатор, тут- роутер
	router := gin.Default()

	router.POST("/employee", handlerLocal.CreateEmployee)
	router.GET("/employee/:id", handlerLocal.GetEmployee)
	router.PUT("/employee/:id", handlerLocal.UpdateEmployee)
	router.DELETE("/employee/:id", handlerLocal.DeleteEmployee)

	router.Run()
}
