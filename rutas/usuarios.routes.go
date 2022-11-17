package rutas

import (
	"encoding/json"
	"net/http"

	"github.com/MariaSusanaGallardo/api-rest/db"
	"github.com/MariaSusanaGallardo/api-rest/models"
	"github.com/gorilla/mux"
)

// funcion para obtener todos los usuarios
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {

	//hacemos una variable slice cuyo tipo es Usuario (desde nuestro modelo)
	var usuarios []models.Usuario

	//usamos la función Find de gorm para buscar la estructura desarrollada en Model
	db.DB.Find(&usuarios)

	//package json => newEncoder codifica en json lo que se escribe
	json.NewEncoder(w).Encode(&usuarios)

}

// funcion para ver usuario por id
func VerUsuarioId(w http.ResponseWriter, r *http.Request) {

	//creamos un objeto del tipo Usuario para llevarlo a db.DB.First(&usuario, params["id"])
	var usuario models.Usuario

	//con el package mux y su función Vars accedemos a las variables da la response, r
	params := mux.Vars(r)

	//de esta forma accedemos a la tabla que tenga el id que ponemos y con FIrst le pedimos el primer elemento que coincida
	db.DB.First(&usuario, params["id"])

	//validamos que exista el usuario, si es 0 no existe

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
		return

	}

	//relacionamos las tareas del usuario con su id
	db.DB.Model(&usuario).Association("Tareas").Find(&usuario.Tareas)

	json.NewEncoder(w).Encode(&usuario)

}

// función para crear un usuario
func CrearUsuario(w http.ResponseWriter, r *http.Request) {

	//declaramos una variable usuario
	var usuario models.Usuario

	//importamos json y asignamos a usuario el contenido del body en json que establecimos
	json.NewDecoder(r.Body).Decode(&usuario)

	//guardamos los datos en la base de datos
	usuarioCreado := db.DB.Create(&usuario)

	err := usuarioCreado.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&usuario)

}

// función para borar usuario
func BorrarUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario
	params := mux.Vars(r)

	db.DB.First(&usuario, params["id"])

	//vbalidamos que exisa el user

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
		return //para que se cierre y muestre el mensaje si no encuetra nada
	}

	//de esta forma se elimina pero se mantiene en la bse de datos comno registro
	// db.DB.Delete(&usuario)

	//de esta forma se elimina definitivametne de la tabla
	db.DB.Unscoped().Delete(&usuario)

	w.WriteHeader(http.StatusOK)

}
