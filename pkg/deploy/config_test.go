package deploy

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestGetConfig(t *testing.T) {
	os.Setenv("CIRCLE_BUILD_NUM", "123")
	os.Setenv("CIRCLE_BUILD_URL", "http://foo.com/build/123")
	os.Setenv("CIRCLE_PROJECT_REPONAME", "deploy")
	os.Setenv("CIRCLE_PROJECT_USERNAME", "bob-inc")
	os.Setenv("CIRCLE_SHA1", "abcdef1234567890")
	os.Setenv("CIRCLE_USERNAME", "bob")
	os.Setenv("GCLOUD_SERVICE_KEY", "foo")
	os.Setenv("GITHUB_ACCESS_TOKEN", "abc123token")
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
container_name: appName
deployment_name: deployImage
image_name: theImage
branches:
  - branch:
      name: master
      project: prod
      cluster: production
      newrelic: a123
  - branch:
      name: staging
      project: dev
      cluster: staging
      newrelic: b098
`)

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	myConfig, err := GetConfig()
	assert.Nil(t, err)
	assert.Equal(t, "appName", myConfig.ContainerName)
	assert.Equal(t, "deployImage", myConfig.DeploymentName)
	assert.Equal(t, "theImage", myConfig.ImageName)
	assert.Equal(t, "master", myConfig.Branches[0].Branch.Name)
	assert.Equal(t, "prod", myConfig.Branches[0].Branch.Project)
	assert.Equal(t, "production", myConfig.Branches[0].Branch.Cluster)
	assert.Equal(t, "a123", myConfig.Branches[0].Branch.Newrelic)
	assert.Equal(t, "staging", myConfig.Branches[1].Branch.Name)
	assert.Equal(t, "dev", myConfig.Branches[1].Branch.Project)
	assert.Equal(t, "staging", myConfig.Branches[1].Branch.Cluster)
	assert.Equal(t, "b098", myConfig.Branches[1].Branch.Newrelic)
	assert.Equal(t, "123", viper.Get("CIRCLE_BUILD_NUM"))
	assert.Equal(t, "http://foo.com/build/123", viper.Get("CIRCLE_BUILD_URL"))
	assert.Equal(t, "deploy", viper.Get("CIRCLE_PROJECT_REPONAME"))
	assert.Equal(t, "bob-inc", viper.Get("CIRCLE_PROJECT_USERNAME"))
	assert.Equal(t, "abcdef1234567890", viper.Get("CIRCLE_SHA1"))
	assert.Equal(t, "bob", viper.Get("CIRCLE_USERNAME"))
	assert.Equal(t, "foo", viper.Get("GCLOUD_SERVICE_KEY"))
	assert.Equal(t, "abc123token", viper.Get("GITHUB_ACCESS_TOKEN"))
}

func TestValidateFail(t *testing.T) {
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
deployment_name: deployImage
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	myConfig, err := GetConfig()
	assert.Nil(t, err)
	shouldNotBeNil := myConfig.Validate()
	assert.NotNil(t, shouldNotBeNil)
}

func TestValidateSuccess(t *testing.T) {
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
container_name: appName
deployment_name: deployImage
image_name: theImage
branches:
  - branch:
      name: master
      project: prod
      cluster: production
      newrelic: a123
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	myConfig, err := GetConfig()
	assert.Nil(t, err)
	shouldBeNil := myConfig.Validate()
	assert.Nil(t, shouldBeNil)
}

func TestValidateBranchesFail(t *testing.T) {
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
container_name: appName
deployment_name: deployImage
image_name: theImage
`)

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	myConfig, _ := GetConfig()
	shouldNotBeNil := myConfig.Validate()
	assert.EqualError(t, shouldNotBeNil, "Missing configuration object branches")
	assert.NotNil(t, shouldNotBeNil)
}

func TestValidateConfigFileFail(t *testing.T) {
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
image_name: []
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	_, err := GetConfig()
	assert.NotNil(t, err)
}

func TestValidateEnvFail(t *testing.T) {
	os.Unsetenv("CIRCLE_BUILD_NUM")
	viper.SetConfigType("yaml")

	var yamlExample = []byte(`
container_name: appName
deployment_name: deployImage
image_name: theImage
branches:
  - branch:
      name: master
      project: prod
      cluster: production
      newrelic: a123
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	myConfig, _ := GetConfig()
	shouldNotBeNil := myConfig.Validate()
	assert.EqualError(t, shouldNotBeNil, "Missing env var CIRCLE_BUILD_NUM")
	assert.NotNil(t, shouldNotBeNil)
}
