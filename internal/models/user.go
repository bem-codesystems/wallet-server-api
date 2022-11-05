package models

import (
	"database/sql"
	"time"
	"wallet-server/helpers"
)

type User struct {
	ID        string    `json:"id"`
	HasWallet bool      `json:"has_wallet"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	WalletID  string    `json:"wallet_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Create(name string, email string) (int, error) {
	stmt := `INSERT INTO users (
                   id,
                   has_wallet,
                   name,
                   email,
                   wallet_id,
                   created_at,
                   updated_at) VALUES (?,?,?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP())`

	hash, err := helpers.CreateRandomID(50)
	if err != nil {
		return 0, err
	}

	res, err := u.DB.Exec(stmt, hash, false, name, email, "")
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

func (u *UserModel) Get(userID string) (*User, error) {
	stmt := `SELECT * FROM users WHERE id = ?`

	var user *User

	row := u.DB.QueryRow(stmt, userID)

	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.HasWallet, &user.WalletID)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
