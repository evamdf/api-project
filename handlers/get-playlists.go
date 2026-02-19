package handlers

import (
	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func GetPlaylists(c *gin.Context) {
	// initialize slice so JSON marshals to [] (empty array) instead of null
	var playlists []models.Playlist = []models.Playlist{}
	database.DB.Find(&playlists)
}
