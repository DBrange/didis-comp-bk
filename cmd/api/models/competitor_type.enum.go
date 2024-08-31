package models

import "github.com/DBrange/didis-comp-bk/cmd/api/utils"

type COMPETITOR_TYPE string

const ( 
	COMPETITOR_TYPE_SINGLE COMPETITOR_TYPE = "S"
	COMPETITOR_TYPE_DOUBLE COMPETITOR_TYPE = "D"
	COMPETITOR_TYPE_TEAM   COMPETITOR_TYPE = "T"
)

func (g COMPETITOR_TYPE) IsValid() bool {
	switch g {
	case COMPETITOR_TYPE_SINGLE, COMPETITOR_TYPE_DOUBLE, COMPETITOR_TYPE_TEAM:
		return true
	}
	return false
}

func ParseCompetitorType(s string) (COMPETITOR_TYPE, error) {
	switch s {
	case string(COMPETITOR_TYPE_SINGLE), string(COMPETITOR_TYPE_DOUBLE), string(COMPETITOR_TYPE_TEAM):
		return COMPETITOR_TYPE(s), nil
	default:
		return "", utils.ParseErr("competitor type")
	}
}
