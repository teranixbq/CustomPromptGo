package validation

import (
	"errors"
	"promptgo/util/constanta"
	"strconv"
	"strings"
)

func CheckDataEmpty(data ...any) error {
	for _, value := range data {
		if value == "" {
			return errors.New(constanta.ERROR_EMPTY)
		}
		if value == 0 {
			return errors.New(constanta.ERROR_EMPTY)
		}
	}
	return nil
}

func CheckEqualData(data string, validData []string) (string, error) {
	inputData := strings.ToLower(data)

	isValidData := false
	for _, data := range validData {
		if inputData == strings.ToLower(data) {
			isValidData = true
			break
		}
	}

	if !isValidData {
		return "", errors.New(constanta.ERROR_INVALID_INPUT)
	}

	return inputData, nil
}


func MaxLength(data string, maxLength int) error {
	if len(data) > maxLength {
		return errors.New(constanta.ERROR_MAX_LENGTH + strconv.Itoa(maxLength))
	}
	return nil
}
