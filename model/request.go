package model

type Request struct {
	QueryString           string            `json:"rawQueryString"`
	Cookies               []string          `json:"cookies"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	Request               RequestContext    `json:"requestContext"`
	Body                  string            `json:"body"`
}

type RequestContext struct {
	DomainName string      `json:"domainName"`
	Http       HttpRequest `json:"http"`
}

type HttpRequest struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Protocal  string `json:"protocal"`
	SourceIP  string `json:"sourceIP"`
	UserAgent string `json:"userAgent"`
}
