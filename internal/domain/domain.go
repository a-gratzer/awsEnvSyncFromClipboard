package domain

type CredentialsGroup struct {
	Name        string
	Credentials Credentials
}

type Credentials struct {
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsSessionToken    string
}
