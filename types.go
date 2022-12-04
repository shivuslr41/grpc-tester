package tester

type Lister struct {
	Server    string `json:"server"`
	ProtoPath string `json:"proto_path"`
	ProtoFile string `json:"proto_file"`
	TLS       bool   `json:"tls"`
}

type Runer struct {
	Lister
	Endpoint string `json:"endpoint"`
}

type Tester struct {
	Runer
	Tests []T `json:"tests"`
}

type T struct {
	ID string `json:"id"`
}
