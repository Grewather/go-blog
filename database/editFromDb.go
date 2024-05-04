package database

import (
	"fmt"
)

func EditDb(file *File) error {
	result := Db.Model(&File{}).Where("id = ?", file.ID).Updates(File{
		Title:       file.Title,
		Description: file.Description,
		FilePath:    file.FilePath,
	})

	if result.Error != nil {
		fmt.Print(result.Error)
		return result.Error
	}

	return nil
}
