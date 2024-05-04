package database

import "fmt"

type ArtInfo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetAllFiles() ([]ArtInfo, error) {
	var files []ArtInfo
	result := Db.Table("files").Select("id, title, description").Find(&files)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get files: %v", result.Error)
	}
	return files, nil
}
