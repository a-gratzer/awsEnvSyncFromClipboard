package aws

import (
	"awsEnvSyncFromClipboard/internal/domain"
	"awsEnvSyncFromClipboard/internal/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestAWSClientWriter_MustBackupCredentials(t *testing.T) {
	NewWriter(logger.GetZapLogger(true), getRootPath()).MustBackupCredentials(getBackupPostfix())
	_, err := os.Stat(getRootPath() + "/" + CREDENTIALS + getBackupPostfix())
	assert.NoError(t, err)
}

func TestAWSClientWriter_WriteCredentials(t *testing.T) {
	data := []domain.CredentialsGroup{}
	data = append(data, domain.CredentialsGroup{
		Name: "test",
		Credentials: domain.Credentials{
			AwsAccessKeyId:     "id",
			AwsSecretAccessKey: "key",
			AwsSessionToken:    "token",
		},
	})
	NewWriter(logger.GetZapLogger(true), getRootPath()).WriteCredentials(data)
	cred := NewReader(logger.GetZapLogger(true), getRootPath()).ReadCredentials()

	assert.Equal(t, data, cred)

}

func getBackupPostfix() string {
	return "_" + time.Now().Format("2006-01-02")
}
