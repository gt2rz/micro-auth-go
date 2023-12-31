package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/models"
	"github.com/gt2rz/micro-auth/internal/utils"
)

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

	sql := `
	INSERT INTO users (
		id,
		email, 
		password, 
		firstname, 
		lastname, 
		phone 
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	result, err := r.db.ExecContext(ctx, sql,
		user.Id,
		user.Email,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Phone,
	)

	fmt.Println(err)
	if err != nil {
		return constants.ErrUserNotSaved
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return constants.ErrUserNotSaved
	}

	if rowsAffected != 1 {
		return constants.ErrUserNotSaved
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user = models.User{}

	query := `
		SELECT 
			id, 
			email, 
			password, 
			firstname, 
			lastname, 
			phone, 
			verified, 
			status, 
			created_at, 
			updated_at 
		FROM users 
		WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname,
		&user.Phone,
		&user.Verified,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, constants.ErrUserNotFound
	}

	if err != nil {
		return nil, constants.ErrGettingUser
	}

	return &user, nil
}

// GenerateResetToken generates a reset token for a user
func (r *UserRepositoryImpl) GenerateResetToken(ctx context.Context, id string) (string, error) {
	resetTokenAt := time.Now().Add(2 * time.Minute)
	resetToken, err := utils.GenerateRandomString(64)
	if err != nil {
		return "", constants.ErrGenerateRandomString
	}

	result, err := r.db.ExecContext(ctx, "UPDATE users SET password_reset_token=$1, password_reset_token_at=$2 WHERE id=$3", resetToken, resetTokenAt, id)
	if err != nil {
		return "", constants.ErrResetTokenNotSaved
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", constants.ErrResetTokenNotSaved
	}

	if rowsAffected != 1 {
		return "", constants.ErrResetTokenNotSaved
	}

	return resetToken, nil
}

// GetUserByResetToken gets a user from the database by reset token
func (r *UserRepositoryImpl) GetUserByResetToken(ctx context.Context, resetToken string) (*models.User, error) {
	var user = models.User{}

	query := r.db.QueryRowContext(ctx, "SELECT id, email, firstname, password_reset_token_at  FROM users WHERE password_reset_token=$1", resetToken)
	err := query.Scan(&user.Id, &user.Email, &user.Firstname, &user.PasswordResetTokenAt)
	if err == sql.ErrNoRows {
		return nil, constants.ErrUserNotFound
	}

	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	return &user, nil
}

// UpdatePassword updates a user's password
func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, id string, password string) error {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET password=$1 WHERE id=$2", password, id)

	if err != nil {
		return utils.ErrAnErrorOccurred
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrAnErrorOccurred
	}

	if rowsAffected != 1 {
		return constants.ErrUserNotSaved
	}
	return nil
}
