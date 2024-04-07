package api

type SaveUserResponse struct {
	UID     string `json:"uid"`
	Success bool   `json:"success"`
}
