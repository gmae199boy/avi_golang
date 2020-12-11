package model

type User struct {
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	// Video map[string][]string `json:"video,omitempty"`
	// Image map[string][]string `json:"image,omitempty"`
}
