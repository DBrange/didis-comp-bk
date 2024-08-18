package utils

import (
	"reflect"
	"regexp"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	Validate.RegisterValidation("genre", validateGenre)
	Validate.RegisterValidation("sport", validateSport)
	Validate.RegisterValidation("round", validateRound)
	Validate.RegisterValidation("surface", validateSurface)
	Validate.RegisterValidation("tournamentCapacity", validateTournamentCapacity)
	Validate.RegisterValidation("competitorType", validateCompetitorType)
	Validate.RegisterValidation("rangeMovement", validateRangeMovement)
	Validate.RegisterValidation("bool", validateBool)
	Validate.RegisterValidation("availStatus", validateAvailabilitySatatus)
	Validate.RegisterValidation("day", validateDay)
	Validate.RegisterValidation("timeSlot", validateTimeSlot)
	Validate.RegisterValidation("password", validatePassword)
	Validate.RegisterValidation("string", IsString)
	Validate.RegisterValidation("number", IsNumber)
	Validate.RegisterValidation("date", IsDate)
}

func validateGenre(fl validator.FieldLevel) bool {
	genre := models.GENRE(fl.Field().String())
	return genre.IsValid()
}

func validateSport(fl validator.FieldLevel) bool {
	genre := models.SPORT(fl.Field().String())
	return genre.IsValid()
}

func validateRound(fl validator.FieldLevel) bool {
	genre := models.ROUND(fl.Field().String())
	return genre.IsValid()
}

func validateSurface(fl validator.FieldLevel) bool {
	surface, ok := fl.Field().Interface().(models.TENNIS_SURFACE)
	// Si la superficie es una cadena vacía, no validar y considerar válida
	if surface == "" {
		return true
	}

	if !ok {
		return false
	}

	return surface.IsValid()
}

func validateTournamentCapacity(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.Int {
		return false
	}

	capacity := models.TOURNAMENT_CAPACITY(fl.Field().Int())
	return capacity.IsValid()
}

func validateCompetitorType(fl validator.FieldLevel) bool {
	genre := models.COMPETITOR_TYPE(fl.Field().String())
	return genre.IsValid()
}

func validateRangeMovement(fl validator.FieldLevel) bool {
	genre := models.RANGE_MOVEMENT(fl.Field().String())
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

// IsString verifica si el campo es un string
func IsString(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(string)
	return ok
}

// IsNumber verifica si el campo es un número (entero o flotante)
func IsNumber(fl validator.FieldLevel) bool {
	switch fl.Field().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsDate verifica si el campo es una fecha válida en formato YYYY-MM-DD
func IsDate(fl validator.FieldLevel) bool {
	dateStr, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
