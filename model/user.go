package model

import (
	"database/sql"
	"errors"
)

type User struct {
	ID   int
	Name string
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db}
}

func (m *UserModel) Insert(name string) (int, error) {
	stmt := `insert into user (name) values (?)`

	result, err := m.db.Exec(stmt, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *UserModel) Get(id uint) (*User, error) {
	stmt := `select id, name from user where id = ?`
	row := m.db.QueryRow(stmt, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNoRecord
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (m *UserModel) GetAll() ([]*User, error) {
	stmt := `select id, name from user`
	rows, err := m.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) Update(id uint, name string) error {
	stmt := `update user set name = ? where id = ?`
	_, err := m.db.Exec(stmt, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) Delete(id uint) error {
	stmt := `delete from user where id = ?`
	_, err := m.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
