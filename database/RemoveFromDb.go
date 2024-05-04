package database

func RemoveFromDb(id int) error {
	result := Db.Delete(&File{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
