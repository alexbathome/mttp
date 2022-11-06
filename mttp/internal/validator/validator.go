package validator

import (
	"errors"
	"fmt"
	"strings"
)

const (
	invalidCharactersCommon = "!@#$%^&*():\\+=[]|\"';<>,.?"
	invalidCharactersServer = invalidCharactersCommon + "/-"
)

// validationErrors is a map of the invalid characters to their index in the string
type validationErrors map[byte][]int

// getValidationErrors is a recursive function that works it's way through the given string and returns a
// validationErrors map.
func getValidationErrors(unvalidated string, invalids string, startingIndex int, errors validationErrors) validationErrors {
	if i := strings.IndexAny(unvalidated, invalids); i != -1 {
		errors[unvalidated[i]] = append(errors[unvalidated[i]], i+startingIndex)
		return getValidationErrors(unvalidated[i+1:], invalids, i+startingIndex+1, errors)
	}
	return errors
}

// ValidateServerName validates a provided server name, if the server name consists
// of any invalid characters, it throws an InvalidServerName error with details
// of where the invalid characters can be found
func ValidateServerName(serverName string) error {
	ve := make(validationErrors, 0)
	e := getValidationErrors(serverName, invalidCharactersServer, 0, ve)
	var errorString string
	for k, v := range e {
		errorString += fmt.Sprintf("%c found in string at %v\n", k, v)
	}
	if errorString != "" {
		return errors.New(errorString)
	}
	return nil
}
