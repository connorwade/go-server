package server

import (
	"net/http"

	"github.com/connorwade/go-server/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) getAlbums(ctx *gin.Context) {
	albums, err := server.store.All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

type postAlbumsRequest struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func (server *Server) postAlbums(ctx *gin.Context) {
	var req postAlbumsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateAlbumParams{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	album, err := server.store.Create(args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, album)
}

type GetAlbumByIDRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

func (server *Server) getAlbumByID(ctx *gin.Context) {
	var req GetAlbumByIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	album, err := server.store.GetByID(db.GetAlbumByIDArgs{ID: req.ID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, album)
}
