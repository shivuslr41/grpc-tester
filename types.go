package tester

type Lister struct {
	Server    string `json:"server"`
	ProtoPath string `json:"proto_path"`
	ProtoFile string `json:"proto_file"`
	TLS       bool   `json:"tls"`
}

type Runner struct {
	Lister
	Endpoint      string        `json:"endpoint"`
	Data          []interface{} `json:"-"`
	GrpcurlFlags  string        `json:"-"`
	StreamPayload bool          `json:"-"`
}

type Endpoint struct {
	Runner
	Tests []T `json:"tests"`
}

type T struct {
	ID            string        `json:"id"`
	Description   string        `json:"description"`
	Request       []interface{} `json:"request"`
	StreamPayload bool          `json:"stream-payload"`
	Query         string        `json:"query"`
	Compare       bool          `json:"compare"`
	Expect        []interface{} `json:"expect"`
	Skip          bool          `json:"skip"`
	Response      []byte        `json:"-"`
	Print         bool          `json:"print"`
	GrpcurlFlags  string        `json:"grpcurl-flags"`
	IgnoreOrder   bool          `json:"ignore-order"`
}
