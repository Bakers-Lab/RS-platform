package v1

import (
	"RS-Backend/dal/db"
	// "RS-Backend/models/dao"
	service "RS-Backend/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type InferHandler struct {
	service service.InferService
}

func RegisterInferRoutes(router *gin.RouterGroup, dB db.IDB) {
	api := &InferHandler{service: service.NewInferService(db.NewInferAccesser(dB))}
	router.GET("/inferjobs", api.GetAllInferJobs)
	router.GET("/inferjobs/:id", api.GetInferJobById)
}

func (api *InferHandler) GetAllInferJobs(c *gin.Context) {
	InferJobs, err := api.service.GetAllInferJobs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, InferJobs)
}

func (api *InferHandler) GetInferJobById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{Error: "invalid id"})
		return
	}

	inferJob, err := api.service.GetInferJobById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: "invalid id"})
		return
	}
	c.JSON(http.StatusOK, inferJob)
}