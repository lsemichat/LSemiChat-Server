package repository

import (
	"app/api/domain/entity"
	"app/api/domain/repository"
	"app/api/infrastructure/database"

	"github.com/pkg/errors"
)

type tagRepository struct {
	sqlHandler database.SQLHandler
}

func NewTagRepository(sh database.SQLHandler) repository.TagRepository {
	return &tagRepository{
		sqlHandler: sh,
	}
}

func (tr *tagRepository) Create(tag *entity.Tag) error {
	_, err := tr.sqlHandler.Exec(`
		INSERT INTO tags(id, tag, category_id)
		VALUES (?, ?, ?)
	`,
		tag.ID,
		tag.Tag,
		tag.Category.ID,
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert db")
	}
	return nil
}

func (tr *tagRepository) FindAll() ([]*entity.Tag, error) {
	rows, err := tr.sqlHandler.Query(`
		SELECT id, tag, category_id
		FROM tags
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	var tags []*entity.Tag
	for rows.Next() {
		var tag entity.Tag
		var category entity.Category
		if err = rows.Scan(&tag.ID, &tag.Tag, &category.ID); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan")
		}
		tag.Category = &category
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (tr *tagRepository) FindByID(id string) (*entity.Tag, error) {
	row := tr.sqlHandler.QueryRow(`
		SELECT id, tag, category_id
		FROM tags
		WHERE id=?
	`, id)
	var tag entity.Tag
	var category entity.Category
	if err := row.Scan(&tag.ID, &tag.Tag, &category.ID); err != nil {
		return nil, errors.Wrap(err, "failed to scan")
	}
	tag.Category = &category
	return &tag, nil
}

func (tr *tagRepository) FindByCategoryID(id string) ([]*entity.Tag, error) {
	rows, err := tr.sqlHandler.Query(`
		SELECT id, tag, category_id
		FROM tags
		WHERE category_id=?
	`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	var tags []*entity.Tag
	for rows.Next() {
		var tag entity.Tag
		var category entity.Category
		if err = rows.Scan(&tag.ID, &tag.Tag, &category.ID); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan")
		}
		tag.Category = &category
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (tr *tagRepository) FindByUserUUID(id string) ([]*entity.Tag, error) {
	rows, err := tr.sqlHandler.Query(`
		SELECT t.id, t.tag, t.category_id
		FROM users_tags as r
		INNER JOIN tags as t
		ON t.id = r.tag_id
		WHERE user_id=?
	`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	var tags []*entity.Tag
	for rows.Next() {
		var tag entity.Tag
		var category entity.Category
		if err = rows.Scan(&tag.ID, &tag.Tag, &category.ID); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan")
		}
		tag.Category = &category
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (tr *tagRepository) FindByThreadID(id string) ([]*entity.Tag, error) {
	rows, err := tr.sqlHandler.Query(`
		SELECT t.id, t.tag, t.category_id
		FROM threads_tags as r
		INNER JOIN tags as t
		ON t.id = r.tag_id
		WHERE thread_id=?
	`, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	var tags []*entity.Tag
	for rows.Next() {
		var tag entity.Tag
		var category entity.Category
		if err = rows.Scan(&tag.ID, &tag.Tag, &category.ID); err != nil {
			if rows.CheckNoRows(err) {
				return nil, nil
			}
			return nil, errors.Wrap(err, "failed to scan")
		}
		tag.Category = &category
		tags = append(tags, &tag)
	}
	return tags, nil
}
