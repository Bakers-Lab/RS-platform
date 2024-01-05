package v1

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	service "RS-Backend/services"
	"net/http"
	"strconv"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DatasetHandler struct {
	service service.DatasetService
}

func RegisterDatasetRoutes(router *gin.RouterGroup, dB db.IDB) {
	api := &DatasetHandler{service: service.NewDatasetService(db.NewDatasetAccesser(dB))}
	router.GET("/datasets", api.GetAllDatasets)
	router.GET("/datasets/:id", api.GetDatasetById)
	router.POST("/datasets", api.InsertDataset)
	router.POST("/uploadFile", api.UploadFile)
}

// GetAllDatasets godoc
// @Summary Get all datasets
// @Description Retrieves a list of all datasets
// @Tags datasets
// @Accept  json
// @Produce  json
// @Success 200 {array} dao.Dataset "List of datasets"
// @Failure 500 {object} v1.APIError "Internal server error"
// @Router /api/v1/datasets [get]
func (api *DatasetHandler) GetAllDatasets(c *gin.Context) {
	datasets, err := api.service.GetAllDatasets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, datasets)
}

// GetDatasetById godoc
// @Summary Get a dataset by ID
// @Description Retrieves a dataset based on its ID
// @Tags datasets
// @Accept  json
// @Produce  json
// @Param   id   path    int     true  "Dataset ID"
// @Success 200 {object} dao.Dataset "Dataset"
// @Failure 400 {object} v1.APIError "invalid id"
// @Failure 500 {object} v1.APIError "Internal server error"
// @Router /api/v1/datasets/{id} [get]
func (api *DatasetHandler) GetDatasetById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{Error: "invalid id"})
		return
	}

	dataset, err := api.service.GetDatasetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIError{Error: "invalid id"})
		return
	}
	c.JSON(http.StatusOK, dataset)
}

// InsertDataset godoc
// @Summary Insert a new dataset
// @Description Adds a new dataset to the database
// @Tags datasets
// @Accept  json
// @Produce  json
// @Param   dataset  body    dao.Dataset   true  "Dataset to be added"
// @Success 200 {string} string "Dataset inserted successfully"
// @Failure 400 {object} v1.APIError "Invalid input"
// @Failure 500 {object} v1.APIError "Internal server error"
// @Router /api/v1/datasets [post]
func (api *DatasetHandler) InsertDataset(c *gin.Context) {
	var dataset dao.Dataset
	if err := c.BindJSON(&dataset); err != nil {
		c.JSON(http.StatusBadRequest, APIError{Error: "invalid input"})
		return
	}

	err := api.service.InsertDataset(c, &dataset)
	if err != nil {

		// 这里你可能需要根据错误类型返回不同的状态码
		c.JSON(http.StatusInternalServerError, APIError{"Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dataset inserted successfully"})
}

func (api *DatasetHandler) UploadFile(c *gin.Context) {
	// var requestData map[string]interface{}

    // if err := c.BindJSON(&requestData); err != nil {
    //     // 处理绑定 JSON 数据失败的情况
    //     c.JSON(http.StatusBadRequest, gin.H{"error":  "Invalid input"})
    //     return
    // }
	// datasetIdFloat, ok := requestData["datasetId"].(float64)// 这里需要做类型断言
    // if !ok {
    //     // 处理类型断言失败的情况
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datasetId"})
    //     return
    // }
	// datasetId:= int64(datasetIdFloat)
    // name, ok := requestData["name"].(string)
    // if !ok {
    //     // 处理类型断言失败的情况
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name"})
    //     return
    // }
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}
	datasetIdStr:=  c.PostForm("datasetId")
	name:= c.PostForm("name") 
	datasetId, err := strconv.ParseInt(datasetIdStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datasetId"})
        return
    }


	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse files"})
		return
	}
	files := form.File["files"]

	path, err := api.service.FindDatasetPath(c, datasetId, name)
	if err != nil {
		// 这里你可能需要根据错误类型返回不同的状态码
		c.JSON(http.StatusInternalServerError, APIError{"Internal server error"})
		return
	}

	for _, file := range files {
		// 指定文件保存的完整路径
		// 例如，将文件保存到 /path/to/your/directory/ 目录下，并使用原始文件名作为保存的文件名
		uploadPath := filepath.Join(path+"/", file.Filename)
	
		// 保存文件到指定路径
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add files"})
			return
		}
	}
	// 输出一些信息，表示数据导入成功
	c.JSON(http.StatusOK, gin.H{"message": "Files inserted successfully"})
}


