package app

import (
	"github.com/Arraf18/go-sisko/controller"
	"github.com/Arraf18/go-sisko/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(siswaController controller.SiswaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/siswas", siswaController.FindAll)
	router.GET("/api/siswas/:siswaId", siswaController.FindById)
	router.POST("/api/siswas", siswaController.Create)
	router.PUT("/api/siswas/:siswaId", siswaController.Update)
	router.DELETE("/api/siswas/:siswaId", siswaController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
