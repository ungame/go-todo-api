package todos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-todo-api/api/models"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func Run(baseURL string) {
	todo, err := CreateTODO(baseURL)
	if err != nil {
		log.Fatalf("error on create todo: %s\n", err.Error())
	}

	log.Println("Create TODO successfully!")

	err = GetTODO(baseURL, todo)
	if err != nil {
		log.Fatalf("error on get todo: %s\n", err.Error())
	}

	log.Println("Get TODO successfully!")
}

func CreateTODO(baseURL string) (*models.TODO, error) {

	todo := &models.TODO{
		Title: "Integration Test",
		Task:  "This a Test",
	}

	body, err := json.Marshal(todo)
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s/todos", baseURL)

	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res != nil && res.Body != nil {
		defer res.Body.Close()

		if res.StatusCode == http.StatusCreated {

			body, err = io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(body, todo)
			if err != nil {
				return nil, err
			}

			return todo, nil
		}
	}

	d, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}

	log.Println(string(d))

	return nil, errors.New("can't create TODO")
}

func GetTODO(baseURL string, todo *models.TODO) error {

	uri := fmt.Sprintf("%s/todos/%d", baseURL, todo.ID)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res != nil && res.Body != nil {
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			return nil
		}
	}

	d, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err
	}

	log.Println(string(d))

	return errors.New("can't get TODO")
}
