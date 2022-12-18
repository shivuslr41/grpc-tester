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
	StreamPayload bool          `json:"stream_payload"`
	testerCall    bool
}

type Endpoint struct {
	Runner
	Tests []T `json:"tests"`
}

type T struct {
	ID           string        `json:"id"`
	Description  string        `json:"description"`
	Request      []interface{} `json:"request"`
	Queries      []string      `json:"queries"`
	Compare      bool          `json:"compare"`
	Expectations []interface{} `json:"expectation"`
	Skip         bool          `json:"skip"`
	Response     []interface{} `json:"-"`
	Print        bool          `json:"print"`
	GrpcurlFlags string        `json:"grpcurl_flags"`
	IgnoreOrder  bool          `json:"ignore_order"`
	Pass         bool          `json:"-"`
}
