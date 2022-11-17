package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=admin password=admin dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{}) //se intatará conectar, puede devolver un error

	//manejamos el error
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

}
