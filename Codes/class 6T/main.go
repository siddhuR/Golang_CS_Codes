package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	post := ;
	r := gin.Default()	//application will be configured on port 8080
	pingFunc := func(c *gin.Context) {

		key := map[string]any{
			
		}
		res := gin.H{
			"messgage": "pong",
		}
		key := map[string]any{
			
		}
		c.keys["ping"] = "pingu"
		c.JSON(http.StatusOK, res)
	}

	pongFunc := func(c *gin.Context) {
		res := gin.H{
			"messgage": "pong",
			"contextvalue": c.keys["ping"],
		}
		c.JSON(http.StatusOK, res)
	}

	dingFunc := func(c *gin.Context) {
		res := gin.H{
			"messgage": "pong",
		}
		c.JSON(http.StatusOK, res)
	}

	dongFunc := func(c *gin.Context) {
		res := gin.H{
			"messgage": "pong",
		}
		c.JSON(http.StatusOK, res)
	}

	r,GET("/ping", pingFunc)	//https://localhost:8080/pinng
	r,GET("/pong", pongFunc)
	r,GET("/ding", dingFunc)
	r,GET("/dong", dongFunc)

	r.Run()	 //listen and serve on 0.0.0.0:8080 (for windows "localhost:8080" )
	
}