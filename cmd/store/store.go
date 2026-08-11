package store

import (
	"box-manager/cmd/model"
	"context"
)

type Store interface {
	CreateFighter(ctx context.Context, f model.Fighter) (model.Fighter, error)
	GetFighter(ctx context.Context, id int) (model.Fighter, error)
	UpdateFighter(ctx context.Context, id int, f model.Fighter) error
	DeleteFighter(ctx context.Context, id int) error
	ListFighters(ctx context.Context) []model.Fighter

	CreateClub(ctx context.Context, c model.Club) (model.Club, error)
	GetClub(ctx context.Context, id int) (model.Club, error)
	UpdateClub(ctx context.Context, id int, c model.Club) error
	DeleteClub(ctx context.Context, id int) error
	ListClubs(ctx context.Context) []model.Club

	CreateTournament(ctx context.Context, t model.Tournament) (model.Tournament, error)
	GetTournament(ctx context.Context, id int) (model.Tournament, error)
	UpdateTournament(ctx context.Context, id int, t model.Tournament) error
	DeleteTournament(ctx context.Context, id int) error
	ListTournaments(ctx context.Context) []model.Tournament
	AddParticipantToTournament(ctx context.Context, tournamentID, fighterID, place int) error
	RemoveParticipantFromTournament(ctx context.Context, tournamentID, fighterID int) error

	CreateFight(ctx context.Context, f model.Fight) (model.Fight, error)
	GetFight(ctx context.Context, id int) (model.Fight, error)
	UpdateFight(ctx context.Context, id int, f model.Fight) error
	DeleteFight(ctx context.Context, id int) error
	ListFights(ctx context.Context) []model.Fight
}
