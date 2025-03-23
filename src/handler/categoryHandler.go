package categoryHandler

import (
	"encoding/json"
	"fmt"
	categoryRepository "github.com/Lis1337/easy_restaurant/src/repository"
	"log"
	"net/http"
	"strconv"
)

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		panic(fmt.Sprintf("Invalid id given, message: %s", err))
	}

	category := categoryRepository.FindById(id)

	categoryJson, err := json.Marshal(category)
	if err != nil {
		log.Fatalf("Ошибка при сериализации: %v", err)
	}
	res := string(categoryJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := categoryRepository.FindAll()

	categoriesJson, err := json.Marshal(categories)
	if err != nil {
		log.Fatalf("Ошибка при сериализации: %v", err)
	}
	res := string(categoriesJson)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}
