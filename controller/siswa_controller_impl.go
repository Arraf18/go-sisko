package controller

import (
	"github.com/Arraf18/go-sisko/helper"
	"github.com/Arraf18/go-sisko/model/web"
	"github.com/Arraf18/go-sisko/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type SiswaControllerImpl struct {
	SiswaService service.SiswaService
}

func NewSiswaController(siswaService service.SiswaService) SiswaController {
	return &SiswaControllerImpl{
		SiswaService: siswaService,
	}
}

func (controller *SiswaControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaCreateRequest := web.SiswaCreateRequest{}
	helper.ReadFromRequestBody(request, &siswaCreateRequest)

	siswaResponse := controller.SiswaService.Create(request.Context(), siswaCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SiswaControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaUpdateRequest := web.SiswaUpdateRequest{}
	helper.ReadFromRequestBody(request, &siswaUpdateRequest)

	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	siswaUpdateRequest.Id = id

	categoryResponse := controller.SiswaService.Update(request.Context(), siswaUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SiswaControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	controller.SiswaService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SiswaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	siswaResponse := controller.SiswaService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SiswaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaResponses := controller.SiswaService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
