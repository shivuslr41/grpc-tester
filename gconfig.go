package tester

// global configurations
var GConf = GConfig{}

// GConfig stores configurations that overrides test file configs
type GConfig struct {
	Use bool
	Lister
	Endpoint      string
	GrpcurlFlags  string
	StreamPayload bool
	Compare       bool
	Print         bool
}

// Debug
var Debug bool
