package dto

type GetUserPrimaryDataDTORes struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Username  string   `json:"username"`
	Image     string   `json:"image"`
}
