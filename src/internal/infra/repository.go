package infra

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vnniciusg/src/internal/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db.DB,
	}
}

func (repo *UserRepositoryImpl) Create(user *domain.User) error {
	_, err := repo.db.Exec("INSERT INTO users (ID, email) VALUES ($1, $2)", user.ID, user.Name)
	if err != nil {
		log.Println("Error on insert user: ", err)
		return err
	}
	return nil
}

func (repo *UserRepositoryImpl) GetByID(id int) (*domain.User, error) {
	row := repo.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		log.Println("Error on get user: ", err)
		return nil, err
	}
	return user, nil
}
