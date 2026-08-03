package main

import (
	"box-manager/cmd/api"
	"box-manager/cmd/model"
	"box-manager/cmd/store"
	"fmt"
	"time"
)

func main() {
	memory := store.NewMemoryStore()
	httpHandlers := api.NewHTTPHandler(memory)
	httpServer := api.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start HTTP server:", err)
	}

}

func Foo() {

	s := store.NewMemoryStore()
	f, err := s.CreateFighter(model.Fighter{
		FirstName: "Muhammad",
		LastName:  "Ali",
		Weight:    97.5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created:", f)

	got, err := s.GetFighter(f.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Got:", got)

	err = s.UpdateFighter(f.ID, model.Fighter{
		FirstName: "Muhammad",
		LastName:  "Ali",
		Weight:    99.5,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(s.GetFighter(f.ID))

	err = s.DeleteFighter(f.ID)
	fmt.Println(s.GetFighter(f.ID))

	a := store.NewMemoryStore()
	n, err := a.CreateClub(model.Club{
		Name:     "Gok",
		Address:  "123 str",
		Fighters: []int{1, 2, 3},
	})
	if err != nil {
		panic("sdsds")
	}

	fmt.Println("Id:", n.ID, "nam:", n.Name, "address:", n.Address)
}

func examination() {

	s := store.NewMemoryStore()

	// ======================
	// 1. Проверка Fighter
	// ======================
	fmt.Println("=== FIGHTERS ===")

	// Создание
	ali, err := s.CreateFighter(model.Fighter{
		FirstName: "Muhammad",
		LastName:  "Ali",
		BirthDate: time.Date(1942, 1, 17, 0, 0, 0, 0, time.UTC),
		Weight:    97.5,
		Category:  "Heavyweight",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created fighter: %+v\n", ali)

	tyson, err := s.CreateFighter(model.Fighter{
		FirstName: "Mike",
		LastName:  "Tyson",
		BirthDate: time.Date(1966, 6, 30, 0, 0, 0, 0, time.UTC),
		Weight:    100.0,
		Category:  "Heavyweight",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created fighter: %+v\n", tyson)

	// Получение
	got, err := s.GetFighter(ali.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got fighter: %s %s\n", got.FirstName, got.LastName)

	// Обновление
	ali.Weight = 98.0
	err = s.UpdateFighter(ali.ID, ali)
	if err != nil {
		panic(err)
	}
	updated, _ := s.GetFighter(ali.ID)
	fmt.Printf("Updated weight: %.1f kg\n", updated.Weight)

	// Список всех бойцов
	allFighters := s.ListFighters()
	fmt.Printf("Total fighters: %d\n", len(allFighters))

	// Удаление (проверим на третьем бойце)
	frazier, _ := s.CreateFighter(model.Fighter{
		FirstName: "Joe",
		LastName:  "Frazier",
		Weight:    95.0,
	})
	fmt.Printf("Created: %s %s (id=%d)\n", frazier.FirstName, frazier.LastName, frazier.ID)
	err = s.DeleteFighter(frazier.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted Joe Frazier")
	_, err = s.GetFighter(frazier.ID)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// ======================
	// 2. Проверка Club
	// ======================
	fmt.Println("\n=== CLUBS ===")

	club, err := s.CreateClub(model.Club{
		Name:     "Knockout Gym",
		Address:  "123 Main St",
		Fighters: []int{ali.ID, tyson.ID},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created club: %s (id=%d) with fighters: %v\n", club.Name, club.ID, club.Fighters)

	gotClub, _ := s.GetClub(club.ID)
	fmt.Printf("Got club: %s\n", gotClub.Name)

	// ======================
	// 3. Проверка Tournament
	// ======================
	fmt.Println("\n=== TOURNAMENTS ===")

	tour, err := s.CreateTournament(model.Tournament{
		Name:      "World Heavyweight Championship",
		Date:      time.Date(2026, 12, 25, 0, 0, 0, 0, time.UTC),
		Location:  "Madison Square Garden",
		PrizeFund: 1_000_000.0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created tournament: %s (id=%d)\n", tour.Name, tour.ID)

	// Добавление участников
	err = s.AddParticipantToTournament(tour.ID, ali.ID, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added Ali to tournament (place 1)")

	err = s.AddParticipantToTournament(tour.ID, tyson.ID, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added Tyson to tournament (place 2)")

	// Попытка добавить дубликат
	err = s.AddParticipantToTournament(tour.ID, ali.ID, 3)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Проверка участников
	updatedTour, _ := s.GetTournament(tour.ID)
	fmt.Printf("Tournament participants: %+v\n", updatedTour.Participants)

	// Удаление участника
	err = s.RemoveParticipantFromTournament(tour.ID, ali.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Removed Ali from tournament")
	finalTour, _ := s.GetTournament(tour.ID)
	fmt.Printf("Remaining participants: %+v\n", finalTour.Participants)

	// ======================
	// 4. Проверка Fight
	// ======================
	fmt.Println("\n=== FIGHTS ===")

	fight, err := s.CreateFight(model.Fight{
		Fighter1ID: ali.ID,
		Fighter2ID: tyson.ID,
		Rounds: []model.RoundsScore{
			{Fighter1Score: 10, Fighter2Score: 9},
			{Fighter1Score: 9, Fighter2Score: 10},
			{Fighter1Score: 10, Fighter2Score: 8},
		},
		Result: model.ResultWinFighter1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created fight (id=%d): Ali vs Tyson, result=%d\n", fight.ID, fight.Result)

	// Попытка создать бой с одинаковыми бойцами
	_, err = s.CreateFight(model.Fight{
		Fighter1ID: ali.ID,
		Fighter2ID: ali.ID,
	})
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Попытка создать бой с несуществующим бойцом
	_, err = s.CreateFight(model.Fight{
		Fighter1ID: ali.ID,
		Fighter2ID: 999,
	})
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	gotFight, _ := s.GetFight(fight.ID)
	fmt.Printf("Got fight: %+v\n", gotFight)

	// Обновление результата
	fight.Result = model.ResultDraw
	err = s.UpdateFight(fight.ID, fight)
	if err != nil {
		panic(err)
	}
	updatedFight, _ := s.GetFight(fight.ID)
	fmt.Printf("Updated result: %d (expected %d)\n", updatedFight.Result, model.ResultDraw)

	// Удаление боя
	err = s.DeleteFight(fight.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted fight")
	_, err = s.GetFight(fight.ID)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	fmt.Println("\n=== ALL CHECKS PASSED ===")

}
