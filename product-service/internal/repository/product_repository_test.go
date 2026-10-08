package repository

import (
	"errors"
	"testing"
)

func TestFindByID_Ditemukan(t *testing.T) {
	repo := NewMemoryProductRepository()

	p, err := repo.FindByID(1)
	if err != nil {
		t.Fatalf("tidak diharapkan error, dapat: %v", err)
	}
	if p.Name != "Keripik Singkong" {
		t.Errorf("nama produk = %q, diharapkan %q", p.Name, "Keripik Singkong")
	}
}

func TestFindByID_TidakDitemukan(t *testing.T) {
	repo := NewMemoryProductRepository()

	_, err := repo.FindByID(99)
	if !errors.Is(err, ErrProductNotFound) {
		t.Errorf("diharapkan ErrProductNotFound, dapat: %v", err)
	}
}
