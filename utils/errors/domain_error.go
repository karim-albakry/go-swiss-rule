package errors

type domainErr struct {
	Message string
}

func (thisDE *domainErr) Error() string {
	return thisDE.Message
}

func SimpleError(message string) error {
	return &domainErr{Message: message}
}
