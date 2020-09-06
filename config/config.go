package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"github.com/ghodss/yaml"
	"strings"
)

const defaultApiPort = 8080

var configFile = pflag.String("config", "config.yaml", "set config file name")

type Config struct {
	Bindings []BindingConfig `json:"bindings"`
	ApiPort  int             `json:"api_port"`
}

func (c *Config) Validate() error {
	if c.ApiPort == 0 {
		c.ApiPort = defaultApiPort
	}
	if len(c.Bindings) == 0 {
		return fmt.Errorf("at least one binding must be defined")
	}
	for _, binding := range c.Bindings {
		if err := binding.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func getConfigFormat(in []byte) (string, error) {
	c := &Config{}
	yamlErr := yaml.Unmarshal(in, c)
	if yamlErr == nil {
		return "yaml", nil
	}
	jsonErr := json.Unmarshal(in, c)
	if jsonErr == nil {
		return "json", nil
	}

	return "", fmt.Errorf("yaml parsing error: %s, json parsing error: %s", yamlErr.Error(), jsonErr.Error())
}
func decodeBase64(in string) string {
	// base64 string cannot contain space so this is indication of base64 string
	if !strings.Contains(in, " ") {
		sDec, err := base64.StdEncoding.DecodeString(in)
		if err != nil {
			log.Println(fmt.Sprintf("error decoding config file base64 string: %s ", err.Error()))
			return in
		}
		return string(sDec)
	}
	return in
}
func getConfigDataFromLocalFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	fileExt, err := getConfigFormat(data)
	if fileExt == "" {
		return "", err
	}
	if strings.HasSuffix(filename, "."+fileExt) {
		return filename, nil
	}
	return filename + "." + fileExt, nil
}

func getConfigDataFromEnv() (string, error) {
	envConfigData, ok := os.LookupEnv("CONFIG")
	envConfigData = decodeBase64(envConfigData)
	if ok {
		fileExt, err := getConfigFormat([]byte(envConfigData))
		if fileExt == "" {
			return "", err
		}
		/* #nosec */
		err = ioutil.WriteFile("./config."+fileExt, []byte(envConfigData), 0644)
		if err != nil {
			return "", fmt.Errorf("cannot save environment config file")
		}
		return "./config." + fileExt, nil
	}
	return "", fmt.Errorf("no config data from environment variable")
}
func getConfigFile() (string, error) {
	if *configFile != "" {
		loadedConfigFile, err := getConfigDataFromLocalFile(*configFile)
		if err != nil {
			return "", err
		}
		return loadedConfigFile, nil
	} else {
		loadedConfigFile, err := getConfigDataFromEnv()
		if err != nil {
			return "", err
		}
		return loadedConfigFile, nil
	}
}

func Load() (*Config, error) {
	pflag.Parse()
	viper.AddConfigPath("./")
	loadedConfigFile, err := getConfigFile()
	if err != nil {
		return nil, fmt.Errorf("error loading configuration, %w", err)
	} else {
		viper.SetConfigFile(loadedConfigFile)
	}
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}