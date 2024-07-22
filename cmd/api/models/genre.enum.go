package models

type GENRE string

const (
	GENRE_MALE   GENRE = "M"
	GENRE_FEMALE GENRE = "F"
	GENRE_OTHER  GENRE = "O"
)

func (g GENRE) IsValid() bool {
	switch g {
	case GENRE_MALE, GENRE_FEMALE, GENRE_OTHER:
		return true
	}
	return false
}
