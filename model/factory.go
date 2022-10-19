package model

import "database/sql"

type Models struct {
	UserModel *UserModel
}

func CreateModels(db *sql.DB) *Models {
	return &Models{
		UserModel: &UserModel{db: db},
	}
}
