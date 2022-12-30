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
		`INSERT INTO movie (title,description,rating,image) VALUES ($1,$2,$3,$4)`,
		newMovie.Title,
		newMovie.Description,
		newMovie.Rating,
		newMovie.Image,
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

func (d *Dependency) UpdateMovie(ctx context.Context, id int, newMovie *Movies) error {
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
		`UPDATE movie SET title = $1,description = $2 ,rating = $3, image = $4, updated_at = now() WHERE id = $5`,
		newMovie.Title,
		newMovie.Description,
		newMovie.Rating,
		newMovie.Image,
		id,
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

func (d *Dependency) DeleteMovie(ctx context.Context, id int) error {
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
		`DELETE from movie where id = $1`,
		id,
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
