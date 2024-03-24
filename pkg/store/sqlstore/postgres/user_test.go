package postgres

import (
	"testing"

	"github.com/evergreenies/go-api-tdd/pkg/domain"
)

var (
	oldSqlCreateUser   = sqlCreateUser
	oldDeleteUserByID  = sqlDeleteUserByID
	oldFindUserByEmail = sqlFindUserByEmail
	oldSqlFindUserByID = sqlFindUserByID
)

func TestCreateUser(t *testing.T) {
	pStore := NewPostgresStore(testDB)
	oldPassword := "password"

	user := &domain.User{
		Email:    "test@test.com",
		Password: "password",
		Name:     "Jon Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Error("ID must not be zero")
	}

	if user.Name != createdUser.Name {
		t.Errorf("expected %q; got %q\n", user.Name, createdUser.Name)
	}

	if user.Password == oldPassword {
		t.Error("password not hashed.")
	}

	err = pStore.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expected nil error, got %q\n", err)
	}

	sqlCreateUser = "invalid"
	createdUser, err = pStore.CreateUser(user)
	if err == nil {
		t.Error("expected error as error should not be nil")
	}
	sqlCreateUser = oldSqlCreateUser

	sqlDeleteUserByID = "invalid"
	err = pStore.DeleteUserByID(0)
	if err == nil {
		t.Error("expected non nil error, got an error")
	}
	sqlDeleteUserByID = oldDeleteUserByID
}

func TestFindUserByEmail(t *testing.T) {
	pStore := NewPostgresStore(testDB)

	user := &domain.User{
		Email:    "test1@test.com",
		Password: "password",
		Name:     "Jon Doe",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	userByEmail, err := pStore.FindUserByEmail(createdUser.Email)
	if err != nil {
		t.Fatal(err)
	}

	if user.Email != userByEmail.Email {
		t.Errorf("We expect user by email %q; got %q", user.Email, userByEmail.Email)
	}

	err = pStore.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expected nil error, got %q\n", err)
	}

	userByEmail, err = pStore.FindUserByEmail(createdUser.Email)
	if err == nil {
		t.Errorf("We do not expect user email; got %q\n", userByEmail.Email)
	}

	sqlFindUserByEmail = "invalid"
	userByEmail, err = pStore.FindUserByEmail("invalid@mail.com")
	if err == nil {
		t.Error("We expect to raise erro as SQL query is wrong to find user by email.")
	}
	sqlFindUserByEmail = oldFindUserByEmail
}

func TestFindUserByID(t *testing.T) {
	pStore := NewPostgresStore(testDB)

	user := &domain.User{
		Name:     "John Dow",
		Email:    "tempmail1@mail.com",
		Password: "password",
	}

	createdUser, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	userByID, err := pStore.FindUserByID(createdUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != userByID.ID {
		t.Errorf("We expect user by id %q; got %q", user.ID, userByID.ID)
	}

	err = pStore.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expected nil error, got %q\n", err)
	}

	userByID, err = pStore.FindUserByID(createdUser.ID)
	if err == nil {
		t.Errorf("We do not expect user id; got %q\n", userByID.ID)
	}

	sqlFindUserByID = "invalid"
	userByID, err = pStore.FindUserByID(0)
	if err == nil {
		t.Error("We expect to raise erro as SQL query is wrong to find user by email.")
	}
	sqlFindUserByEmail = oldFindUserByEmail
}

func TestDeleteAllUsers(t *testing.T) {
	pStore := NewPostgresStore(testDB)

	user := &domain.User{
		Email:    "test@test@mail.com",
		Password: "password",
		Name:     "Jon Doe",
	}

	_, err := pStore.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	err = pStore.DeleteAllUsers()
	if err != nil {
		t.Fatal(err)
	}
}
