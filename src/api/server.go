package api

import (
	"app-example/src/api/controllers"
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	"syscall"
)

// Init -
func Init() {
	r := NewRouter()
	r.Run(":8080")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()

	<-done
}

// NewRouter -
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	appCtrl := new(controllers.AppController)
	router.GET("/", appCtrl.Home)

	return router
}
