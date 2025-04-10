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

	//POST запрос по адресу localhost:8080/employee
	//в теле запроса данные json о "сотруднике"
	//{
	// 	"name": "Pete",
	// 	"age": 25,
	// 	"salary": 2000,
	// 	"sex": "M"
	// }
	router.POST("/employee", handlerLocal.CreateEmployee)
	//в ответ получаем id
	// {
	// 	"id": 1
	// }

	//Теперь GET по тому же адресу, но с требуемым id,
	//получаем информацию о сотруднике
	router.GET("/employee/:id", handlerLocal.GetEmployee)

	//PUT не понял как работает
	router.PUT("/employee/:id", handlerLocal.UpdateEmployee)

	//DELETE удаляет сотрудника с введенным id
	router.DELETE("/employee/:id", handlerLocal.DeleteEmployee)

	router.Run()
}
