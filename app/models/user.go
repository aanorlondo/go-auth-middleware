package models

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"app/config"
	"app/utils"
)

// ORM
type User struct {
	Username string
	Password string
}

var logger = utils.GetLogger()
var databaseUrl = ""
var databaseTableName = ""

// INIT
func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	databaseUrl = cfg.GetDatabaseURL()
	databaseTableName = cfg.GetDatabaseTableName()
}

// CRUD
// // - CREATE (INSERT)
func (u *User) Save() error {
	logger.Info("Saving new user ", u.Username, " to database...")
	logger.Info("Connecting to database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to database: ", err)
		return err
	}
	defer db.Close()
	logger.Info("Executing query...")
	_, err = db.Exec("INSERT INTO "+databaseTableName+" (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		logger.Error("ERROR when inserting new user to database: ", err)
		return err
	}
	logger.Info("User successfully inserted.")
	return nil
}

// // - READ (SELECT)
func GetUserByUsername(username string) (*User, error) {
	logger.Info("Getting user by username: ", username)
	logger.Info("Connecting to the database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to the database: ", err)
		return nil, err
	}
	defer db.Close()
	logger.Info("Executing query...")
	row := db.QueryRow("SELECT username, password FROM "+databaseTableName+" WHERE username = ?", username)
	user := &User{}
	err = row.Scan(&user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Info("User not found with username: ", username)
			return nil, nil
		}
		logger.Error("ERROR when retrieving user: ", err)
		return nil, err
	}
	logger.Info("User retrieved successfully: ", user.Username)
	return user, nil
}

// // - UPDATE
func (u *User) Update() error {
	logger.Info("Updating user: ", u.Username)
	logger.Info("Connecting to the database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to the database: ", err)
		return err
	}
	defer db.Close()
	logger.Info("Executing query...")
	_, err = db.Exec("UPDATE "+databaseTableName+" SET password = ? WHERE username = ?", u.Password, u.Username)
	if err != nil {
		logger.Error("ERROR when updating user: ", err)
		return err
	}
	logger.Info("User updated successfully: ", u.Username)
	return nil
}
