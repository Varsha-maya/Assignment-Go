package main

import (
	"awesomeProject3/model"
	db "awesomeProject3/repository"
	. "awesomeProject3/router"
)

func main() {
	dbHost := "127.0.0.1:27017"
	db.Init(&model.Database{
		Driver:   "mongodb",
		Endpoint: dbHost})
	defer db.Exit()

	Router()

}