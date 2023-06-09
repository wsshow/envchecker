package config

import (
	"encoding/json"
	"envchecker/pkg/dl"
	"envchecker/pkg/pterm"
	"envchecker/utils"
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type CheckObj struct {
	Name          string
	GetVersion    string
	ExpectVersion string
	Url           string
	Steps         []string
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
		execPath, err := utils.FindExecPath(obj.Name)
		if err != nil || len(execPath) == 0 {
			pterm.Warning(fmt.Sprintf("not found %s", obj.Name))
			if pterm.Confirm(fmt.Sprintf("Whether to download %s", obj.Name)) {
				if !strings.HasPrefix(obj.Url, "http") {
					pterm.Warning(fmt.Sprintf("not found valid url, current: %s", obj.Url))
					continue
				}
				err = dl.SingleDownload(obj.Url)
				if err != nil {
					pterm.Error(fmt.Sprintf("%s download failed", obj.Name))
				}
			}
			continue
		}
		pterm.Info(obj.Name, "check...")
		ss := strings.Split(obj.GetVersion, " ")
		if len(ss) < 2 || ss[0] != obj.Name {
			pterm.Error(fmt.Sprintf("config GetVersion %s name not match with config name %s", obj.GetVersion, obj.Name))
			continue
		}
		currentVersion, err := utils.Cmd(obj.Name, ss[1:]...)
		if err != nil {
			return err
		}
		expectVersion := obj.ExpectVersion
		if strings.Contains(currentVersion, expectVersion) {
			pterm.Success(fmt.Sprintf("current version is %s, which meets requirements", currentVersion))
			continue
		} else {
			pterm.Warning(fmt.Sprintf("current verison: %s, expect version: %s", currentVersion, expectVersion))
			if pterm.Confirm(fmt.Sprintf("Whether to download %s", obj.Name)) {
				if !strings.HasPrefix(obj.Url, "http") {
					pterm.Warning(fmt.Sprintf("not found valid url, current: %s", obj.Url))
					continue
				}
				err = dl.SingleDownload(obj.Url)
				if err != nil {
					pterm.Error(fmt.Sprintf("%s download failed", obj.Name))
				}
			}
			continue
		}

	}
	return nil
}

func Init() error {
	c := Config{
		[]CheckObj{
			{
				Name:          "node",
				GetVersion:    "node -v",
				ExpectVersion: "16.16.0",
				Url:           "https://nodejs.org/dist/v16.16.0/node-v16.16.0.tar.gz",
			},
			{
				Name:          "npm",
				GetVersion:    "npm -v",
				ExpectVersion: "8.11.0",
				Steps:         []string{"npm -g install npm@8.11.0"},
			},
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
