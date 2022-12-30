package movie

import (
	"context"
	"database/sql"
	"fmt"
)

func (d *Dependency) CreateMovie(ctx context.Context, newMovie *Movies) error {
	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return fmt.Errorf("acquiring connection pool: %w", err)
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO movie (title,description,rating,image, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`,
		newMovie.Title,
		newMovie.Description,
		newMovie.Rating,
		newMovie.Image,
		newMovie.CreatedAt,
		newMovie.UpdatedAt,
	)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return fmt.Errorf("rolling back transaction: %w", err)
		}

		return fmt.Errorf("executing update transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return fmt.Errorf("rolling back transaction: %w", e)
		}

		return fmt.Errorf("commiting transaction: %w", err)
	}

	return nil
}
