package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"market/models"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) categoryRepo {
	return categoryRepo{db: db}
}

func (c categoryRepo) InsertCategory(category models.Category) error {
	id := uuid.New()
	if _, err := c.db.Exec(`insert into category (id,name) values ($1,$2)`, &id, &category.Name); err != nil {
		return err
	}
	return nil
}

func (c categoryRepo) GetByID(id uuid.UUID) (models.Category, error) {
	cat := models.Category{}

	if err := c.db.QueryRow(`select * from category where id = $1`, id).Scan(&cat.ID, &cat.Name, &cat.Created_at, &cat.Updated_at); err != nil {
		return models.Category{}, err
	}

	return cat, nil
}

func (c categoryRepo) GetList() ([]models.Category, error) {
	cats := []models.Category{}

	rows, err := c.db.Query(`select * from category`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cat := models.Category{}
		if err = rows.Scan(&cat.ID, &cat.Name, &cat.Created_at, &cat.Updated_at); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	return cats, nil
}
