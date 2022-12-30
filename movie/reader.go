package movie

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (d *Dependency) GetMovie(ctx context.Context) ([]Movies, error) {

	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return []Movies{}, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil && !errors.Is(err, sql.ErrConnDone) {
			return
		}
	}()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true})
	if err != nil {

		return []Movies{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	rows, err := tx.QueryContext(
		ctx,
		"select * from movie",
	)

	if err != nil {
		if e := tx.Rollback(); e != nil {

			return []Movies{}, fmt.Errorf("failed to rollback transaction: %w", e)
		}

		return []Movies{}, fmt.Errorf("failed to query movies: %w", err)
	}

	var movies []Movies
	for rows.Next() {
		var movie Movies
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Rating,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {

			if e := tx.Rollback(); e != nil {
				return []Movies{}, fmt.Errorf("failed to rollback transaction: %w", e)
			}

			return []Movies{}, fmt.Errorf("failed to scan movie: %w", err)
		}

		movies = append(movies, movie)
	}

	if err := tx.Commit(); err != nil {
		return []Movies{}, fmt.Errorf("committing transaction: %w", err)
	}

	return movies, nil
}

func (d *Dependency) GetMovieById(ctx context.Context, movieId int) (Movies, error) {

	conn, err := d.DB.Conn(ctx)
	if err != nil {
		return Movies{}, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil && !errors.Is(err, sql.ErrConnDone) {
			return
		}
	}()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true})
	if err != nil {

		return Movies{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var movie Movies
	err = tx.QueryRowContext(
		ctx,
		"select * from movie where id = $1",
		movieId,
	).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Rating,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		if e := tx.Rollback(); e != nil {

			return Movies{}, fmt.Errorf("failed to rollback transaction: %w", e)
		}

		return Movies{}, fmt.Errorf("failed to query movies: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return Movies{}, fmt.Errorf("committing transaction: %w", err)
	}

	return movie, nil
}
