package db

func (db *DB) GetTrack(title string) (*Track, error) {
	var track Track
	result := db.db.Where(&Track{Title: title}).First(&track)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}
