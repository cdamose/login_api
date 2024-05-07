package dto

type ErrorResponse struct {
	Code           *string `json:"code,omitempty"`
	InnerException *struct {
		StackTrace *string `json:"stack_trace,omitempty"`
	} `json:"inner_exception,omitempty"`
	Message *string `json:"message,omitempty"`
}

type UrlShortnerResponse struct {
	ShortenUrl *string `json:"shorten_url,omitempty"`
}

type UrlShortnerResponseV1 struct {
	Data  *UrlShortnerResponse `json:"data,omitempty"`
	Error *ErrorResponse       `json:"error,omitempty"`
}
