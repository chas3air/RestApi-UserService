package sqlite

import (
	"database/sql"
	"log"
	"userservice/internal/domain/models"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepo struct {
	Path string
}

func New(path string) *UserRepo {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			surname VARCHAR(50),
			name VARCHAR(50),
			age INT,
			CONSTRAINT age_lt_0 CHECK (age > 0)
		);

		CREATE INDEX IF NOT EXISTS user_id_idx ON users(id);
	`)

	if err != nil {
		log.Fatalln("failed to create users table:", err)
	}

	return &UserRepo{
		Path: path,
	}
}

func (ur *UserRepo) Get() ([]models.User, error) {
	panic("unimplement")
}

func (ur *UserRepo) GetById(id int) (models.User, error) {
	panic("unimplement")
}

func (ur *UserRepo) Insert(user models.User) error {
	panic("unimplement")
}

func (ur *UserRepo) Update(id int, user models.User) error {
	panic("unimplement")
}

func (ur *UserRepo) Delete(id int) error {
	panic("unimplement")
}
