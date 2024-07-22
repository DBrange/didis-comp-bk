package utils

import (
	"reflect"
	"regexp"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	Validate.RegisterValidation("genre", validateGenre)
	Validate.RegisterValidation("sport", validateSport)
	Validate.RegisterValidation("bool", validateBool)
	Validate.RegisterValidation("availStatus", validateAvailabilitySatatus)
	Validate.RegisterValidation("day", validateDay)
	Validate.RegisterValidation("timeSlot", validateTimeSlot)
	Validate.RegisterValidation("password", validatePassword)
}

func validateGenre(fl validator.FieldLevel) bool {
	genre := models.GENRE(fl.Field().String())
	return genre.IsValid()
}

func validateSport(fl validator.FieldLevel) bool {
	genre := models.SPORT(fl.Field().String())
	return genre.IsValid()
}

func validateBool(fl validator.FieldLevel) bool {
	return fl.Field().Kind() == reflect.Bool
}

func validateAvailabilitySatatus(fl validator.FieldLevel) bool {
	genre := models.AVAILABILITY_STATUS(fl.Field().String())
	return genre.IsValid()
}

func validateDay(fl validator.FieldLevel) bool {
	genre := models.DAY(fl.Field().String())
	return genre.IsValid()
}

func validateTimeSlot(fl validator.FieldLevel) bool {
	timeSlotRegex := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d$`)
	return timeSlotRegex.MatchString(fl.Field().String())
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return len(password) >= 10 &&
		regexp.MustCompile(`[A-Z]`).MatchString(password) &&
		regexp.MustCompile(`[0-9]`).MatchString(password) &&
		regexp.MustCompile(`[@\-_.,;+]`).MatchString(password)
}
