package repository

import (
	"api/adapter/logger"
	"api/domain"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
	l logger.Logger
}

func NewUserRepository(db *sql.DB, l logger.Logger) *UserRepository {
	return &UserRepository{db, l}
}

func (ur *UserRepository) Store(u domain.User) error {
	stmt, err := ur.DB.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Username, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindById() {}
