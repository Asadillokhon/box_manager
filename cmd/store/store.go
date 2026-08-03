package store

import "box-manager/cmd/model"

type Store interface {
	CreateFighter(f model.Fighter) (model.Fighter, error)
	GetFighter(id int) (model.Fighter, error)
	UpdateFighter(id int, f model.Fighter) error
	DeleteFighter(id int) error
	ListFighters() []model.Fighter

	CreateClub(c model.Club) (model.Club, error)
	GetClub(id int) (model.Club, error)
	UpdateClub(id int, c model.Club) error
	DeleteClub(id int) error
	ListClubs() []model.Club

	CreateTournament(t model.Tournament) (model.Tournament, error)
	GetTournament(id int) (model.Tournament, error)
	UpdateTournament(id int, t model.Tournament) error
	DeleteTournament(id int) error
	ListTournaments() []model.Tournament
	AddParticipantToTournament(tournamentID, fighterID, place int) error
	RemoveParticipantFromTournament(tournamentID, fighterID int) error

	CreateFight(f model.Fight) (model.Fight, error)
	GetFight(id int) (model.Fight, error)
	UpdateFight(id int, f model.Fight) error
	DeleteFight(id int) error
	ListFights() []model.Fight
}
