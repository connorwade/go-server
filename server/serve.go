package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/connorwade/go-server/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(cors.Default())

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "/var/www")

	router.GET("/index", server.indexRequest)
	router.GET("/albums", server.getAlbums)
	router.GET("/albums/:id", server.getAlbumByID)
	router.POST("/newalbums", server.postAlbums)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) indexRequest(ctx *gin.Context) {
	music, err := server.store.All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.HTML(http.StatusOK, "index.html", music)
}
