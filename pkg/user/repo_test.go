package user

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	userID := "1"

	// good query
	rows := sqlmock.
		NewRows([]string{"username", "admin", "id", "passwordHash", "token"})
	user := &User{
		Username:     "Coolguy",
		Admin:        false,
		ID:           userID,
		PasswordHash: "CoolHash",
		Token:        "CoolToken",
	}
	expect := []*User{
		user,
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, item.PasswordHash, item.Token)
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash, token").
		WithArgs(userID).
		WillReturnRows(rows)

	repo := &Repo{
		DB: db,
	}
	item, err := repo.GetByID(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query db error
	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash, token").
		WithArgs(userID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByID(userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	userID := "1"
	login := "Coolguy"

	// good query
	rows := sqlmock.
		NewRows([]string{"username", "admin", "id", "passwordHash"})
	user := &User{
		Username:     login,
		Admin:        false,
		ID:           userID,
		PasswordHash: "CoolHash",
	}
	expect := []*User{
		user,
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, item.PasswordHash)
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnRows(rows)

	repo := &Repo{
		DB: db,
	}
	item, err := repo.GetByUserName(login)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query db error
	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByUserName(login)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAuthorize(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	userID := "1"
	login := "Coolguy"
	password := "CoolPassword"
	hash := getSHA256(password)

	// good query
	rows := sqlmock.
		NewRows([]string{"username", "admin", "id", "passwordHash"})
	user := &User{
		Username:     login,
		Admin:        false,
		ID:           userID,
		PasswordHash: hash,
	}
	expect := []*User{
		user,
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, item.PasswordHash)
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnRows(rows)

	repo := &Repo{
		DB: db,
	}
	item, err := repo.Authorize(login, password)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query db error
	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnError(fmt.Errorf("db_error"))

	item, err = repo.Authorize(login, password)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// incorrect hash error
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, "incorrectHash")
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnRows(rows)

	item, err = repo.Authorize(login, password)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	userID := "1"
	login := "Coolguy"
	password := "CoolPassword"
	hash := getSHA256(password)

	// good query
	user := &User{
		Username:     login,
		Admin:        false,
		ID:           userID,
		PasswordHash: hash,
	}

	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			user.PasswordHash,
			user.Token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &Repo{
		DB: db,
	}
	id, err := repo.add(user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != userID {
		t.Errorf("bad id: want %s, have %v", userID, id)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// db error on insert
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			user.PasswordHash,
			user.Token).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.add(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// erroneous insert exec result
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			user.PasswordHash,
			user.Token).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.add(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("can't create mock: %s", err)
	}
	defer db.Close()

	userID := "1"
	login := "Coolguy"
	password := "CoolPassword"
	hash := getSHA256(password)

	// good query
	rows := sqlmock.
		NewRows([]string{"username", "admin", "id", "passwordHash", "token"})
	user := &User{
		Username: login,
		Admin:    false,
		ID:       userID,
		Password: password,
	}
	expect := []*User{
		user,
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, hash, item.Token)
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnError(fmt.Errorf("user_not_found"))

	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			hash,
			user.Token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash, token").
		WithArgs(userID).
		WillReturnRows(rows)

	repo := &Repo{
		DB: db,
	}

	item, err := repo.Register(user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// insert user error
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			user.PasswordHash,
			user.Token).
		WillReturnError(fmt.Errorf("db_error"))

	item, err = repo.Register(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// GetByID error
	mock.
		ExpectExec("INSERT INTO users").
		WithArgs(user.Username,
			user.Admin,
			hash,
			user.Token).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash, token").
		WithArgs(userID).
		WillReturnError(fmt.Errorf("db_error"))

	item, err = repo.Register(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// user already exists error
	rows = sqlmock.
		NewRows([]string{"username", "admin", "id", "passwordHash"})
	user = &User{
		Username:     login,
		Admin:        false,
		ID:           userID,
		PasswordHash: "CoolHash",
	}
	expect = []*User{
		user,
	}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Admin, item.PasswordHash)
	}

	mock.
		ExpectQuery("SELECT id, username, admin, passwordHash").
		WithArgs(login).
		WillReturnRows(rows)

	item, err = repo.Register(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
