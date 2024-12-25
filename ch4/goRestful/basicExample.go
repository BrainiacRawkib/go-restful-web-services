package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a web service
	webservice := new(restful.WebService)

	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))

	// Add the service to application
	restful.Add(webservice)

	log.Println("Running Server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func pingTime(req *restful.Request, res *restful.Response) {
	// Write to the response
	io.WriteString(res, fmt.Sprintf("%s\n", time.Now()))
}
