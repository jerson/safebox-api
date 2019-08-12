package safebox

import (
	"safebox.jerson.dev/api/services"
)

// SetSessionResponse ...
func (s *SafeBox) SetSessionResponse(response *AuthResponse) {
	s.Session.response = &services.AuthResponse{
		AccessToken: response.AccessToken,
		DateExpire:  response.DateExpire,
		Date:        response.Date,
	}
}

// SetSessionPassword ...
func (s *SafeBox) SetSessionPassword(password string) {
	s.Session.password = password

}

// Session ...
type Session struct {
	response *services.AuthResponse
	password string
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
