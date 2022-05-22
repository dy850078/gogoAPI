package controllers

import (
	"net/http"

	"github.com/dy850078/gogoAPI/api/models"
	"github.com/dy850078/gogoAPI/api/responses"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}
