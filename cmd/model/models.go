package model

import (
	"time"
)

type BoutResult int

const (
	ResultPending     BoutResult = iota // 0
	ResultWinFighter1                   // 1
	ResultWinFighter2                   // 2
	ResultDraw                          // 3
	ResultKO                            // 4
	ResultTKO                           // 5
)

type Fighter struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Weight    float64
	Category  string
}

type Club struct {
	ID       int
	Name     string
	Address  string
	Fighters []int
}

type Fight struct {
	ID         int
	Fighter1ID int
	Fighter2ID int
	Rounds     []RoundsScore
	Result     BoutResult
}
type RoundsScore struct {
	Fighter1Score int
	Fighter2Score int
}

type Tournament struct {
	ID           int
	Name         string
	Date         time.Time
	Location     string
	Participants []TournamentParticipant
	PrizeFund    float64
}
type TournamentParticipant struct {
	FighterID int
	Place     int
}
