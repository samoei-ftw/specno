package internal

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/samoei-ftw/specno/backend/common/utils"
	service "github.com/samoei-ftw/specno/backend/project_service/internal/service"
	auth "github.com/samoei-ftw/specno/backend/user_service/pkg"
)
func GetProjectOwnerHandler(service *service.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		vars := mux.Vars(r)
		projectIdStr := vars["project_id"]
		//var rf := auth.ExtractTokenFromHeader(r);
		projectId, err := strconv.Atoi(projectIdStr)
		log.Printf("Invalid project id, %s", projectIdStr)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.Response{
				Status:  "error",
				Message: "Invalid project ID:",
			})
			return
		}
		projectUserId, err := service.GetUserForProject(uint(projectId), tokenStr) 
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, utils.Response{
				Status:  "error",
				Message: "Error fetching user for project",
			})
			return
		}
		claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
		if !ok || claims == nil {
			utils.RespondWithJSON(w, http.StatusUnauthorized, utils.Response{
				Status:  "error",
				Message: "Unauthorized: user ID not found",
			})
			return
		}
		log.Printf("Claims User ID: %d, Project Owner ID: %d", claims.UserID, projectUserId)
		if uint(claims.UserID) != projectUserId {
			utils.RespondWithJSON(w, http.StatusForbidden, utils.Response{
				Status:  "error",
				Message: "User is not the project owner",
			})
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, utils.Response{
			Status: "success",
			Data: utils.GetOwnerResponse{
				IsOwner: true,
			},
		})
	}
}