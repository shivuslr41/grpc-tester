package tester

// Lister stores List configurtions
type Lister struct {
	Server    string `json:"server"`
	ProtoPath string `json:"proto_path"`
	ProtoFile string `json:"proto_file"`
	TLS       bool   `json:"tls"`
}

// Runner stores Run configurtions
type Runner struct {
	Lister
	Endpoint      string `json:"endpoint"`
	Data          []any  `json:"-"`
	GrpcurlFlags  string `json:"-"`
	StreamPayload bool   `json:"stream_payload"`
	testerCall    bool
}

// Endpoint stores tests configurations of multiple endpoints
type Endpoint struct {
	Runner
	Skip  bool `json:"skip"`
	Tests []T  `json:"tests"`
}

// T stores multiple test case configurtions per endpoint
type T struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Requests     []any    `json:"requests"`
	Queries      []string `json:"jqq"`
	Compare      bool     `json:"compare"`
	Expectations []any    `json:"expectations"`
	Skip         bool     `json:"skip"`
	Response     []any    `json:"-"`
	Print        bool     `json:"print"`
	GrpcurlFlags string   `json:"grpcurl_flags"`
	IgnoreOrder  bool     `json:"ignore_order"`
	Pass         bool     `json:"-"`
	State
}

// State stores variables(extracted data) state
type State struct {
	Replace     []string `json:"replace"`
	ReplaceFrom []string `json:"replace_from"`
	Extract     []string `json:"extract"`
	ExtractTo   []string `json:"extract_to"`
}
