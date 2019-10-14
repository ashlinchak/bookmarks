package validators

import (
	"errors"

	"gopkg.in/go-playground/validator.v9"
)

//
// type Account struct {
//     Name string `validate:"required,max=5"`
// }
//
// func (a *Account) Validate() map[string]map[string]string
//     return map[string]map[string]string{
//         "Name": map[string]string{
//             "required": "Name is required",
//             "max":      "Name cannot exceed 5 characters."
//         }
//     }
// }
//
// messages, err := validator.Validate(&Account{Name: "foobarbaz"})
// if err != nil {
//     panic(err)
// }
// if len(messages) > 0 {
//     fmt.Println(messages) // []string{"Name cannot exceed 5 characters."}
// }

var (
	v = validator.New()
)

// Validatable is an interface representing objects that can be validated.
type Validatable interface {
	Validate() map[string]map[string]string
}

// ValidationMessage is a string type alias that represents validation
// a validation message.
type ValidationMessage string

func (v ValidationMessage) String() string {
	return string(v)
}

// Validate will validate a Validatable object and return a slice of
// validation messages as well as an error if any is encountered.
func Validate(obj Validatable) ([]ValidationMessage, error) {
	return validate(obj)
}

func validate(obj Validatable) ([]ValidationMessage, error) {
	messages := make([]ValidationMessage, 0)
	err := v.Struct(obj)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return []ValidationMessage{}, errors.New(err.Error())
	} else if _, ok := err.(validator.ValidationErrors); ok {
		messageMap := obj.Validate()

		for _, err := range err.(validator.ValidationErrors) {
			f := err.StructField()
			t := err.Tag()

			if _, ok := messageMap[f]; ok {
				if _, ok := messageMap[f][t]; ok {
					messages = append(messages, ValidationMessage(messageMap[f][t]))
				}
			}
		}
	}

	return messages, nil
}
