package repository

import (
	"database/sql"
	"github.com/dian/bank-api/internal/model"
	"github.com/dian/bank-api/pkg/errors"
	"github.com/google/uuid"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func InsertUser(db *sql.DB, name, email string) error {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	return err
}

func GetAllUsers(db *sql.DB) ([]model.User, error) {

	rows, err := db.Query("SELECT id, name, email, created_at FROM users")

	if err == sql.ErrNoRows {
		return nil, errors.NewBusinessError("user not found")
	}

	if err != nil {
		return nil, errors.NewInternalError("failed to query user", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, errors.NewInternalError("failed to scan user", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT id, name, email, created_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.NewBusinessError("user not found")
	}
	if err != nil {
		return nil, errors.NewInternalError("failed to query user", err)
	}

	return &user, nil
}

func (r *UserRepo) SaveSession(userID int, token string, expiredAt time.Time) error {

	_, err := r.db.Exec(`
		INSERT INTO user_sessions (session_id, user_id, token, expired_at)
		VALUES ($1, $2, $3, $4)
	`, uuid.NewString(), userID, token, expiredAt)

	return err

}
