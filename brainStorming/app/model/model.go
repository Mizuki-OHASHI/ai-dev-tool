package model

import "errors"

type SomeEntity struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e *SomeEntity) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.Value == "" {
		return errors.New("value is required")
	}
	return nil
}
