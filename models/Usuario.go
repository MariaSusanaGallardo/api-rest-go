package models

import "gorm.io/gorm"

//creamos la estructura para Usuarios

type Usuario struct {
	gorm.Model //convierte la estructura en una tabla

	Nombre   string `gorm:"not null" json:"nombre"`
	Apellido string `gorm:"not null" json:"apellido"`
	Email    string `gorm:"not null;unique_index json:email"`
	Tareas   []Tarea
}
