package datavalidator

import (
    "gopkg.in/go-playground/validator.v9"
    "github.com/arturoguerra/goautoplex/internal/datavalidator/revars"
)

type CustomValidator struct {
    Validator *validator.Validate
}

func inArray(value string, array []string) bool {
    for _, i := range array {
        if i == value {
            return true
        }
    }
    return false
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.Validator.Struct(i)
}

func LinuxPath(fl validator.FieldLevel) bool {
    return revars.LinuxPath.MatchString(fl.Field().String())
}

func NzbGetStatus(fl validator.FieldLevel) bool {
    statuses := []string{"SUCCESS","WARNING","FAILURE","DELETED"}

    return inArray(fl.Field().String(), statuses)
}

func NzbGetCategory(fl validator.FieldLevel) bool {
    categories := []string{"movies", "shows", "anime"}

    return inArray(fl.Field().String(), categories)
}

func StrInt(fl validator.FieldLevel) bool {
    return revars.StrInt.MatchString(fl.Field().String())
}

func New() *CustomValidator {
    validator := validator.New()

    // Shared validators
    validator.RegisterValidation("linuxpath", LinuxPath)

    // NzbGet Validators
    validator.RegisterValidation("nzbgetstatus", NzbGetStatus)
    validator.RegisterValidation("nzbgetcategory", NzbGetCategory)

    // Deluge Validators
    validator.RegisterValidation("strint", StrInt)

    return &CustomValidator{validator}
}
