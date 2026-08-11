package api

type FighterDTO struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	BirthDate string  `json:"birth_date"` // "2006-01-02"
	Weight    float64 `json:"weight"`
	Category  string  `json:"category"`
	ClubID    int     `json:"club_id"`
}

type TournamentDTO struct {
	Name      string  `json:"name"`
	Date      string  `json:"date"` // "2006-01-02"
	Location  string  `json:"location"`
	PrizeFund float64 `json:"prize"`
}
type AddParticipantDTO struct {
	FighterID int `json:"fighter_id"`
	Place     int `json:"place"`
}
