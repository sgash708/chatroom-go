package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// h is host?
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serverWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:9000")
}
