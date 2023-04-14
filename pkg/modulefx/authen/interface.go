package authen


type Authenticator interface {
	Authen(apiKey string) bool
}
