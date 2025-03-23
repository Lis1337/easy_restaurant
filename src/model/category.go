package model

import "fmt"

// Category dish category (ex. soup, drinks etc.)
type Category struct {
	Id          int8
	Name        string
	Description *string
	Order       *int16
}

// Метод String для преобразования данных в строку
func (c Category) String() string {
	description := "NULL"
	if c.Description != nil {
		description = fmt.Sprintf("%s", *c.Description)
	}

	order := "NULL"
	if c.Order != nil {
		order = fmt.Sprintf("%d", *c.Order)
	}

	return fmt.Sprintf("Category(ID=%d, Name=%s, Description=%s, Order=%s)",
		c.Id, c.Name, description, order)
}
