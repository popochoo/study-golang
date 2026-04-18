package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type CompleteDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func ErrorHandling(w http.ResponseWriter, err error, customErr error) {
	errDTO := NewErrorDTO(err)

	if errors.Is(err, customErr) {
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
	}
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
