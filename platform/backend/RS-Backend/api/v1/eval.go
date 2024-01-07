package v1

import (
	"RS-Backend/dal/db"
	// "RS-Backend/models/dao"
	service "RS-Backend/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type EvalHandler struct {
	service service.EvalService
}

func RegisterEvalRoutes(router *gin.RouterGroup, dB db.IDB) {
	api := &EvalHandler{service: service.NewEvalService(db.NewEvalAccesser(dB))}
	router.GET("/evaljobs", api.GetAllEvalJobs)
	router.GET("/evaljobs/:id", api.GetEvalJobById)
}

func (api *EvalHandler) GetAllEvalJobs(c *gin.Context) {
	EvalJobs, err := api.service.GetAllEvalJobs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, EvalJobs)
}

func (api *EvalHandler) GetEvalJobById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{Error: "invalid id"})
		return
	}

	evaljob, err := api.service.GetEvalJobById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: "invalid id"})
		return
	}
	c.JSON(http.StatusOK, evaljob)
}