package errors

import (
	"errors"
	"testing"
)

type UniversalError struct{}

func (err *UniversalError) Error() string {
	return "universal error"
}

func (err *UniversalError) Is(_ error) bool {
	return true
}

func (err *UniversalError) As(target interface{}) bool {
	if targetErr, ok := target.(**entityError); ok {
		*targetErr = &entityError{
			err: err,
			id:  -1,
		}

		return true
	}

	return true
}

func TestErrors(t *testing.T) {
	err1 := NewEntityNotFound(1)

	// равны, потому что err1.Unwrap() == EntityNotFound
	if !errors.Is(err1, EntityNotFound) {
		t.Errorf("expected %v, got %v", EntityNotFound, err1)
	}

	err2 := NewEntityExists(2)
	err3 := NewEntityNotSaved(2)
	err4 := errors.Join(err2, err3)

	// равны, потому что err4.Unwrap() == []error{err2, err3}
	// и err2.Unwrap() == EntityExists, err3.Unwrap() == EntityNotSaved
	if !errors.Is(err4, err2) {
		t.Errorf("expected %v, got %v", err2, err4)
	}

	if !errors.Is(err4, err3) {
		t.Errorf("expected %v, got %v", err3, err4)
	}

	// равны, потому что universalError.Is() == true
	if !errors.Is(&UniversalError{}, errors.New("error")) {
		t.Errorf("expected universal error")
	}

	// Приводимы поскольку они все являются entityError
	var err5 *entityError
	if !errors.As(err1, &err5) && err5.id != 1 {
		t.Errorf("expected %v, got %v", err1, err5)
	}

	var err6 *entityError
	if !errors.As(err2, &err6) && err6.id != 2 {
		t.Errorf("expected %v, got %v", err2, err5)
	}

	var err7 *entityError
	if !errors.As(err3, &err7) && err7.id != 2 {
		t.Errorf("expected %v, got %v", err3, err5)
	}

	var err8 *entityError
	if !errors.As(err4, &err8) && err8.id != 2 {
		t.Errorf("expected %v, got %v", err4, err5)
	}

	// Приводимо, потому что UniversalError реализует As
	var err9 *entityError
	if !errors.As(&UniversalError{}, &err9) && err9.id != -1 {
		t.Errorf("expected universal error")
	}
}
