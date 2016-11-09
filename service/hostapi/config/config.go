package config

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/golang/glog"
	configuration "github.com/rancher/agent/utilities/config"
	"io/ioutil"
	"os"
	"strconv"
)

type config struct {
	CAdvisorUrl     string
	DockerUrl       string
	Systemd         bool
	NumStats        int
	Auth            bool
	HaProxyMonitor  bool
	Key             string
	HostUuid        string
	Port            int
	Ip              string
	ParsedPublicKey interface{}
	HostUuidCheck   bool
	EventsPoolSize  int
	CattleUrl       string
	CattleAccessKey string
	CattleSecretKey string
	PidFile         string
	LogFile         string
}

var Config config

func ParsedPublicKey() error {
	keyBytes, err := ioutil.ReadFile(Config.Key)
	if err != nil {
		glog.Error("Error reading file")
		return err
	}

	block, _ := pem.Decode(keyBytes)
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	Config.ParsedPublicKey = pubKey

	return nil
}

func Parse() error {
	uuid, err := configuration.DockerUUID()
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(configuration.HostAPIPort())
	if err != nil {
		return err
	}
	Config.HaProxyMonitor = false
	Config.Port = port
	Config.Ip = configuration.HostAPIIP()
	Config.DockerUrl = "unix:///var/run/docker.sock"
	Config.Auth = true
	Config.HostUuid = uuid
	Config.HostUuidCheck = true
	Config.Key = configuration.JwtPublicKeyFile()
	Config.EventsPoolSize = 10
	Config.CattleUrl = configuration.APIURL("")
	Config.CattleAccessKey = configuration.AccessKey()
	Config.CattleSecretKey = configuration.SecretKey()

	if len(Config.Key) > 0 {
		if err := ParsedPublicKey(); err != nil {
			glog.Error("Error reading file")
			return err
		}
	}

	s, err := os.Stat("/run/systemd/system")
	if err != nil || !s.IsDir() {
		Config.Systemd = false
	} else {
		Config.Systemd = true
	}

	return nil
}
