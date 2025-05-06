package competitor

// Хранение характеристик участников по их ID
var AllCompetitors map[int]*Competitor = make(map[int]*Competitor)

// Структура, хранящая характеристики участника
type Competitor struct {
	ScheduledStart uint32
	ActualStart    uint32
	PrevStart      uint32
	FinalTime      uint32
	AllTime        uint32
	Laps           []uint32
	PenaltyStart   uint32
	PenaltyTime    uint32
	PenaltyAmount  uint32
	Hits           int
	Shots          int
	ExtraInfo      string
}

func NewCompetitor() *Competitor {
	return &Competitor{}
}
