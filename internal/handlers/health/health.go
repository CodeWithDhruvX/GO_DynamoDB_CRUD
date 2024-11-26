package Health

import (
	"errors"
	"net/http"
	"github.com/akhil/dynamodb-go-crud-yt/internal/repository/adapter"
	"github.com/akhil/dynamodb-go-crud-yt/internal/handlers"
)

type Handler struct{
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository:repository
	}
}

func(h *Handler) Get(w http.ResponseWriter,r *http.Request){

	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w,r,errors.New("Relational database not live"))
		return
	}

	HttpStatus.StatusOK(w,r,"Service OK")
}

func (h *Handler) Put(w http.ReponseWriter,r *http.HttpRequest){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Post(w http.ResponseWriter,r *http.Request){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Delete(w http.ReponseWriter,r *http.Request){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Options(w http.ReponseWriter,r *http.Request){
	HttpStatus.StatuseNoContent(w,r)
}