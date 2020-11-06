package gofr

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getDB(t *testing.T) (*DB, sqlmock.Sqlmock){
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return &DB{mockDB}, mock
}

func TestDB_SelectSingleColumnFromIntToString(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		AddRow(2)
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	ids := make([]string, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []string{"1", "2"}, ids)
}

func TestDB_SelectSingleColumnFromStringToString(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1").
		AddRow("2")
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	ids := make([]string, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []string{"1", "2"}, ids)
}

func TestDB_SelectSingleColumnFromIntToInt(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		AddRow(2)
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	ids := make([]int, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []int{1, 2}, ids)
}

func TestDB_SelectSingleColumnFromIntToCustomInt(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		AddRow(2)
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	type CustomInt int
	ids := make([]CustomInt, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []CustomInt{1, 2}, ids)
}

func TestDB_SelectSingleColumnFromStringToCustomInt(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1").
		AddRow("2")
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	type CustomInt int
	ids := make([]CustomInt, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []CustomInt{1, 2}, ids)
}

func TestDB_SelectSingleColumnFromIntToCustomString(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		AddRow(2)
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	type CustomStr string
	ids := make([]CustomStr, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []CustomStr{"1", "2"}, ids)
}

func TestDB_SelectSingleColumnFromStringToCustomString(t *testing.T) {
	db, mock := getDB(t)
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1").
		AddRow("2")
	mock.ExpectQuery("^select id from users*").
		WillReturnRows(rows)

	type CustomStr string
	ids := make([]CustomStr, 0)
	db.Select(context.TODO(), &ids, "select id from users")
	assert.Equal(t, []CustomStr{"1", "2"}, ids)
}