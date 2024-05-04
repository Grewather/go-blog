package database

func GetFromDb(id int) (*File, error) {
	var file File

	result := Db.First(&file, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &file, nil
}
