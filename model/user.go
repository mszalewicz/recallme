package model

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
)

type UserModel struct {
	username string
	DB       *sql.DB
}

func (user *UserModel) UsernameExists(username string) (bool, error) {

	query, err := user.DB.Query(`SELECT * FROM users WHERE username = ? LIMIT 1`, username)
	if err != nil {
		return false, err
	}

	return query.Next(), nil

}

func (user *UserModel) EmailExists(email string) (bool, error) {

	query, err := user.DB.Query(`SELECT * FROM users WHERE email = ? LIMIT 1`, email)
	if err != nil {
		return false, err
	}

	return query.Next(), nil

}

func (user *UserModel) Add(username string, email string, password string) (bool, error) {

	hashedPassword := sha256.Sum256([]byte(password))
	hashedPasswordString := hex.EncodeToString(hashedPassword[:])

	sqlResult, err := user.DB.Exec(
		`INSERT INTO users (
		    username,
		    email,
		    password,
		    created_at,
		    updated_at
		    ) VALUES (
		    ?,
		    ?,
		    ?,
		    CURRENT_TIMESTAMP,
		    CURRENT_TIMESTAMP
	    )`,
		username,
		email,
		hashedPasswordString)

	if err != nil {
		return false, err
	}

	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return false, err
	}

	fmt.Println("Number of inserted users: ", rowsAffected)
	return true, nil // todo: right know placeholder, check if rowsAffected returns values, if not return false

}
