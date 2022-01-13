package repository

import (
	"database/sql"
	"errors"
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/model/domain"
	"context"
)

type CategoryRepositoryImplementation struct {

}

func (repository *CategoryRepositoryImplementation) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) values (?)"
	result, err:= tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	helper.PanicError(err)
	category.Id=int(id)
	return category
}

func (repository *CategoryRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id=?"
	_, err := tx.ExecContext(ctx,SQL,category.Name,category.Id)
	helper.PanicError(err)
	return category
}

func (repository *CategoryRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id=?"
	_, err := tx.ExecContext(ctx,SQL,category.Name,category.Id)
	helper.PanicError(err)
}

func (repository *CategoryRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category,error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx,SQL,categoryId)
	helper.PanicError(err)
	category := domain.Category{}
	if rows.Next(){
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		return category, nil
	}else{
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	var categories []domain.Category
	for rows.Next(){
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}
	return categories
}
