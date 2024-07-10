package utils

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate = validator.New()

func init() {
    // Creamos una nueva instancia del validador
    Validate = validator.New()
    
    // Registramos validaciones personalizadas para GENRE
    Validate.RegisterValidation("genre", validateGenre)
}

// Función de validación personalizada para GENRE
func validateGenre(fl validator.FieldLevel) bool {
    // Convertimos el valor del campo a tipo GENRE
    genre := models.GENRE(fl.Field().String())
    
    // Llamamos al método IsValid() para verificar si el valor es válido
    return genre.IsValid()
}