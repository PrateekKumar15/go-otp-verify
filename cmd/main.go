package main

import (
	"github.com/PrateekKumar15/go-otp-verify/api"
	 "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// initialize configuration
	app := api.Config{Router: router}
	app.Routes()

	router.Run(":8000")
	// run the server
	// log.Fatal(http.ListenAndServe(":8000", router))
}