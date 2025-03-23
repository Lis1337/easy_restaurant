package categoryRepository

import (
	"database/sql"
	"fmt"
	"github.com/Lis1337/easy_restaurant/config/db"
	"github.com/Lis1337/easy_restaurant/src/model"
	"log"
)

func FindById(id int) model.Category {
	query := `SELECT * FROM category WHERE id = $1`

	rows, err := db.ExecuteSql(query, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	rows.Next()
	res, err := buildSingleResult(rows)
	if err != nil {
		log.Fatalf("Error getting category by id = %v", id)
	}

	return res
}

func FindAll() []model.Category {
	query := `SELECT * FROM category`

	rows, err := db.ExecuteSql(query)
	if err != nil {
		panic(err)
	}

	res := buildArrayResult(rows)

	return res
}

func buildArrayResult(rows *sql.Rows) []model.Category {
	defer rows.Close()

	result := make([]model.Category, 0)
	for rows.Next() {
		category, err := buildSingleResult(rows)
		if err != nil {
			log.Fatalf("Ошибка при построении категории: %v", err)
		}

		result = append(result, category)
	}

	// Проверяем, не произошла ли ошибка при итерации
	if err := rows.Err(); err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

	return result
}

func buildSingleResult(rows *sql.Rows) (model.Category, error) {
	var id int8
	var name string
	var description sql.NullString
	var order sql.NullInt16

	err := rows.Scan(
		&id,
		&name,
		&description,
		&order,
	)
	if err != nil {
		return model.Category{}, fmt.Errorf("ошибка при сканировании строки: %w", err)
	}

	// Обрабатываем NULL-значения
	var descPtr *string
	if description.Valid {
		descPtr = &description.String
	}

	var orderPtr *int16
	if order.Valid {
		orderPtr = &order.Int16
	}

	category := model.Category{
		Id:          id,
		Name:        name,
		Description: descPtr,
		Order:       orderPtr,
	}

	return category, nil
}
