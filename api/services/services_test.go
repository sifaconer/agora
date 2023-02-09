package services

import (
	"agora/api/proto/libgo"
	"agora/core/format"
	"agora/core/models"
	"context"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

type MockService struct{}

func (m *MockService) Reverse(model *models.Transport) (*models.Transport, error) {
	return &models.Transport{
		Value: "aloH",
	}, nil
}
func (m *MockService) Upper(model *models.Transport) (*models.Transport, error) {
	return nil, nil
}
func (m *MockService) Lower(model *models.Transport) (*models.Transport, error) {
	return nil, nil
}

func TestReverse(t *testing.T) {
	s := ServicesFormat{
		Usecase: &MockService{},
	}

	expected := "aloH"
	got, err := s.Reverse(context.Background(), &libgo.DataRequest{
		Value: "Hola",
	})
	if err != nil {
		t.Fail()
	}
	if got.Value != expected {
		t.Errorf("expected %s, got %s", expected, got.Value)
	}
}
func TestUpper(t *testing.T) {
	u := format.NewFormat()
	s := ServicesFormat{
		Usecase: u,
	}

	expected := "HOLA"
	got, err := s.Upper(context.Background(), &libgo.DataRequest{
		Value: "Hola",
	})
	if err != nil {
		t.Fail()
	}
	if got.Value != expected {
		t.Errorf("expected %s, got %s", expected, got.Value)
	}
}

func TestLower(t *testing.T) {
	u := format.NewFormat()
	s := ServicesFormat{
		Usecase: u,
	}

	expected := "hola"
	got, err := s.Lower(context.Background(), &libgo.DataRequest{
		Value: "Hola",
	})
	if err != nil {
		t.Fail()
	}
	if got.Value != expected {
		t.Errorf("expected %s, got %s", expected, got.Value)
	}
}
