package config

import (
	"encoding/json"
	"envchecker/pkg/pterm"
	"envchecker/utils"
	"fmt"

	"github.com/samber/lo"
)

type CheckObj struct {
	Name    string
	Version string
	Url     string
	Command string
}

type Config struct {
	Objs []CheckObj
}

func (c *Config) Check() error {
	objs := c.Objs
	names := []string{}
	for _, obj := range objs {
		names = append(names, obj.Name)
	}
	names = pterm.Multiselect("Please select the items that you want to detect", names)
	for _, obj := range objs {
		if !lo.Contains(names, obj.Name) {
			continue
		}
		s, err := utils.Cmd("powershell", fmt.Sprintf("Get-Command -Name %s -ErrorAction SilentlyContinue", obj.Name))
		if err != nil {
			return err
		}
		if len(s) > 0 {
			pterm.Info(obj.Name, "check...")
			currentVersion, err := utils.Cmd(obj.Name, "-v")
			if err != nil {
				return err
			}
			currentVersion = utils.TrimUntilNum(currentVersion)
			expectVersion := obj.Version
			if expectVersion != currentVersion {
				pterm.Warning(fmt.Sprintf("current verison: %s, expect version: %s", currentVersion, expectVersion))
				continue
			} else {
				pterm.Success(fmt.Sprintf("current version is %s, which meets requirements", currentVersion))
				continue
			}
		}
	}
	return nil
}

func Init() error {
	c := Config{
		[]CheckObj{
			{Name: "node", Version: "16.16.0", Url: "https://nodejs.org/dist/v16.16.0/node-v16.16.0.tar.gz"},
			{Name: "npm", Version: "8.11.0", Command: " npm -g install npm@8.11.0"},
		},
	}
	return To(&c, "envchecker.json")
}

func From(confPath string) (c *Config, err error) {
	bs, err := utils.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, err
	}
	err = json.Unmarshal(bs, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func To(c *Config, confPath string) error {
	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}
	return utils.WriteFile(confPath, data)
}
