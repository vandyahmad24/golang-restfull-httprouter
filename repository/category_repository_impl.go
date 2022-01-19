package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restapi-httprouter/helper"
	"golang-restapi-httprouter/model/domain"
	"log"
)

type CategoryRepositoryImplementation struct {

}

func NewCategoryRepository() *CategoryRepositoryImplementation {
	return &CategoryRepositoryImplementation{}
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
	defer rows.Close()
	if rows.Next(){
		log.Println("Masuk found")
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		//log.Println(category)

		//rows.Close()
		return category, nil
	}else{
		log.Println("Masuk not found")
		//rows.Close()
		return category, errors.New("CATEGORY NOT FOUND")
	}

}

func (repository *CategoryRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	var categories []domain.Category
	defer rows.Close()
	for rows.Next(){
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}
	return categories
}
