package main

import (
	"net/http"
	"play-go/moviedb"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/push/movie", func(ctx *gin.Context) {

		var m moviedb.Movie

		if err := ctx.ShouldBindJSON(&m); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}

		moviedb.CreateOrUpdate(&m)

		color.Green("postjson value is %#v", m)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
		})
	})

	r.GET("/movie/:movieName", func(ctx *gin.Context) {
		movieName := ctx.Param("movieName")

		m, err := moviedb.GetFromName(movieName)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, m)
	})

	r.Run()
}
