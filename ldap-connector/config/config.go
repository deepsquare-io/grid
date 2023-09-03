package config

import (
	"fmt"
	"os"

	"github.com/deepsquare-io/the-grid/ldap-connector/validate"
	"gopkg.in/yaml.v3"
)

type Config struct {
	PeopleDN       string `yaml:"peopleDN"       validate:"required"`
	GroupDN        string `yaml:"groupDN"        validate:"required"`
	AddUserToGroup struct {
		MemberAttributes []string `yaml:"memberAttributes" validate:"required"`
	} `yaml:"addUserToGroup"`
	CreateUser struct {
		RDNAttribute       string              `yaml:"rdnAttribute" validate:"required"`
		ObjectClasses      []string            `yaml:"objectClasses"`
		UserNameAttributes []string            `yaml:"userNameAttributes"`
		DefaultAttributes  map[string][]string `yaml:"defaultAttributes"`
	} `yaml:"createUser"`
}

func (c Config) GetUserDN(user string) string {
	return fmt.Sprintf("%s=%s,%s", c.CreateUser.RDNAttribute, user, c.PeopleDN)
}

func (c *Config) Validate() error {
	return validate.I.Struct(c)
}

func ParseFile(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
