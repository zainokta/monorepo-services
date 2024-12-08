package transaction

import (
	"context"
	"log"
	"time"
	"transaction/pkg/database"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	UUID      string    `json:"id" db:"uuid"`
	Status    string    `json:"status" db:"status"`
	UserID    string    `json:"userId" db:"user_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func StoreTransaction(ctx context.Context, req Transaction) error {
	query := "INSERT INTO transactions (status, user_id) VALUES (@status, @user_id)"
	args := pgx.NamedArgs{
		"status":  req.Status,
		"user_id": req.UserID,
	}

	_, err := database.DB.Exec(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
