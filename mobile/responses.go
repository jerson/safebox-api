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

// AccountsResponse ...
type AccountsResponse struct {
	Accounts []*AccountSingle
}

// AccountSingle ...
type AccountSingle struct {
	ID       int64
	Label    string
	Username string
	Hint     string
}

// LogoutResponse ...
type LogoutResponse struct {
}
