package format

import (
	"agora/core/models"
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	f := NewFormat()
	f.reverseFunc = func(str string) string {
		return "hola"
	}
	cases := []struct {
		input *models.Transport
		want  string
	}{
		{input: &models.Transport{Value: "Hola"}, want: "aloH"},
		{input: &models.Transport{Value: "ana"}, want: "ana"},
		{input: &models.Transport{Value: "martes"}, want: "setram"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Reverse #%d", i), func(t *testing.T) {
			got, err := f.Reverse(c.input)
			if err != nil {
				t.Fail()
			}
			if got.Value != c.want {
				t.Errorf("Expected %s, got %s", c.want, got.Value)
			}
		})
	}
}

func TestReverseError(t *testing.T) {
	f := NewFormat()
	cases := []struct {
		input *models.Transport
	}{
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Reverse nil #%d", i), func(t *testing.T) {
			_, err := f.Reverse(c.input)
			if err == nil {
				t.Errorf("Expected error, got %v", err)
			}
		})

	}
}

func TestUpper(t *testing.T) {
	f := NewFormat()
	cases := []struct {
		input *models.Transport
		want  string
	}{
		{input: &models.Transport{Value: "Hola"}, want: "HOLA"},
		{input: &models.Transport{Value: "ana"}, want: "ANA"},
		{input: &models.Transport{Value: "martes"}, want: "MARTES"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Upper #%d", i), func(t *testing.T) {
			got, err := f.Upper(c.input)
			if err != nil {
				t.Fail()
			}
			if got.Value != c.want {
				t.Errorf("Expected %s, got %s", c.want, got.Value)
			}
		})
	}
}

func TestUpperError(t *testing.T) {
	f := NewFormat()
	cases := []struct {
		input *models.Transport
	}{
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Upper nil #%d", i), func(t *testing.T) {
			_, err := f.Upper(c.input)
			if err == nil {
				t.Errorf("Expected error, got %v", err)
			}
		})

	}
}

func TestLower(t *testing.T) {
	f := NewFormat()
	cases := []struct {
		input *models.Transport
		want  string
	}{
		{input: &models.Transport{Value: "Hola"}, want: "hola"},
		{input: &models.Transport{Value: "ANA"}, want: "ana"},
		{input: &models.Transport{Value: "MARTES"}, want: "martes"},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Lower #%d", i), func(t *testing.T) {
			got, err := f.Lower(c.input)
			if err != nil {
				t.Fail()
			}
			if got.Value != c.want {
				t.Errorf("Expected %s, got %s", c.want, got.Value)
			}
		})
	}
}

func TestLowerError(t *testing.T) {
	f := NewFormat()
	cases := []struct {
		input *models.Transport
	}{
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
		{input: &models.Transport{Value: ""}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Lower nil #%d", i), func(t *testing.T) {
			_, err := f.Lower(c.input)
			if err == nil {
				t.Errorf("Expected error, got %v", err)
			}
		})

	}
}
