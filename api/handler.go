package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jbground/kalman-filter/app"
	"log"
	"net/http"
	"text/template"
)

func KalmanIndex(c *gin.Context) {

	files, err := template.ParseFiles("asset/filter.html")
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	err = files.Execute(c.Writer, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return
}

func KalmanUniform(c *gin.Context) {
	var data []*app.Coordinate
	data = app.GenerateUniformData()

	va, _ := json.Marshal(data)
	fmt.Println(string(va))

	c.JSON(http.StatusOK, app.GenerateUniformData())
}

func LoggingBodyContent(ctx *gin.Context) {
	jsons := make([]byte, ctx.Request.ContentLength)
	if _, err := ctx.Request.Body.Read(jsons); err != nil {
		if err.Error() != "EOF" {
			return
		}
	}
	fmt.Println(string(jsons))
}

func KalmanPdf(c *gin.Context) {
	//LoggingBodyContent(c)

	var val []app.Coordinate
	if err := c.ShouldBind(&val); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, app.GeneratePdfData())
}

func KalmanExecute(c *gin.Context) {
	var val []app.Coordinate
	if err := c.ShouldBind(&val); err != nil {
		fmt.Println(err)
	}

	app.ExecuteKalmanFilterByLib(val)

	c.JSON(http.StatusOK, val)

}
