package models

type GENRE string

const (
	Male   GENRE = "M"
	Female GENRE = "F"
	Other  GENRE = "O"
)

func (g GENRE) IsValid() bool {
	switch g {
	case Male, Female, Other:
		return true
	}
	return false
}
