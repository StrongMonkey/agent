package util

import (
	"github.com/rancher/agent/service/hostapi/config"
	rclient "github.com/rancher/go-rancher/client"
)

func GetRancherClient() (*rclient.RancherClient, error) {
	apiUrl := config.Config.CattleUrl
	accessKey := config.Config.CattleAccessKey
	secretKey := config.Config.CattleSecretKey

	if apiUrl == "" || accessKey == "" || secretKey == "" {
		return nil, nil
	}

	apiClient, err := rclient.NewRancherClient(&rclient.ClientOpts{
		Url:       apiUrl,
		AccessKey: accessKey,
		SecretKey: secretKey,
	})
	if err != nil {
		return nil, err
	}
	return apiClient, nil
}
