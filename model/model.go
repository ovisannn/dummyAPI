package book

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo"
)

type M map[string]interface{}

// Define our errors:
var internalError = M{"message": "internal error"}
var bookNotFound = M{"message": "book not found"}

type Book struct {
	mgm.DefaultModel `bson:",inline"`

	Name      string `json:"name" bson:"name"`
	Author    string `json:"author" bson:"author"`
	PageCount int    `json:"page_count" bson:"page_count"`
}

type Request struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	PageCount int    `json:"page_count"`
}

// Create handler create new book.
func Create(c echo.Context) error {

	requestData := &Request{}

	// skip checking bind errors.
	_ = c.Bind(requestData)

	//Validate our data:
	if err := requestData.Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok {
			return c.JSON(http.StatusInternalServerError, internalError)
		}

		return c.JSON(http.StatusBadRequest, err)
	}

	book := &Book{
		Name:      requestData.Name,
		Author:    requestData.Author,
		PageCount: requestData.PageCount,
	}

	err := mgm.Coll(book).Update(book)

	if err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, book)
}

func (r *Request) Validate() error {
	return validation.ValidateStruct(r,
		// Name can not be empty.
		validation.Field(&r.Name, validation.Required),

		// Author name can not be empty.
		validation.Field(&r.Author, validation.Required),

		validation.Field(&r.PageCount, validation.Required, validation.Min(10)),
	)
}

func Update(c echo.Context) error {

	requestData := &Request{}
	_ = c.Bind(requestData)

	if err := requestData.Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok {
			return c.JSON(http.StatusInternalServerError, internalError)
		}

		return c.JSON(http.StatusBadRequest, err)
	}

	book := &Book{}
	coll := mgm.Coll(book)

	err := coll.FindByID(c.Param("id"), book)

	if err != nil {
		return c.JSON(http.StatusNotFound, bookNotFound)
	}

	// Update our book
	book.Name = requestData.Name
	book.Author = requestData.Author
	book.PageCount = requestData.PageCount

	if err = coll.Save(book); err != nil {
		return c.JSON(http.StatusBadRequest, internalError)
	}

	return c.JSON(http.StatusOK, book)
}

func Delete(c echo.Context) error {

	book := &Book{}
	coll := mgm.Coll(book)
	err := coll.FindByID(c.Param("id"), book)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "can not find book.",
		})
	}

	if err := coll.Delete(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "internal error",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
