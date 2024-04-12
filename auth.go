package s1

import (
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func New(url, token string) (*S1, error) {
	auth, err := securityprovider.NewSecurityProviderApiKey("header", "Authorization", fmt.Sprintf("ApiToken %s", token))
	if err != nil {
		return nil, err
	}
	client, err := NewClient(url, WithRequestEditorFn(auth.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}
