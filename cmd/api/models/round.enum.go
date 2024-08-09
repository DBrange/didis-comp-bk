package models

type ROUND string

const (
	ROUND_GROUP ROUND = "GROUP"
	ROUND_1R    ROUND = "1R"
	ROUND_2R    ROUND = "2R"
	ROUND_3R    ROUND = "3R"
	ROUND_4R    ROUND = "4R"
	ROUND_5R    ROUND = "5R"
	ROUND_OF    ROUND = "OF"
	ROUND_CF    ROUND = "CF"
	ROUND_SF    ROUND = "SM"
	ROUND_F     ROUND = "F"
)

func (g ROUND) IsValid() bool {
	switch g {
	case ROUND_GROUP,
		ROUND_1R,
		ROUND_2R,
		ROUND_3R,
		ROUND_4R,
		ROUND_5R,
		ROUND_OF,
		ROUND_CF,
		ROUND_SF,
		ROUND_F:
		return true
	}
	return false
}
