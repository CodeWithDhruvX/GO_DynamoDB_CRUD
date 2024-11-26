package handlers

import "net/http"

type Interface interface {
	Get(w http.ReponseWriter,r *http.Request)
	Post(w http.ReponseWriter,r *http.Request)
	Put(w http.ReponseWriter,r *http.Request)
	Delete(w http.ReponseWriter,r *http.Request)
	Options(w http.ReponseWriter,r *http.Request)
}