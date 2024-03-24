package postgres

import (
	"database/sql"
	"errors"

	"github.com/evergreenies/go-api-tdd/pkg/common"
	"github.com/evergreenies/go-api-tdd/pkg/domain"
)

var (
	sqlCreateUser = `insert into users (name, email, password) 
	values ($1, $2, $3)
	returning id, name, email, password`
	sqlDeleteUserByID  = `delete from users where id = $1`
	sqlFindUserByEmail = `select id, name, email, password from users where email = $1`
	sqlFindUserByID    = `select id, name, email, password from users where id = $1`
)

func (p *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	password, err := common.PasswordHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	err = p.db.QueryRow(sqlCreateUser, user.Name, user.Email, user.Password).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *postgresStore) DeleteUserByID(id int64) error {
	_, err := p.db.Exec(sqlDeleteUserByID, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresStore) FindUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	err := p.db.QueryRow(sqlFindUserByEmail, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (p *postgresStore) FindUserByID(id int64) (*domain.User, error) {
	user := &domain.User{}

	err := p.db.QueryRow(sqlFindUserByID, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}
