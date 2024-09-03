package database

import (
	"rate/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	defer db.Close()

	emailExpected := models.Email{
		ID:        1,
		Email:     "bigrabbit54@carrot.com",
		CreatedAt: time.Now(),
	}

	repo := NewSubscriptionRepo(db)
	rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).AddRow(emailExpected.ID, emailExpected.Email, emailExpected.CreatedAt)
	mock.ExpectQuery("SELECT \\* FROM emails WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

	gotEmailObj, err := repo.GetByID(1)
	require.NoError(t, err)
	require.NotNil(t, gotEmailObj)
	require.Equal(t, emailExpected.ID, gotEmailObj.ID)
	require.Equal(t, emailExpected.Email, gotEmailObj.Email)
	require.WithinDuration(t, emailExpected.CreatedAt, gotEmailObj.CreatedAt, time.Second)
}

func TestSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	defer db.Close()
	emailExpected := models.Email{
		ID:        1,
		Email:     "bigrabbit54@carrot.com",
		CreatedAt: time.Now(),
	}

	repo := NewSubscriptionRepo(db)
	rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).AddRow(emailExpected.ID, emailExpected.Email, emailExpected.CreatedAt)
	mock.ExpectQuery("INSERT into emails\\(email\\) VALUES \\(\\$1\\) RETURNING id, email, created_at;").WithArgs(emailExpected.Email).WillReturnRows(rows)

	got, err := repo.Create(emailExpected)

	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, emailExpected.ID, got.ID)
	require.Equal(t, emailExpected.Email, got.Email)
	require.WithinDuration(t, emailExpected.CreatedAt, got.CreatedAt, time.Second)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

}

func TestListEmails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}
	repo := NewSubscriptionRepo(db)
	defer db.Close()

	expectedList := []*models.Email{
		{
			ID:        1,
			Email:     "bigrabbit54@carrot.com",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Email:     "bigrabbit54@carrot.com",
			CreatedAt: time.Now(),
		},
		{
			ID:        3,
			Email:     "bigrabbit54@carrot.com",
			CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "email", "created_at"}).AddRow(
		expectedList[0].ID, expectedList[0].Email, expectedList[0].CreatedAt,
	).AddRow(
		expectedList[1].ID, expectedList[1].Email, expectedList[1].CreatedAt,
	).AddRow(expectedList[2].ID, expectedList[2].Email, expectedList[2].CreatedAt)

	mock.ExpectQuery("SELECT \\* FROM emails").WillReturnRows(rows)
	gotList, err := repo.List()

	require.NoError(t, err)
	require.NotNil(t, gotList)
	require.ElementsMatch(t, expectedList, gotList)
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

}
