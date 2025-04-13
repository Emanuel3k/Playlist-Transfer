package config

import (
	"fmt"
	"github.com/emanuel3k/playlist-transfer/cmd/http/routes"
	"log"
	"net/http"
	"os"
)

var (
	HttpServerPort = "HTTP_SERVER_PORT"
)

func InitHTTPServer() error {
	r := routes.NewRouter()

	port := os.Getenv(HttpServerPort)

	log.Println("starting http server on port", port, "🚀")

	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.InitRoutes())
}
