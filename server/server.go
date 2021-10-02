package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"go/server/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {

	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	log.Print("Server is running at port:", s.port)

	router := routes.ConfingRoutes(s.server)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
	log.Fatal(router.Run(":" + s.port))

}
