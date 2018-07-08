package HueApiModels

import "encoding/json"

type CreateUserResults []CreateUserResult

func UnmarshalCreateUserResults(data []byte) (CreateUserResults, error) {
	var r CreateUserResults
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateUserResults) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreateUserResult struct {
	Error   Error   `json:"error"`
	Success Success `json:"success"`
}

type Error struct {
	Type        int64  `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Success struct {
	Username string `json:"username"`
}
