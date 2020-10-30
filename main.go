package main

import (
	"fmt"

	"github.com/lrs-studies/poc-go-firebase-auth/middleware"

	"github.com/lrs-studies/poc-go-firebase-auth/api"
	"github.com/lrs-studies/poc-go-firebase-auth/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize new gin engine (for server)
	r := gin.Default()
	// create/configure database instance
	db := config.CreateDatabase()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set db to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})
	r.Use(middleware.AuthMiddleware)
	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)
	// start the server
	fmt.Println("server started...")
	r.Run(":5000")
}
