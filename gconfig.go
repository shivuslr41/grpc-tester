package tester

var GConf = GConfig{}

type GConfig struct {
	Use bool
	Lister
	Endpoint      string
	GrpcurlFlags  string
	StreamPayload bool
	Compare       bool
	Print         bool
}
