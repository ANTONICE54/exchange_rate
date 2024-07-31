package database

import (
	"database/sql"
	"rate/internal/models"
)

type IEmailRepo interface {
	Subscribe(email models.Email) error
	GetEmails() ([]*models.Email, error)
}

type EmailRepo struct {
	*sql.DB
}

func NewEmailRepo(db *sql.DB) *EmailRepo {
	return &EmailRepo{
		db,
	}
}

func (repo EmailRepo) Subscribe(email models.Email) error {
	query := "INSERT into emails(email) VALUES ($1)"

	_, err := repo.Exec(query, email.Email)

	if err != nil {
		return err
	}

	return nil
}

func (repo EmailRepo) GetEmails() ([]*models.Email, error) {
	query := "SELECT * FROM emails"

	rows, err := repo.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	emails := []*models.Email{}

	for rows.Next() {
		var curr models.Email
		err = rows.Scan(
			&curr.ID,
			&curr.Email,
			&curr.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		emails = append(emails, &curr)
	}

	return emails, nil
}
