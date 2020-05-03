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
	message := "hello world"
	lat := 100.00001
	long := 200.00002

	db, DBStore, mock := PrepareMockStore(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO notes").WithArgs(message).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`
		INSERT INTO breadcrumbs \(data_type, data_id, geog\) VALUES \( 'note', \(SELECT id from notes WHERE note=\$1\), 'SRID=4326;POINT\(100.000010 200.000020\)'\)
		`).WithArgs(message).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := DBStore.SaveNote(message, lat, long)
	require.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRetrieveNotes(t *testing.T) {
	radius := 123
	lat := 100.00001
	long := 200.00002

	db, DBStore, mock := PrepareMockStore(t)
	defer db.Close()
	mock.ExpectQuery(`
		SELECT n.note, ST_X\(b.geog::geometry\), ST_Y\(geog::geometry\) FROM breadcrumbs as b LEFT JOIN notes as n ON b.data_id = n.id WHERE ST_DWithin\(b.geog, ST_MakePoint\(100.000010, 200.000020\), \$1\)
		`).WithArgs(radius)
	err, _ := DBStore.RetrieveNotes(radius, lat, long)
	require.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveNotes(t *testing.T) {
	notes := []Note{
		{
			Message:   "one",
			Latitude:  100.00001,
			Longitude: 200.00002,
		},
		{
			Message:   "two",
			Latitude:  50.00001,
			Longitude: 90.00002,
		},
	}
	db, DBStore, mock := PrepareMockStore(t)
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO notes \(note\) VALUES \('one'\), \('two'\);`).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(
		`INSERT INTO breadcrumbs \(data_type, data_id, geog\) VALUES 
		\( 'note', \(SELECT id from notes WHERE note='one'\), 'SRID=4326;POINT\(100.000010 200.000020\)'\), 
		\( 'note', \(SELECT id from notes WHERE note='two'\), 'SRID=4326;POINT\(50.000010 90.000020\)'\);`).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	err := DBStore.SaveNotes(notes)
	require.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
