package rutas

import (
	"encoding/json"
	"net/http"

	"github.com/MariaSusanaGallardo/api-rest/db"
	"github.com/MariaSusanaGallardo/api-rest/models"
	"github.com/gorilla/mux"
)

//funciones para tareas

func ListarTareas(w http.ResponseWriter, r *http.Request) {

	//creamos una variable para poder usar el modelo
	var tareas []models.Tarea

	//buscará la variable y traerá los datos
	db.DB.Find(&tareas)

	//para traerlos los codificamos con json
	json.NewEncoder(w).Encode(&tareas)

}

func CrearTarea(w http.ResponseWriter, r *http.Request) {

	//creamos una variable para poder usar el modelo
	var tarea models.Tarea

	//leer lo que viene del request body

	json.NewDecoder(r.Body).Decode(&tarea)

	//creamos la nueva tarea con la función create de db
	tareaCreada := db.DB.Create(&tarea)

	//si me devuelve un error lo extremos

	err := tareaCreada.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//si no hay error
	json.NewEncoder(w).Encode(tarea)

}

func VerTarea(w http.ResponseWriter, r *http.Request) {
	//creamos variable de Tareas para guardar parametros
	var tarea models.Tarea

	//usamos VARS DE MUX para accedeor a las variables de request
	params := mux.Vars(r)

	//hacemos la consulta para que muestre al primero que encuentre con first
	db.DB.First(&tarea, params["id"])

	//validamos que exista la tarea
	if tarea.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tarea no encontrada"))
	}

	//si pasa la validación de arriba, lo devuelvo con json
	json.NewEncoder(w).Encode(&tarea)

}

func BorrarTarea(w http.ResponseWriter, r *http.Request) {
	//creamos variable de Tareas para guardar parametros
	var tarea models.Tarea

	//usamos VARS DE MUX para accedeor a las variables de request
	params := mux.Vars(r)

	//hacemos la consulta para que muestre al primero que encuentre con first
	db.DB.First(&tarea, params["id"])

	//validamos que exista la tarea
	if tarea.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tarea no encontrada"))
		return
	}

	//eliminamos la tarea
	db.DB.Unscoped().Delete(&tarea)
	w.WriteHeader(http.StatusNoContent)

}
