package models

type SPORT string

const (
	SPORT_TENNIS       SPORT = "TENNIS"
	SPORT_PADDLE       SPORT = "PADDLE"
	SPORT_TABLE_TENNIS SPORT = "TABLE_TENNIS"
	SPORT_FOOTBALL     SPORT = "FOOTBALL"
)

func (g SPORT) IsValid() bool {
	switch g {
	case SPORT_TENNIS, SPORT_PADDLE, SPORT_TABLE_TENNIS, SPORT_FOOTBALL:
		return true
	}
	return false
}
