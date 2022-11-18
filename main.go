package main

import (
	"net/http"

	"github.com/MariaSusanaGallardo/api-rest/db"
	"github.com/MariaSusanaGallardo/api-rest/models"
	"github.com/MariaSusanaGallardo/api-rest/rutas"
	"github.com/gorilla/mux"
)

//FUNCIÓN PRINCIPAL

func main() {

	//para conectarse a la bsae de datos que creamos

	db.DBConnection()

	//usando los modelos que creamos , que ejecute y cree las tablas

	db.DB.AutoMigrate(models.Tarea{})   //importamos la estrcutura Tarea
	db.DB.AutoMigrate(models.Usuario{}) //importamos la estrcutura Usuario

	//usamos el paquete gorilla mux para crear las rutas
	r := mux.NewRouter()

	//pagina principal
	r.HandleFunc("/", rutas.RutaHome)

	//obtener todos los usuarios
	r.HandleFunc("/usuarios", rutas.ListarUsuarios).Methods("GET")

	//obtener  un usuario por ID
	r.HandleFunc("/usuarios/{id}", rutas.VerUsuarioId).Methods("GET")

	//crear usuario
	r.HandleFunc("/usuarios", rutas.CrearUsuario).Methods("POST")

	//eliminar usuarios
	r.HandleFunc("/usuarios/{id}", rutas.BorrarUsuario).Methods("DELETE")

	//funciones para tareas

	//listar tareas
	r.HandleFunc("/tareas", rutas.ListarTareas).Methods("GET")

	//crear tarea
	r.HandleFunc("/tareas", rutas.CrearTarea).Methods("POST")

	//ver una tarea por id
	r.HandleFunc("/tareas/{id}", rutas.VerTarea).Methods("GET")

	//eliminar tarea
	r.HandleFunc("/tareas/{id}", rutas.BorrarTarea).Methods("DELETE")

	//activamos el servidor con el puerto y con la funcionalidad de arriba
	http.ListenAndServe(":3000", r)

}
