package models

type PlaylistTrack struct {
	PlaylistId int `json:"playlist_id" gorm:"column:PlaylistId;primaryKey"`
	TrackId    int `json:"track_id" gorm:"column:TrackId;primaryKey"`
}

// TableName specifies the exact table name in the Chinook DB so GORM
// uses the correct (case-sensitive) identifier.
func (PlaylistTrack) TableName() string {
	return "PlaylistTrack"
}
