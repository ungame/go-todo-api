package repository

import (
	"go-todo-api/api/fakedb"
	"go-todo-api/api/models"
)

type (
	TODOSRepository interface {
		Create(todo *models.TODO)
		Get(id int) (*models.TODO, error)
		GetAll() []*models.TODO
		Update(todo *models.TODO) error
		Delete(id int) error
		Count() int
	}

	todosRepository struct {
		db fakedb.Memory
	}
)

func NewTODOSRepository() TODOSRepository {
	return &todosRepository{db: fakedb.New()}
}

func (r *todosRepository) Create(todo *models.TODO) {
	r.db.Set(todo.ID, todo)
}

func (r *todosRepository) Get(id int) (*models.TODO, error) {
	if r.db.Has(id) {
		item, ok := r.db.Get(id).(*models.TODO)
		if ok {
			return item, nil
		}
	}
	return nil, models.NotFound(models.TODO{}, id)
}

func (r *todosRepository) GetAll() []*models.TODO {
	todos := make([]*models.TODO, 0, r.db.Len())

	r.db.ForEach(func(item interface{}) bool {
		if todo, ok := item.(*models.TODO); ok {
			todos = append(todos, todo)
		}
		return true
	})

	return todos
}

func (r *todosRepository) Update(todo *models.TODO) error {
	if r.db.Has(todo.ID) {
		r.db.Set(todo.ID, todo)
		return nil
	}
	return models.NotFound(models.TODO{}, todo.ID)
}

func (r *todosRepository) Delete(id int) error {
	if r.db.Has(id) {
		r.db.Del(id)
		return nil
	}
	return models.NotFound(models.TODO{}, id)
}

func (r *todosRepository) Count() int {
	return r.db.Len()
}
