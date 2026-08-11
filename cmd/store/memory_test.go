package store_test

import (
	"box-manager/cmd/model"
	"box-manager/cmd/store/memorystore"
	"context"
	"testing"
	"time"
)

func TestCreateFighter_Success(t *testing.T) {
	// Arrange: создаём хранилище и бойца для вставки
	memStore := memorystore.NewMemoryStore()
	input := model.Fighter{
		FirstName: "Muhammad",
		LastName:  "Ali",
		BirthDate: time.Date(1942, 1, 17, 0, 0, 0, 0, time.UTC),
		Weight:    97.5,
		Category:  "Heavyweight",
	}

	// Act: вызываем CreateFighter
	created, err := memStore.CreateFighter(context.Background(), input)

	// Assert: ошибки быть не должно
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// ID должен стать 1 (первый в хранилище)
	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}

	// Поля должны совпадать с исходными
	if created.FirstName != input.FirstName {
		t.Errorf("FirstName: want %q, got %q", input.FirstName, created.FirstName)
	}
	if created.Weight != input.Weight {
		t.Errorf("Weight: want %f, got %f", input.Weight, created.Weight)
	}
	// Проверяем дату рождения — Compare или Equal
	if !created.BirthDate.Equal(input.BirthDate) {
		t.Errorf("BirthDate: want %v, got %v", input.BirthDate, created.BirthDate)
	}

	// Проверяем, что боец действительно сохранён в хранилище
	retrieved, err := memStore.GetFighter(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetFighter failed: %v", err)
	}
	// Сравниваем двух бойцов (можно по полям или reflect.DeepEqual)
	if retrieved.ID != created.ID ||
		retrieved.FirstName != created.FirstName ||
		retrieved.LastName != created.LastName ||
		!retrieved.BirthDate.Equal(created.BirthDate) ||
		retrieved.Weight != created.Weight ||
		retrieved.Category != created.Category {
		t.Errorf("retrieved fighter %+v doesn't match created %+v", retrieved, created)
	}
}

func TestCreateFighter_Multiple(t *testing.T) {
	memStore := memorystore.NewMemoryStore()

	f1 := model.Fighter{FirstName: "Mike", LastName: "Tyson"}
	f2 := model.Fighter{FirstName: "Evander", LastName: "Holyfield"}

	c1, err1 := memStore.CreateFighter(context.Background(), f1)
	if err1 != nil {
		t.Fatalf("unexpected error: %v", err1)
	}
	c2, err2 := memStore.CreateFighter(context.Background(), f2)
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	if c1.ID != 1 {
		t.Errorf("first ID expected 1, got %d", c1.ID)
	}
	if c2.ID != 2 {
		t.Errorf("second ID expected 2, got %d", c2.ID)
	}
}
