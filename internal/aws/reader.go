package aws

import (
	"awsEnvSyncFromClipboard/internal/domain"
	"bufio"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
)

const (
	CREDENTIALS           = "credentials"
	CONFIG                = "config"
	AWS_ACCESS_KEY_ID     = "aws_access_key_id"
	AWS_SECRET_ACCESS_KEY = "aws_secret_access_key"
	AWS_SESSION_TOKEN     = "aws_session_token"
)

type AWSClientReader struct {
	logger   *zap.Logger
	rootPath string
}

func NewAWSClientReader(logger *zap.Logger, rootPath string) *AWSClientReader {
	return &AWSClientReader{logger: logger, rootPath: rootPath}
}

func (r AWSClientReader) GetCredentials() []domain.CredentialsGroup {

	var groups = []domain.CredentialsGroup{}
	r.sequentialRead(CREDENTIALS, func(line string) {
		if line[0] == '[' {
			groups = r.attachNewGroup(groups, line)
		} else {
			r.updateLastGroup(groups, line)
		}
	})

	return groups
}

func (r AWSClientReader) attachNewGroup(groups []domain.CredentialsGroup, line string) []domain.CredentialsGroup {
	if line == "" || line[0] != '[' {
		return nil
	}

	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")

	group := domain.CredentialsGroup{}
	group.Name = line
	group.Credentials = domain.Credentials{}

	return append(groups, group)
}

func (r AWSClientReader) updateLastGroup(groups []domain.CredentialsGroup, line string) {
	if line == "" || line[0] == '[' {
		return
	}
	lastGroup := &groups[len(groups)-1]

	if strings.Contains(line, AWS_ACCESS_KEY_ID) {
		value := strings.Replace(line, AWS_ACCESS_KEY_ID, "", -1)
		value = strings.Replace(value, "=", "", -1)
		lastGroup.Credentials.AwsAccessKeyId = value
	} else if strings.Contains(line, AWS_SECRET_ACCESS_KEY) {
		value := strings.Replace(line, AWS_SECRET_ACCESS_KEY, "", -1)
		value = strings.Replace(value, "=", "", -1)
		lastGroup.Credentials.AwsSecretAccessKey = value
	} else if strings.Contains(line, AWS_SESSION_TOKEN) {
		value := strings.Replace(line, AWS_SESSION_TOKEN, "", -1)
		value = strings.Replace(value, "=", "", -1)
		lastGroup.Credentials.AwsSessionToken = value
	}
}

func (r AWSClientReader) LogRawCredentials() {
	r.sequentialRead(CREDENTIALS, func(line string) {
		r.logger.Info(line)
	})
}

func (r AWSClientReader) LogRawConfig() {
	r.sequentialRead(CONFIG, func(line string) {
		r.logger.Info(line)
	})
}

func (r AWSClientReader) sequentialRead(fileName string, callback func(line string)) {
	file := r.mustOpenFile(r.rootPath + "/" + fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
}

func (r *AWSClientReader) mustOpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		r.logger.Fatal("failed to open file", zap.String("file", filePath), zap.Error(err))
	}
	return file
}

func (r *AWSClientReader) mustReadFile(file *os.File) []byte {
	b, err := io.ReadAll(file)
	if err != nil {
		r.logger.Fatal("failed to open file", zap.String("file", file.Name()), zap.Error(err))
	}
	return b
}
