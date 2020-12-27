package repository

import (
	"app/api/application/repository"
	"app/api/domain/entity"
	"app/api/infrastructure/database"

	"github.com/pkg/errors"
)

type categoryRepository struct {
	sqlHandler database.SQLHandler
}

func NewCategoryRepository(sh database.SQLHandler) repository.CategoryRepository {
	return &categoryRepository{
		sqlHandler: sh,
	}
}

func (cr *categoryRepository) GetAll() ([]*entity.Category, error) {
	rows, err := cr.sqlHandler.Query(`
		SELECT id, category
		FROM categories
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select db")
	}
	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err = rows.Scan(&category.ID, &category.Category); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan category")
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cr *categoryRepository) GetByID(id string) (*entity.Category, error) {
	row := cr.sqlHandler.QueryRow(`
		SELECT id, category
		FROM categories
		WHERE id=?
	`, id)
	var category entity.Category
	if err := row.Scan(&category.ID, &category.Category); err != nil {
		return nil, errors.Wrap(err, "failed to select from db")
	}
	return &category, nil
}

func (cr *categoryRepository) Create(category *entity.Category) error {
	_, err := cr.sqlHandler.Exec(`
		INSERT INTO categories(id, category)
		VALUES (?, ?)
	`,
		category.ID,
		category.Category,
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert db")
	}
	return nil
}

func (cr *categoryRepository) Update(category *entity.Category) error {
	_, err := cr.sqlHandler.Exec(`
		UPDATE categories
		SET category=?
	`, category.Category)
	if err != nil {
		return errors.Wrap(err, "failed to update category")
	}
	return nil
}

func (cr *categoryRepository) Delete(id string) error {
	_, err := cr.sqlHandler.Exec(`
		DELETE FROM categories
		WHERE id=?
	`, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete category")
	}
	return nil
}
