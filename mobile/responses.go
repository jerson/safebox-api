package safebox

// AuthResponse ...
type AuthResponse struct {
	AccessToken string
	DateExpire  string
	Date        string
	KeyPair     *KeyPairResponse
}

// KeyPairResponse ...
type KeyPairResponse struct {
	PublicKey  string
	PrivateKey string
}

// AddAccountResponse ...
type AddAccountResponse struct {
	ID int64
}

// AccountResponse ...
type AccountResponse struct {
	Account *Account
}

// Account ...
type Account struct {
	Label    string
	Username string
	Password string
	Hint     string
}

// AccountSingle ...
type AccountSingle struct {
	ID       int64
	Label    string
	Username string
	Hint     string
}

// Equal ...
func (s *AccountSingle) Equal(s2 *AccountSingle) bool {
	return s.ID == s2.ID
}

// LogoutResponse ...
type LogoutResponse struct {
}
