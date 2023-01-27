package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user UserModel) error
}

type UserModel struct {
	id        UserID
	name      string
	email     string
	password  string
	createdAt string
	updatedAt string
	deletedAt string
}

func NewUser(id, name, email, password string) (UserModel, error) {

	idVO, err := NewUserID(id)
	if err != nil {
		return UserModel{}, err
	}

	return UserModel{
		id:        idVO,
		name:      name,
		email:     email,
		password:  password,
		createdAt: time.Now().Format("2006-01-02"),
		updatedAt: time.Now().Format("2006-01-02"),
	}, nil
}

// Value Objects
var ErrInvalidUserID = errors.New("invalid User ID")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

func (id UserID) String() string {
	return id.value
}

//  Getters

func (c UserModel) ID() UserID {
	return c.id
}

func (c UserModel) Name() string {
	return c.name
}

func (c UserModel) Email() string {
	return c.email
}

func (c UserModel) Password() string {
	return c.password
}

func (c UserModel) CreatedAt() string {
	return c.createdAt
}

func (c UserModel) UpdatedAt() string {
	return c.updatedAt
}
