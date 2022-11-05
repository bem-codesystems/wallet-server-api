package models

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

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

	res, err := u.DB.ExecContext(ctx, stmt, hash, false, name, email, "")
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `SELECT * FROM users WHERE id = ?`

	var user *User

	row := u.DB.QueryRowContext(ctx, stmt, userID)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.HasWallet,
		&user.WalletID)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (u *UserModel) GetAll() ([]*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	stmt := `SELECT * FROM users ORDER BY name LIMIT 20`

	var userList []*User

	rows, err := u.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &User{}

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.HasWallet,
			&user.WalletID)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
