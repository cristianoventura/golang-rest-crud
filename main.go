package main

import (
	"net/http"

	Controllers "./controllers"
	Database "./database"
	// "./services"
	// "./structs"
)

func main() {
	db := Database.Connect()

	// person := structs.Person{
	// 	"Test",
	// 	"Lastname",
	// 	25,
	// }
	// services.AddPerson(person, db)

	http.ListenAndServe(":8080", Controllers.PersonController(db))
}
