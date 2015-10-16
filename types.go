package main

// Github configuration structure
type Github struct {
	Token string `json:"token"`
}

// Config structure
type Config struct {
	Github `json:"github"`
}

// UpdateAttr json to update an issue
type UpdateAttr struct {
	Status string `json:"status"`
}
