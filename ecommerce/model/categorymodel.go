package model

import(
	"github.com/myrachanto/ecommerce/httperrors"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	// ID 	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	Majorcat string `json:"majorcat"`
	Code string `json:"code"`
	Base
}
func (category Category) Validate() *httperrors.HttpError{
	if category.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if category.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	if category.Description == "" {
		return httperrors.NewNotFoundError("Invalid Description")
	}
	return nil
}