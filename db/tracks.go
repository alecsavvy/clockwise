package db

func (db *DB) GetTrack(title string) (*Track, error) {
	var track Track
	result := db.db.Where(&Track{Title: title}).First(&track)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}

func (db *DB) InsertTrack(title, streamUrl, description string, userID uint) (*Track, error) {
	model := Track{
		Title:       title,
		StreamUrl:   streamUrl,
		Description: description,
		UserID:      userID,
	}

	result := db.db.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
