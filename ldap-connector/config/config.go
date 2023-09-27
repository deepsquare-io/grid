// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package config

import (
	"fmt"
	"os"

	"github.com/deepsquare-io/grid/ldap-connector/validate"
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
