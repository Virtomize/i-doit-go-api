package goidoit

// globals
var (
	id              int
	debug, insecure bool = false, false
)

// API struct used for implementing the apiMethods interface
type API struct {
	URL, APIkey, Username, Password, SessionID string
}

// Request implements i-doit api request structure
type Request struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// Response implements i-doit api response structure
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   IdoitError  `json:"error"`
}

// GenericResponse implements a generic i-doit api response structure
// the map is used to handle type assertions
type GenericResponse struct {
	Jsonrpc string
	Result  []map[string]interface{}
	Error   IdoitError
}

// IdoitError implements i-doit api error structure
type IdoitError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// APIkey used for requests
type APIkey struct {
	APIkey string `json:"apikey"`
}

// F1 implements an object filter type int or []int
type F1 struct {
	Data []int `json:"ids"`
}

// F2 implements an Object filter type string
type F2 struct {
	Data string `json:"title"`
}

// OF1 implements a more complex object type filter
// for ids and type
type OF1 struct {
	Title []int  `json:"ids"`
	Type  string `json:"type"`
}

// OF2 implements a more complex object type filter
// for title and type
type OF2 struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}

// OSF1 implements object type only filter
type OSF1 struct {
	Type string `json:"type"`
}
