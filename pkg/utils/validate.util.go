package utils

import (
	"regexp"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Validate *validator.Validate

func init() {
    Validate = validator.New()

    // Registramos validaciones personalizadas
    Validate.RegisterValidation("genre", validateGenre)
  	Validate.RegisterValidation("mongoid", validateMongoDBObjectID)

}

// Función de validación personalizada para GENRE
func validateGenre(fl validator.FieldLevel) bool {
    genre := models.GENRE(fl.Field().String())
    return genre.IsValid()
}

// Validación personalizada para UUID sin guiones
var mongoObjectIDRegex = regexp.MustCompile("^[a-fA-F0-9]{24}$")

func validateMongoDBObjectID(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	if !mongoObjectIDRegex.MatchString(id) {
		return false
	}
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}