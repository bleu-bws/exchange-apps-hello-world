package services

// AppService -
type AppService interface {
	GetAcessToken() error
	AccessEndpoint(string) ([]byte, error)
}
