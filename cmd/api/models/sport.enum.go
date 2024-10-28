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

func ParseSport(s string) (SPORT, error) {
	switch s {
	case string(SPORT_TENNIS), string(SPORT_PADDLE), string(SPORT_TABLE_TENNIS), string(SPORT_FOOTBALL):
		return SPORT(s), nil
	default:
		return "", ParseErr("sport")
	}
}
