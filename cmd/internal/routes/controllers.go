package routes

import (
	"net/http"

	"github.com/kaaninchen/laterna/server/cmd/internal/db"
	"github.com/kaaninchen/laterna/server/cmd/internal/middleware"
	"github.com/kaaninchen/laterna/server/cmd/internal/models"
	"github.com/kaaninchen/laterna/server/cmd/utils"
)

func RegisterControllerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/controllers", middleware.WithAdminAuth(createController))
	mux.HandleFunc("DELETE /api/v1/controllers/{ID}", middleware.WithAdminAuth(deleteController))

	mux.HandleFunc("POST /api/v1/controllers/toggle/{ID}", middleware.WithAdminAuth(toggleControllerState))
}

func createController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ID, err := db.CreateController()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "An error occured. Check the server logs for more information")
	}
	utils.SuccessResponse(w, http.StatusOK, models.ControllerRequests{ID: int(ID)})
}

func deleteController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue(("ID"))
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !db.ControllerExists(id) {
		utils.ErrorResponse(w, http.StatusNotFound, "Lamp not found")
		return
	}

	if db.DeleteController(id) != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "An error occured. Check the server logs for more information")
		return
	}
	utils.SuccessResponse(w, http.StatusOK, models.DeleteData{Deleted: id})
}

func toggleControllerState(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt((idStr))
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !db.ControllerExists((id)) {
		utils.ErrorResponse(w, http.StatusNotFound, "Controller not found")
		return
	}

	currentState, err := db.GetControllerState(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = db.SetControllerstate(!currentState, id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, models.ToggleController{ID: id, Active: !currentState})
}
