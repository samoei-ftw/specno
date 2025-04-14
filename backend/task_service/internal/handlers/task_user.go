package handlers

import (
	//"bytes" for debugging
	"encoding/json"
	"fmt"

	//"io" for debugging
	//"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/enums"
	"github.com/samoei-ftw/specno/backend/common/utils"
	dto "github.com/samoei-ftw/specno/backend/task_service/internal/models"
	services "github.com/samoei-ftw/specno/backend/task_service/internal/services"
)
func AssignUserHandler(service *services.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskIdStr := vars["id"]
		taskId, err := strconv.Atoi(taskIdStr)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid task ID",
			})
			return
		}

		var updateUserTaskRequest dto.UpdateTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&updateUserTaskRequest); err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid JSON payload",
			})
			return
		}
		//fmt.Printf("Decoded updateUserTaskRequest: %+v\n", updateUserTaskRequest)
		// TODO: validate user id
		
		//userId := updateUserTaskRequest.UserID
		//fmt.Printf("Request user id is: %s", userId)
		task, err := service.UpdateTask(uint(taskId), (*enums.TaskStatus)(updateUserTaskRequest.Status), updateUserTaskRequest.Name, updateUserTaskRequest.Description, updateUserTaskRequest.UserID)
		if err != nil {
			fmt.Printf("Error in updating task status: %v\n", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status:  "success",
			Message: "Task updated successfully",
			Data:    task,
		})
	}
}