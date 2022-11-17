package models

import "gorm.io/gorm"

type Tarea struct {
	gorm.Model

	Titulo      string `gorm:"not null;unique-index" json:"titulo"`
	Descripcion string `json:"descripcion"`
	Terminada   bool   `gorm:"default:false" json:"terminada"`
	UsuarioID   uint   `json:"usuarioId"`
}
