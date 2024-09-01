package aws

import (
	"awsEnvSyncFromClipboard/internal/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_read_config_must_panic(t *testing.T) {

	assert.Panics(t, func() {
		NewReader(logger.GetZapLogger(true), "non_existent").LogRawConfig()
	}, "no such file or directory")

}

func Test_read_config(t *testing.T) {
	assert.NotPanics(t, func() {
		NewReader(logger.GetZapLogger(true), getRootPath()).LogRawConfig()
	})
}

func Test_read_credentials(t *testing.T) {
	credentials := NewReader(logger.GetZapLogger(true), getRootPath()).ReadCredentials()
	assert.Equal(t, 3, len(credentials))
	assert.Equal(t, "develop", credentials[0].Name)
	assert.Equal(t, "develop_1", credentials[0].Credentials.AwsAccessKeyId)
	assert.Equal(t, "develop_2", credentials[0].Credentials.AwsSecretAccessKey)
	assert.Equal(t, "develop_3", credentials[0].Credentials.AwsSessionToken)

	assert.Equal(t, "preview", credentials[1].Name)
	assert.Equal(t, "preview_1", credentials[1].Credentials.AwsAccessKeyId)
	assert.Equal(t, "preview_2", credentials[1].Credentials.AwsSecretAccessKey)
	assert.Equal(t, "preview_3", credentials[1].Credentials.AwsSessionToken)

	assert.Equal(t, "prod", credentials[2].Name)
	assert.Equal(t, "prod_1", credentials[2].Credentials.AwsAccessKeyId)
	assert.Equal(t, "prod_2", credentials[2].Credentials.AwsSecretAccessKey)
	assert.Equal(t, "prod_3", credentials[2].Credentials.AwsSessionToken)
}

func getRootPath() string {
	return "../../tests/.aws"
}
