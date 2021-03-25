package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ameernormie/go-api-template/pkg/handlers"
	"github.com/ameernormie/go-api-template/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.GetEnv("GQL_SERVER_HOST")
	port = utils.GetEnv("GQL_SERVER_PORT")
	gqlPath = utils.GetEnv("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.GetEnv("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.GetBoolEnv("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

type App struct {
	Db *gorm.DB
}

func (a *App) Close() {
	a.Db.Close()
}

type RequestQuery struct {
	BusinessNumber                string `form:"business_number" json:"business_number"`
	CorporationRegistrationNumber string `form:"corporation_registration_number" json:"corporation_registration_number"`
}

func (a *App) GetRouter() *gin.Engine {
	endpoint := "http://" + host + ":" + port
	r := gin.Default()
	if isPgEnabled {
		r.GET("/", handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func (a *App) Start() {
	router := a.GetRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// close db and socket connection
	a.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

var now = time.Now

func today() string {
	return now().Format("20060102")
}
