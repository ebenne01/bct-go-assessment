package model

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const _uniqueKeyViolation = "23505"

var (
	db                   *sql.DB
	ErrDuplicateUserName = errors.New("Username already exists")
)

func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func CloseDB() {
	db.Close()
}

type User struct {
	ID         int     `json:"user_id"`
	UserName   *string `json:"user_name"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	Email      *string `json:"email"`
	Status     *string `json:"user_status"`
	Department *string `json:"department"`
}

func GetAllUsers() ([]User, error) {
	selectStmt := sq.
		Select("*").
		From("users")

	rows, err := selectStmt.RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Department)
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

func Create(u User) (User, error) {
	insert := sq.
		Insert("users").
		Columns("user_name", "first_name", "last_name", "email", "user_status", "department").
		Values(u.UserName, u.FirstName, u.LastName, u.Email, u.Status, u.Department).
		Suffix(`RETURNING user_id, user_name, first_name, last_name, email, user_status, department`).
		RunWith(db).
		PlaceholderFormat(sq.Dollar)

	var user User

	err := insert.QueryRow().Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Department)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == _uniqueKeyViolation {
			return User{}, ErrDuplicateUserName
		}
		return User{}, err
	}
	return user, nil
}

func Update(id int, u User) error {
	_, err := sq.Update("users").
		Set("first_name", u.FirstName).
		Set("last_name", u.LastName).
		Set("email", u.Email).
		Set("user_status", u.Status).
		Set("department", u.Department).
		Where(sq.Eq{"user_id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func Delete(id int) error {
	_, err := sq.Delete("users").
		Where(sq.Eq{"user_id": id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(db).
		Exec()

	if err != nil {
		return err
	}

	return nil
}
