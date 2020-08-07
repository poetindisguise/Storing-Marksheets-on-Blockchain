package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)


func CreateChannel(c *gin.Context) {
	name := c.Query("name")
	output, err:= exec.Command("./createchannel.sh", name).Output()
	c.Writer.Header().Set("Access-Control-Allow-Origin","*")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"err": err})
	}
	c.JSON(http.StatusOK,  string(output))



}


func main() {
	r := gin.Default()

	r.GET("/",CreateChannel)
	r.Static("/form",".")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

