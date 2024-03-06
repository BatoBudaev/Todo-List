package handlers

import (
	"encoding/json"
	"github.com/BatoBudaev/Todo-List/internal/db"
	"github.com/BatoBudaev/Todo-List/internal/models"
	"github.com/BatoBudaev/Todo-List/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(r *mux.Router, db *db.DB) {
	r.HandleFunc("/tasks", getTasks(db)).Methods("GET")
	r.HandleFunc("/tasks", createTask(db)).Methods("POST")
}

func createTask(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err = db.CreateTask(task)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, task)
	}
}

func getTasks(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := db.GetTasks()

		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tasks)
	}
}
