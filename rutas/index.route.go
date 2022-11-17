package rutas

import "net/http"

//funcion para ruta HOME

func RutaHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 3"))
}
