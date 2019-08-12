package safebox

import "safebox.jerson.dev/api/services"

// Session ...
type Session struct {
	response *services.AuthResponse
	password string
}

// SetAccessToken ...
func (s Session) SetAccessToken(token string) {
	s.response = &services.AuthResponse{
		AccessToken: token,
	}
}
func (s Session) login(response *services.AuthResponse) {
	s.response = response
}

func (s Session) logout() {
	s.response = nil
	s.password = ""
}

func (s Session) setPassword(password string) {
	s.password = password
}
