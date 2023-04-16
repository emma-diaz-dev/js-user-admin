package entities

type (
	User struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Address  string `json:"address"`
	}
	UserCreateResponse struct {
		ID        string `json:"id"`
		Documents int64  `json:"documents"`
	}
	UserResponse struct {
		Users     []User `json:"users"`
		Documents int    `json:"documents"`
	}
)
