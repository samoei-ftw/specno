package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/utils"
)

func CreateTaskHandler(service *service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var taskCreateRequest struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			ProjectID   uint   `json:"project_id"`
			UserID      uint   `json:"user_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&taskCreateRequest); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid request payload",
			})
			return
		}
		if taskCreateRequest.Title == "" || taskCreateRequest.Description == "" || taskCreateRequest.ProjectID == 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Validation failure",
			})
			return
		}
		task, err := service.CreateTask(taskCreateRequest.Title, taskCreateRequest.Description, taskCreateRequest.UserID, taskCreateRequest.ProjectID, enums.Todo)
		if err != nil {
			if err.Error() == "user not found" {
				utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
					Status:  "error",
					Message: "User not found",
				})
				return
			}
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to create task",
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusCreated, utils.Response{
			Status: "success",
			Data:   task,
		})
	}
}

func ListTasksHandler(service *service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projIdStr := vars["project_id"]

		projId, err := strconv.Atoi(projIdStr)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid project ID",
			})
			return
		}

		tasks, err := service.ListTasks(uint(projId))
		if err != nil {
			if err.Error() == "not found" {
				utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
					Status:  "error",
					Message: "not found",
				})
				return
			}

			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to retrieve tasks",
			})
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data:   tasks,
		})
	}
}

func GetTaskHandler(service *service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskIdStr := r.URL.Query().Get("id")

		taskId, err := strconv.Atoi(taskIdStr)
		if err != nil || taskId <= 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid task ID",
			})
			return
		}

		task, err := service.GetTask(uint(taskId))
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Failed to retrieve task",
			})
			return
		}

		if task == nil {
			utils.RespondWithJSON(w, http.StatusNotFound, utils.Response{
				Status:  "error",
				Message: "Task not found",
			})
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data:   task,
		})
	}
}