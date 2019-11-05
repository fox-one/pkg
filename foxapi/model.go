package foxapi

type User struct {
	ID          string `json:"id,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Name        string `json:"fullname,omitempty"`
	Language    string `json:"language,omitempty"`
	PhoneCode   string `json:"phone_code,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}
