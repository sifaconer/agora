package format

import (
	"agora/core/models"
	"fmt"
	"strings"
)

type Format struct {
	reverseFunc func(str string) string
}

func NewFormat() *Format {
	return &Format{
		reverseFunc: reverseString,
	}
}

func (f *Format) Reverse(model *models.Transport) (*models.Transport, error) {
	if model.Value == "" {
		return nil, fmt.Errorf("empty value")
	}

	v := f.reverseFunc(model.Value)

	return &models.Transport{
		Value:     v,
		FuncApply: "reverse",
	}, nil
}

func (f *Format) Upper(model *models.Transport) (*models.Transport, error) {
	if model.Value == "" {
		return nil, fmt.Errorf("empty value")
	}
	return &models.Transport{
		Value:     strings.ToUpper(model.Value),
		FuncApply: "upper",
	}, nil
}

func (f *Format) Lower(model *models.Transport) (*models.Transport, error) {
	if model.Value == "" {
		return nil, fmt.Errorf("empty value")
	}
	return &models.Transport{
		Value:     strings.ToLower(model.Value),
		FuncApply: "lower",
	}, nil
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
