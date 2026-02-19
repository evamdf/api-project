package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPlaylistTracks returns all tracks belonging to the playlist specified by
// the :id URL parameter. Response uses models.Response for a consistent API.
func GetPlaylistTracks(c *gin.Context) {
	// Parse and validate id from url
	idParam := c.Param("id")
	playlistID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Success: false, Error: "invalid playlist id"})
		return
	}

	// Ensure playlist exists
	var playlist models.Playlist
	if err := database.DB.First(&playlist, "PlaylistId = ?", playlistID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, models.Response{Success: false, Error: "playlist not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Error: "failed to fetch playlist"})
		return
	}

	// Query playlist tracks
	var tracks []models.PlaylistTrack
	if err := database.DB.Where("PlaylistId = ?", playlistID).Find(&tracks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Success: false, Error: "failed to fetch playlist tracks"})
		return
	}

	// Return wrapped response
	payload := map[string]interface{}{
		"playlist_id":   playlist.PlaylistId,
		"playlist_name": playlist.Name,
		"track_count":   len(tracks),
		"tracks":        tracks,
	}

	c.JSON(http.StatusOK, models.Response{Success: true, Data: payload})
}
