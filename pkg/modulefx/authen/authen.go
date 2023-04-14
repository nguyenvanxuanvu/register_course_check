package authen

import (

	"github.com/spf13/viper"
)


type authenticatorService struct {

}

func NewAuthenticator() Authenticator {
	return &authenticatorService{}
}

func (s *authenticatorService) Authen(apiKey string) bool {
	apiKeyConfig := viper.GetString("authen.api-key")
	if apiKey == apiKeyConfig {
		return true
	} else {
		return false
	}
}
