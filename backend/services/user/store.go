package user

import (
	"database/sql"
	"fmt"

	"backend/types"
)

type SMS struct {
	db *sql.DB
}

func NewStore(
	db *sql.DB,
) *SMS {

	return &SMS{
		db: db,
	}
}

// =========================
// CREATE USER
// =========================

func (s *SMS) CreateUser(
	user types.User,
) error {

	_, err := s.db.Exec(
		`
		INSERT INTO users
		(name, email, password, role)
		VALUES (?, ?, ?, ?)
		`,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	)

	if err != nil {
		return err
	}

	return nil
}

// =========================
// GET USER BY EMAIL
// =========================

func (s *SMS) GetUserByEmail(
	email string,
) (*types.User, error) {

	rows, err := s.db.Query(
		"SELECT * FROM users WHERE email = ?",
		email,
	)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {

		u, err = scanRowsIntoUser(
			rows,
		)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {

		return nil, fmt.Errorf(
			"user not found",
		)
	}

	return u, nil
}

// =========================
// GET USER BY ID
// =========================

func (s *SMS) GetUserByID(
	id int,
) (*types.User, error) {

	rows, err := s.db.Query(
		"SELECT * FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {

		u, err = scanRowsIntoUser(
			rows,
		)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {

		return nil, fmt.Errorf(
			"user not found",
		)
	}

	return u, nil
}

// =========================
// GET ALL USERS
// =========================

func (s *SMS) GetUsers() (
	[]types.User,
	error,
) {

	rows, err := s.db.Query(
		"SELECT * FROM users",
	)

	if err != nil {
		return nil, err
	}

	var users []types.User

	for rows.Next() {

		user, err := scanRowsIntoUser(
			rows,
		)

		if err != nil {
			return nil, err
		}

		users = append(
			users,
			*user,
		)
	}

	return users, nil
}

// =========================
// DELETE USER
// =========================

func (s *SMS) DeleteUser(
	id int,
) error {

	_, err := s.db.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

// =========================
// SCAN USER
// =========================

func scanRowsIntoUser(
	rows *sql.Rows,
) (*types.User, error) {

	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}