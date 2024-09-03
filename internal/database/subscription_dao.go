package database

import (
	"database/sql"
	"log"
	"rate/internal/models"
)

type SubscriptionRepo struct {
	*sql.DB
}

func NewSubscriptionRepo(db *sql.DB) *SubscriptionRepo {
	return &SubscriptionRepo{
		db,
	}
}

func (repo SubscriptionRepo) Create(email models.Email) (*models.Email, error) {
	query := "INSERT into emails(email) VALUES ($1) RETURNING id, email, created_at;"

	row := repo.QueryRow(query, email.Email)
	var res models.Email

	err := row.Scan(
		&res.ID,
		&res.Email,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	log.Println(res)
	return &res, nil
}

func (repo SubscriptionRepo) List() ([]*models.Email, error) {
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

func (repo SubscriptionRepo) GetByID(emailID uint) (*models.Email, error) {
	query := "SELECT * FROM emails WHERE id = $1"
	row := repo.QueryRow(query, emailID)
	var email models.Email

	err := row.Scan(
		&email.ID,
		&email.Email,
		&email.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &email, nil

}
