package url

import (
	"encoding/json"
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/user"
	"github.com/GabriellaAmah/go-url-shortner/util"
)

type IUrlController struct {
	urlService IUrlService
}

func (controller IUrlController) CreateUrl(res http.ResponseWriter, req *http.Request){
	var createUrl  IShortenUrl

	user := req.Context().Value("user").(user.User)

	err := json.NewDecoder(req.Body).Decode(&createUrl)
	if err != nil {
		util.HttpErrorResponse(res, util.MakeError{Code: 500, Message: "Internal server error", Error: err})
		return
	}

	data, serviceErr := controller.urlService.CreateShortenUrl(user.Id, createUrl)
	if serviceErr.IsError {
		util.HttpErrorResponse(res, serviceErr)
		return
	}

	util.HttpSuccessResponse(res, 200, "Url successfully created", data)
}

var UrlController = IUrlController{urlService: UrlService}