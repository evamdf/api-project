package models

type Artist struct {
	ArtistId int    `json:"artist_id" gorm:"column:ArtistId;primaryKey"`
	Name     string `json:"name" gorm:"column:Name"`
}

func (Artist) TableName() string {
	return "Artist"
}
