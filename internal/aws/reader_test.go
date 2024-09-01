package aws

import (
	"awsEnvSyncFromClipboard/internal/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_read_config_must_panic(t *testing.T) {

	assert.Panics(t, func() {
		NewAWSClientReader(logger.GetZapLogger(true), "non_existent").LogRawConfig()
	}, "no such file or directory")

}

func Test_read_config(t *testing.T) {
	assert.NotPanics(t, func() {
		NewAWSClientReader(logger.GetZapLogger(true), getRootPath()).LogRawConfig()
	})
}

func Test_read_credentials(t *testing.T) {
	NewAWSClientReader(logger.GetZapLogger(true), getRootPath()).GetCredentials()
}

func getRootPath() string {
	return "../../tests/.aws"
}
