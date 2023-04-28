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

	r.GET("/movie/id/:movieid", func(ctx *gin.Context) {
		movieid := ctx.Param("movieid")

		m, err := moviedb.GetFromName(movieid)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    m,
		})
	})

	r.GET("/movie/list", func(ctx *gin.Context) {

		var p moviedb.Page

		if err := ctx.ShouldBindQuery(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}

		m, err := moviedb.GetMovieList(p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "success",
			"data":    m,
		})
	})

	r.Run("0.0.0.0:8000")
}
