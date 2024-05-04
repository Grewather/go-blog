package database

import "fmt"

func AddToDb(file *File) error {
	result := Db.Create(file)
	if result.Error != nil {
		fmt.Print(result.Error)
		return result.Error
	}
	return nil
}
