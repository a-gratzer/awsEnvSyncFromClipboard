package aws

import (
	"awsEnvSyncFromClipboard/internal/domain"
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
)

type AWSClientWriter struct {
	logger   *zap.Logger
	rootPath string
}

func NewWriter(logger *zap.Logger, rootPath string) *AWSClientWriter {
	return &AWSClientWriter{logger: logger, rootPath: rootPath}
}

func (w *AWSClientWriter) MustBackupCredentials(postFix string) {

	input, err := os.ReadFile(w.getCredentialsFilePath())
	if err != nil {
		fmt.Println(err)
		return
	}

	destinationFilePath := w.getCredentialsFilePath() + postFix
	err = os.WriteFile(destinationFilePath, input, 0777)
	if err != nil {
		w.logger.Fatal("Failed to backup credentials", zap.String("Path", destinationFilePath), zap.Error(err))
	}
}

func (w *AWSClientWriter) WriteCredentials(data []domain.CredentialsGroup) {

	if err := os.Truncate(w.getCredentialsFilePath(), 0); err != nil {
		w.logger.Fatal("Failed to truncate old credentials file", zap.Error(err))
	}

	file := w.mustOpenFile(w.getCredentialsFilePath())
	writer := bufio.NewWriter(file)

	for _, d := range data {
		writer.WriteString("[" + d.Name + "]")
		writer.WriteString("\n")

		writer.WriteString(w.keyValueString(AWS_ACCESS_KEY_ID, d.Credentials.AwsAccessKeyId))
		writer.WriteString("\n")

		writer.WriteString(w.keyValueString(AWS_SECRET_ACCESS_KEY, d.Credentials.AwsSecretAccessKey))
		writer.WriteString("\n")

		writer.WriteString(w.keyValueString(AWS_SESSION_TOKEN, d.Credentials.AwsSessionToken))
		writer.WriteString("\n")
	}
	writer.Flush()
}

func (w *AWSClientWriter) getCredentialsFilePath() string {
	return w.rootPath + "/" + CREDENTIALS
}

func (w *AWSClientWriter) mustOpenFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		w.logger.Fatal("failed to open file", zap.String("file", filePath), zap.Error(err))
	}
	return file
}

func (w *AWSClientWriter) keyValueString(key, value string) string {
	return key + "=" + value
}
