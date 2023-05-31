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
	ID       int
	Username string
	Password string
	// Add other fields as required
}

var logger = utils.GetLogger()
var databaseUrl = ""

// INIT
func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	databaseUrl = cfg.GetDatabaseURL()
}

// CRUD

// // - CREATE (INSERT)
func (u *User) Save() error {
	logger.Info("Saving new user ", u, " to database...")

	// Open a connection to the MySQL database
	logger.Info("Connecting to database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to database: ", err)
		return err
	}
	defer db.Close()

	// Execute the SQL query to save the user
	logger.Info("Executing query...")
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
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

	// Open a connection to the MySQL database
	logger.Info("Connecting to the database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to the database: ", err)
		return nil, err
	}
	defer db.Close()

	// Execute the SQL query to retrieve the user by username
	logger.Info("Executing query...")
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)

	// Scan the row data into a User struct
	user := &User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found
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

	// Open a connection to the MySQL database
	logger.Info("Connecting to the database...")
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		logger.Error("ERROR when connecting to the database: ", err)
		return err
	}
	defer db.Close()

	// Execute the SQL query to update the user
	logger.Info("Executing query...")
	_, err = db.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", u.Username, u.Password, u.ID)
	if err != nil {
		logger.Error("ERROR when updating user: ", err)
		return err
	}

	logger.Info("User updated successfully: ", u.Username)
	return nil
}
