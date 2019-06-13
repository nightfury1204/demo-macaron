package pkg

import "fmt"

type Book struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description,omitempty"`
	Price *int `json:"price,omitempty"`
}

func (b *Book) Validate() error {
	if b.Name == "" {
		return fmt.Errorf("book name must be non empty")
	}
	return nil
}

func Merge(old, nu Book) Book {
	if nu.Name != "" {
		old.Name = nu.Name
	}
	if nu.Author != "" {
		old.Author = nu.Author
	}
	if nu.Description != "" {
		old.Description = nu.Description
	}
	if nu.Price != nil {
		old.Price = nu.Price
	}
	return old
}