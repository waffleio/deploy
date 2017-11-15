package deploy

import (
	"fmt"

	"github.com/spf13/viper"
)

var envVars = []string{
	"CIRCLE_BUILD_NUM",
	"CIRCLE_BUILD_URL",
	"CIRCLE_PROJECT_REPONAME",
	"CIRCLE_PROJECT_USERNAME",
	"CIRCLE_SHA1",
	"CIRCLE_USERNAME",
	"GCLOUD_SERVICE_KEY",
	"GITHUB_ACCESS_TOKEN",
}

// Config representation of our configuration
type Config struct {
	ContainerName  string `mapstructure:"container_name"`
	DeploymentName string `mapstructure:"deployment_name"`
	ImageName      string `mapstructure:"image_name"`
	Branches       []struct {
		Branch struct {
			Name     string `mapstructure:"name"`
			Project  string `mapstructure:"project"`
			Cluster  string `mapstructure:"cluster"`
			Newrelic string `mapstructure:"newrelic"`
		} `mapstructure:"branch"`
	} `mapstructure:"branches"`
}

// GetConfig loads the configuration
func GetConfig() (Config, error) {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("Failure to properly load the config file: %v", err)

	}

	for _, e := range envVars {
		err := viper.BindEnv(e)
		if err != nil {
			return c, fmt.Errorf("Unable to read: %s:%v", e, err)
		}
	}
	return c, nil
}

//Validate ensures we have all the config items required to run
func (c Config) Validate() error {
	requiredConfigItems := map[string]string{
		"container_name":  c.ContainerName,
		"deployment_name": c.DeploymentName,
		"image_name":      c.ImageName,
	}

	for k, v := range requiredConfigItems {
		if v == "" {
			return fmt.Errorf("Missing configuration object %s", k)
		}
	}

	if len(c.Branches) < 1 {
		return fmt.Errorf("Missing configuration object branches")
	}

	for _, e := range envVars {
		if viper.Get(e) == nil {
			return fmt.Errorf("Missing env var %s", e)
		}
	}
	return nil
}
