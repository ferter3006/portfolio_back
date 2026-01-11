package controllers

import (
	"encoding/json"
	"net/http"
	"new-test/models"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

// Simulación de almacenamiento en memoria
var posts = []models.Post{}

// Crear un nuevo post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al decodificar el post"))
		return
	}
	post.Created = time.Now()
	post.Updated = time.Now()
	posts = append(posts, post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func HiWorld(e *core.RequestEvent) error {
	return e.JSON(http.StatusOK, map[string]string{"message": "Hello, World! desde controller"})
}
