package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type FileController struct{}

func NewFileController() FileController {
	return FileController{}
}

func (f *FileController) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}
	c.SaveUploadedFile(file, filepath.Join("static", filepath.Base(file.Filename)))
	c.JSON(200, gin.H{"msg": "File saved", "data": "/api/v2/file/" + file.Filename})
}

func (f *FileController) SendFileUpload(c *gin.Context) {
	c.File("static/" + c.Param("file"))
}
