package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"userservice/internal/domain/models"
	"userservice/internal/storage"
	"userservice/pkg/logger/sl"

	_ "github.com/mattn/go-sqlite3"
)

const UsersTableName = "users"

type UserRepo struct {
	log  *slog.Logger
	Path string
}

func New(log *slog.Logger, path string) *UserRepo {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS ` + UsersTableName + ` (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			surname VARCHAR(50),
			name VARCHAR(50),
			age INT,
			CONSTRAINT age_lt_0 CHECK (age > 0)
		);

		CREATE INDEX IF NOT EXISTS user_id_idx ON ` + UsersTableName + `(id);
	`)

	if err != nil {
		panic("failed to create " + UsersTableName + " table:" + err.Error())
	}

	return &UserRepo{
		Path: path,
		log:  log,
	}
}

func (ur *UserRepo) Get(ctx context.Context) ([]models.User, error) {
	const op = "sqlite.get"
	db, err := sql.Open("sqlite3", ur.Path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to connect to database: %w", op, err)
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, `
		SELECT * FROM `+UsersTableName+`;
	`)
	if err != nil {
		ur.log.Error("error querying db", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var user models.User
	users := make([]models.User, 0, 5)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Surname, &user.Name, &user.Age)
		if err != nil {
			ur.log.Warn("cannot scan row", sl.Err(err))
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepo) GetById(ctx context.Context, id int) (models.User, error) {
	const op = "sqlite.getById"
	db, err := sql.Open("sqlite3", ur.Path)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: failed to connect to database: %w", op, err)
	}
	defer db.Close()

	row := db.QueryRowContext(ctx, `
		SELECT * FROM `+UsersTableName+`
		WHERE id=$1;
	`, id)

	var user models.User
	err = row.Scan(&user.Id, &user.Surname, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		ur.log.Warn("cannot scan row", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (ur *UserRepo) Insert(ctx context.Context, user models.User) error {
	const op = "sqlite.insert"
	db, err := sql.Open("sqlite3", ur.Path)
	if err != nil {
		return fmt.Errorf("%s: failed to connect to database: %w", op, err)
	}
	defer db.Close()

	var existingUser models.User
	err = db.QueryRowContext(ctx, `
		SELECT * FROM `+UsersTableName+`
		WHERE surname = $1 AND name = $2 AND age = $3
	`, user.Surname, user.Name, user.Age).Scan(&existingUser.Id, &existingUser.Surname, &existingUser.Name, &existingUser.Age)

	if err == nil {
		return fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
	} else if err != sql.ErrNoRows {
		ur.log.Error("error checking for existing user", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.ExecContext(ctx, `
		INSERT INTO `+UsersTableName+` (surname, name, age)
		VALUES ($1, $2, $3)
	`, user.Surname, user.Name, user.Age)
	if err != nil {
		ur.log.Error("error inserting row", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (ur *UserRepo) Update(ctx context.Context, id int, user models.User) error {
	const op = "sqlite.update"
	db, err := sql.Open("sqlite3", ur.Path)
	if err != nil {
		return fmt.Errorf("%s: failed to connect to database: %w", op, err)
	}
	defer db.Close()

	result, err := db.ExecContext(ctx, `
		UPDATE `+UsersTableName+` 
		SET surname=$1, name=$2, age=$3 
		WHERE id=$4
	`, user.Surname, user.Name, user.Age, id)

	if err != nil {
		ur.log.Error("error updating row", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ur.log.Error("error checking affected rows", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
	}

	return nil
}

func (ur *UserRepo) Delete(ctx context.Context, id int) error {
	const op = "sqlite.delete"
	db, err := sql.Open("sqlite3", ur.Path)
	if err != nil {
		return fmt.Errorf("%s: failed to connect to database: %w", op, err)
	}
	defer db.Close()

	result, err := db.ExecContext(ctx, `
		DELETE FROM `+UsersTableName+` WHERE id=$1;
	`, id)
	if err != nil {
		ur.log.Error("error deleting row", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ur.log.Error("error checking affected rows", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
	}

	return nil
}
