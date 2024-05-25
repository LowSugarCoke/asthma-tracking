package handlers

import (
	"reflect"
	"main/pkg/types"
)


func ValidateUser(user types.User) bool {
	// Check if the Username and Email fields are empty
	if user.Username == "" || user.Email == "" {
		return false
	}
	return true
}


// Check if the field is required
func isRequired(fieldName string) bool {
	// Here you can implement your logic to determine if a field is required
	// For example, you can use struct tags to mark fields as required
	// or maintain a separate list of required field names
	// For demonstration purposes, let's assume all fields are required
	return true
}

// Check if the field is empty
func isEmpty(field reflect.Value) bool {
	// Here you can implement your logic to determine if a field is empty
	// For demonstration purposes, let's assume the field is empty if its zero value
	zero := reflect.Zero(field.Type())
	return field.Interface() == zero.Interface()
}