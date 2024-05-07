package db

func (db *DB) GetUser(handle string) (*User, error) {
	var user User
	result := db.db.Where(&User{Handle: handle}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *DB) InsertUser(handle, bio string) (*User, error) {
	model := User{
		Handle: handle,
		Bio:    bio,
	}

	result := db.db.Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
