package persistence

import (
	"database/sql"
	"errors"
	"genesis/internal/models"
)

func (r *Repository) GetSubscriber(id int) (*models.Subscriber, error) {
	const query = `select id, email, password_hash
					from subscribers
					where id = $1`
	row := r.db.QueryRow(query, id)
	var subscriber models.Subscriber
	err := row.Scan(&subscriber.Id, &subscriber.Email, &subscriber.Password_hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("subscriber not found")
		}
		return nil, err
	}
	return &subscriber, nil
}
