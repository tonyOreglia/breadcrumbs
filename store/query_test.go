package store

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func PrepareMockStore(t *testing.T) (*sqlx.DB, *Store, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := sqlx.NewDb(mockDB, "sqlmock")

	dbStore := &Store{
		DB: db,
	}
	return db, dbStore, mock
}

func TestSaveNote(t *testing.T) {
	testNote := "hello world"
	testLat := 100.00001
	testLon := 200.00002
	db, DBStore, mock := PrepareMockStore(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO notes").WithArgs(testNote).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO breadcrumbs").WithArgs(testNote).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := DBStore.SaveNote(testNote, testLat, testLon)
	require.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
