package repository

import (
	"app/api/application/repository"
	"app/api/domain/entity"
	"app/api/infrastructure/database"
	"time"

	"github.com/pkg/errors"
)

type userRepository struct {
	sqlHandler database.SQLHandler
}

func NewUserRepository(sqlHandler database.SQLHandler) repository.UserRepository {
	return &userRepository{
		sqlHandler: sqlHandler,
	}
}

func (repo *userRepository) Create(user *entity.User) error {
	_, err := repo.sqlHandler.Exec(`
		INSERT INTO users(id, user_id, name, image, profile, is_admin, mail, login_at, created_at, updated_at, password)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		user.ID,
		user.UserID,
		user.Name,
		user.Image,
		user.Profile,
		user.IsAdmin,
		user.Mail,
		user.LoginAt,
		user.CreatedAt,
		user.UpdatedAt,
		user.Password,
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert user")
	}
	return nil
}

func (repo *userRepository) UpdateProfile(user *entity.User) error {
	_, err := repo.sqlHandler.Exec(`
		UPDATE users
		SET name=?, mail=?, image=?, profile=?, updated_at=?
		WHERE id=?;
	`,
		user.Name,
		user.Mail,
		user.Image,
		user.Profile,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (repo *userRepository) UpdateUserID(id, userID string, updatedAt *time.Time) error {
	_, err := repo.sqlHandler.Exec(`
		UPDATE users
		SET user_id=?, updated_at=?
		WHERE id=?;
	`,
		userID,
		updatedAt,
		id,
	)
	if err != nil {
		return errors.Wrap(err, "failed to update userID")
	}
	return nil
}

func (repo *userRepository) UpdatePassword(id, password string, updatedAt *time.Time) error {
	_, err := repo.sqlHandler.Exec(`
		UPDATE users
		SET password=?, updated_at=?
		WHERE id=?;
	`,
		password,
		updatedAt,
		id,
	)
	if err != nil {
		return errors.Wrap(err, "failed to update password")
	}
	return nil
}

func (repo *userRepository) FindByID(id string) (*entity.User, error) {
	row := repo.sqlHandler.QueryRow(`
		SELECT id, user_id, name, image, profile, is_admin, mail, login_at, created_at, updated_at, password
		FROM users
		WHERE id=?
	`, id)
	var user entity.User
	if err := row.Scan(&user.ID, &user.UserID, &user.Name, &user.Image, &user.Profile, &user.IsAdmin, &user.Mail, &user.LoginAt, &user.CreatedAt, &user.UpdatedAt, &user.Password); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}
	return &user, nil
}

func (repo *userRepository) FindByUserID(userID string) (*entity.User, error) {
	row := repo.sqlHandler.QueryRow(`
		SELECT id, user_id, name, image, profile, is_admin, mail, login_at, created_at, updated_at, password
		FROM users
		WHERE user_id=?
	`, userID)
	var user entity.User
	if err := row.Scan(&user.ID, &user.UserID, &user.Name, &user.Image, &user.Profile, &user.IsAdmin, &user.Mail, &user.LoginAt, &user.CreatedAt, &user.UpdatedAt, &user.Password); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}
	return &user, nil
}

func (repo *userRepository) FindByMail(mail string) (*entity.User, error) {
	row := repo.sqlHandler.QueryRow(`
		SELECT id, user_id, name, image, profile, is_admin, mail, login_at, created_at, updated_at, password
		FROM users
		WHERE mail=?
	`, mail)
	var user entity.User
	if err := row.Scan(&user.ID, &user.UserID, &user.Name, &user.Image, &user.Profile, &user.IsAdmin, &user.Mail, &user.LoginAt, &user.CreatedAt, &user.UpdatedAt, &user.Password); err != nil {
		return nil, errors.Wrap(err, "failed to scan user")
	}
	return &user, nil
}

func (repo *userRepository) FindAll() ([]*entity.User, error) {
	rows, err := repo.sqlHandler.Query(`
		SELECT id, user_id, name, image, profile, is_admin, mail, login_at, created_at, updated_at, password
		FROM users
	`)
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.ID, &user.UserID, &user.Name, &user.Image, &user.Profile, &user.IsAdmin, &user.Mail, &user.LoginAt, &user.CreatedAt, &user.UpdatedAt, &user.Password); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *userRepository) DeleteByID(id string) error {
	_, err := repo.sqlHandler.Exec(`
		DELETE FROM users WHERE id=?
	`, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete from db")
	}
	return nil
}

func (repo *userRepository) FindFollows(id string) ([]*entity.User, error) {
	rows, err := repo.sqlHandler.Query(`
		SELECT u.id, u.user_id, u.name, u.mail, u.image, u.profile, u.created_at, u.updated_at, u.login_at
		FROM users_followers as f
		INNER JOIN users as u
		ON u.id = f.followed_user_id
		WHERE f.user_id=?
	`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get follows")
	}
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.ID, &user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile, &user.CreatedAt, &user.UpdatedAt, &user.LoginAt); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *userRepository) FindFollowers(id string) ([]*entity.User, error) {
	rows, err := repo.sqlHandler.Query(`
		SELECT u.id, u.user_id, u.name, u.mail, u.image, u.profile, u.created_at, u.updated_at, u.login_at
		FROM users_followers as f
		INNER JOIN users as u
		ON u.id = f.followed_user_id
		WHERE f.followed_user_id=?
	`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get follows")
	}
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(&user.ID, &user.UserID, &user.Name, &user.Mail, &user.Image, &user.Profile, &user.CreatedAt, &user.UpdatedAt, &user.LoginAt); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
