package health

import (
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/util"
)

type healthController struct{}

func (h *healthController) GetAppHealth(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	util.HttpSuccessResponse(w, http.StatusOK, "App is running fine", data)
}

var HealthController = healthController{}
