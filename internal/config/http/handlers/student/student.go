package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/akshaykathwate/students_api/internal/types"
	"github.com/akshaykathwate/students_api/internal/utils"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		var student types.Student 
		
		err:= json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("empty Body")))
			return
		}
		println("")
		slog.Info("creating  a student")
		fmt.Println(student.Name)


		// request validator
		// validator.New().Struct(student)
		e :=validator.New().Struct(student)
		if e!=nil{
			validateErrs:=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
		}
		
		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"ok"})

	}
}