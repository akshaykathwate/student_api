package response

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"github.com/go-playground/validator/v10"

)

type Response struct{
	Status string `json:"status"`
	Error string `json:"error"`
}

const(
	StatusOk ="OK"
	StatusError ="Error"
)

func WriteJson(w http.ResponseWriter , status int, data interface{}) error{

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError (err error) Response{
	return Response{
		Status : StatusError,
		Error : err.Error(),
	}
}
func ValidationError(validationErrors validator.ValidationErrors) Response {
	var errMsgs []string

	for _, fieldErr := range validationErrors {
		switch fieldErr.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("Field %s is required", fieldErr.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Field %s is invalid", fieldErr.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
