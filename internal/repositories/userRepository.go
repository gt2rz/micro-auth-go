package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gt2rz/micro-auth/internal/models"
	"github.com/gt2rz/micro-auth/internal/utils"
)

// errUserNotSaved is the error returned when a user is not saved
var errUserNotSaved = errors.New("User not saved")

// errUserNotFound is the error returned when a user is not found
var errUserNotFound = errors.New("User not found")

// errGenerateRandomString is the error returned when a random string cannot be generated
var errGenerateRandomString = errors.New("Error generating random string")

// errResetTokenNotSaved is the error returned when a reset token is not saved
var errResetTokenNotSaved = errors.New("Reset token not saved")

// UserRepository is the interface for the user repository
type UserRepository interface {
	SaveUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, id string) (*models.User, error)
	GenerateResetToken(ctx context.Context, id string) (string, error)
	GetUserByResetToken(ctx context.Context, resetToken string) (*models.User, error)
	UpdatePassword(ctx context.Context, id string, password string) error
}

// UserRepositoryImpl is the implementation of the UserRepository interface
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) (*UserRepositoryImpl, error) {
	return &UserRepositoryImpl{db}, nil
}

// SaveUser saves a user to the database
func (r *UserRepositoryImpl) SaveUser(ctx context.Context, user models.User) error {

	result, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password, first_name, last_name, phone, verified, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?,?, ?, ?)", user.Id, user.Email, user.Password, user.Firstname, user.Lastname, user.Phone, user.Verified, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return errUserNotSaved
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errUserNotSaved
	}

	if rowsAffected != 1 {
		return errUserNotSaved
	}
	return nil
}

// GetUserByEmail gets a user from the database by email
func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user = models.User{}

	query := r.db.QueryRowContext(ctx, "SELECT id, email, password, first_name, last_name, phone, verified, created_at, updated_at  FROM users WHERE email=?", email)
	err := query.Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Phone, &user.Verified, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errUserNotFound
	}

	if err != nil {
		return nil, errUserNotFound
	}

	return &user, nil

}

// GenerateResetToken generates a reset token for a user
func (r *UserRepositoryImpl) GenerateResetToken(ctx context.Context, id string) (string, error) {
	resetTokenAt := time.Now().Add(2 * time.Minute)
	resetToken, err := utils.GenerateRandomString(64)
	if err != nil {
		return "", errGenerateRandomString
	}

	result, err := r.db.ExecContext(ctx, "UPDATE users SET password_reset_token=?, password_reset_token_at=? WHERE id=?", resetToken, resetTokenAt, id)
	if err != nil {
		return "", errResetTokenNotSaved
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", errResetTokenNotSaved
	}

	if rowsAffected != 1 {
		return "", errResetTokenNotSaved
	}

	return resetToken, nil
}

// GetUserByResetToken gets a user from the database by reset token
func (r *UserRepositoryImpl) GetUserByResetToken(ctx context.Context, resetToken string) (*models.User, error) {
	var user = models.User{}

	query := r.db.QueryRowContext(ctx, "SELECT id, email, first_name, password_reset_token_at  FROM users WHERE password_reset_token=?", resetToken)
	err := query.Scan(&user.Id, &user.Email, &user.Firstname, &user.PasswordResetTokenAt)
	if err == sql.ErrNoRows {
		return nil, errUserNotFound
	}

	if err != nil {
		return nil, errUserNotFound
	}

	return &user, nil
}

// UpdatePassword updates a user's password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, id string, password string) error {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET password=? WHERE id=?", password, id)

	if err != nil {
		return utils.ErrAnErrorOccurred
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrAnErrorOccurred
	}

	if rowsAffected != 1 {
		return errUserNotSaved
	}
	return nil
}
