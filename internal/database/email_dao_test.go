package database

/*
import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestEmailDaoSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)

	}

	defer db.Close()

	repo := NewEmailRepo(db)
	creationTime := time.Now()
	mock.ExpectPrepare("INSERT into emails").ExpectExec().WithArgs("bigrabbit54@carrot.com", creationTime).WillReturnResult(sqlmock.NewResult(1, 1))

	got, err := repo.Subscribe()

}
*/
