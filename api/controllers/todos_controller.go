package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-todo-api/api/httpext"
	"go-todo-api/api/models"
	"go-todo-api/api/repository"
	"io"
	"net/http"
	"strconv"
	"time"
)

type todosControllers struct {
	repo repository.TODOSRepository
}

func RegisterTODOSControllers(router *mux.Router) {
	c := &todosControllers{repo: repository.NewTODOSRepository()}

	router.Path("/todos").HandlerFunc(c.PostTODO).Methods(http.MethodPost)
	router.Path("/todos").HandlerFunc(c.GetTODOS).Methods(http.MethodGet)
	router.Path("/todos/{id}").HandlerFunc(c.GetTODO).Methods(http.MethodGet)
	router.Path("/todos/{id}").HandlerFunc(c.PutTODO).Methods(http.MethodPut)
	router.Path("/todos/{id}").HandlerFunc(c.DeleteTODO).Methods(http.MethodDelete)
}

func (c *todosControllers) PostTODO(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	todo := new(models.TODO)

	err = json.Unmarshal(body, todo)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	err = todo.Validate()
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	todo.ID = c.repo.Count() + 1
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = todo.CreatedAt

	c.repo.Create(todo)

	location := fmt.Sprintf("%s/%d", r.RequestURI, todo.ID)
	w.Header().Set("Location", location)

	httpext.WriteJSON(w, http.StatusCreated, todo)
}

func (c *todosControllers) GetTODO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	todo, err := c.repo.Get(id)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	httpext.WriteJSON(w, http.StatusOK, todo)
}

func (c *todosControllers) GetTODOS(w http.ResponseWriter, r *http.Request) {
	httpext.WriteJSON(w, http.StatusOK, c.repo.GetAll())
}

func (c *todosControllers) PutTODO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	current, err := c.repo.Get(id)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	todo := new(models.TODO)

	err = json.Unmarshal(body, todo)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	err = todo.Validate()
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	todo.ID = id
	todo.CreatedAt = current.CreatedAt
	todo.UpdatedAt = time.Now()

	err = c.repo.Update(todo)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	httpext.WriteJSON(w, http.StatusOK, todo)
}

func (c *todosControllers) DeleteTODO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	err = c.repo.Delete(id)
	if err != nil {
		httpext.WriteJSONError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprint(id))
	w.WriteHeader(http.StatusNoContent)
}
