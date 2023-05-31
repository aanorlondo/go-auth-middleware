package models

import (
	"database/sql"
	"errors"

	"app/config"
)

type User struct {
	ID       int
	Username string
	Password string
	// Add other fields as required
}

var databaseUrl = ""

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	databaseUrl = cfg.GetDatabaseURL()
}

func (u *User) Save() error {
	// TODO: Implement the logic to save the user to the MySQL database
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	// Execute the SQL query to save the user
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*User, error) {
	// TODO: Implement the logic to retrieve a user by username from the MySQL database
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Execute the SQL query to retrieve the user by username
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)

	// Scan the row data into a User struct
	user := &User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (u *User) Update() error {
	// TODO: Implement the logic to update the user in the MySQL database
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	// Execute the SQL query to update the user
	_, err = db.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", u.Username, u.Password, u.ID)
	if err != nil {
		return err
	}

	return nil
}
