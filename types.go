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
	testerCall    bool
}

type Endpoint struct {
	Runner
	Tests []T `json:"tests"`
}

type T struct {
	ID            string        `json:"id"`
	Description   string        `json:"description"`
	Request       []interface{} `json:"requests"`
	StreamPayload bool          `json:"stream-payload"`
	Queries         []string      `json:"queries"`
	Compare       bool          `json:"compare"`
	Expectations  []interface{} `json:"expectations"`
	Skip          bool          `json:"skip"`
	Response      []byte        `json:"-"`
	Print         bool          `json:"print"`
	GrpcurlFlags  string        `json:"grpcurl-flags"`
	IgnoreOrder   bool          `json:"ignore-order"`
}
