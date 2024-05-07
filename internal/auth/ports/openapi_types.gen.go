// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package ports

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code           *string `json:"code,omitempty"`
	InnerException *struct {
		StackTrace *string `json:"stack_trace,omitempty"`
	} `json:"inner_exception,omitempty"`
	Message *string `json:"message,omitempty"`
}

// MetricsData defines model for MetricsData.
type MetricsData struct {
	URL   *string `json:"URL,omitempty"`
	Count *int    `json:"count,omitempty"`
}

// MetricsResponse defines model for MetricsResponse.
type MetricsResponse struct {
	MetricsData *[]MetricsData `json:"metrics_data,omitempty"`
}

// UrlShortnerResponse defines model for UrlShortnerResponse.
type UrlShortnerResponse struct {
	ShortenUrl *string `json:"shorten_url,omitempty"`
}

// MetricsResponseV1 defines model for MetricsResponseV1.
type MetricsResponseV1 struct {
	Data  *MetricsResponse `json:"data,omitempty"`
	Error *ErrorResponse   `json:"error,omitempty"`
}

// PingResponse defines model for PingResponse.
type PingResponse struct {
	Message *string `json:"message,omitempty"`
}

// UrlShortnerResponseV1 defines model for UrlShortnerResponseV1.
type UrlShortnerResponseV1 struct {
	Data  *UrlShortnerResponse `json:"data,omitempty"`
	Error *ErrorResponse       `json:"error,omitempty"`
}

// MetricsDataJSONBody defines parameters for MetricsData.
type MetricsDataJSONBody = map[string]interface{}

// PostUrlShortnerJSONBody defines parameters for PostUrlShortner.
type PostUrlShortnerJSONBody struct {
	Url *string `json:"url,omitempty"`
}

// PostUrlShortnerParams defines parameters for PostUrlShortner.
type PostUrlShortnerParams struct {
	XTraceId *string `json:"x-trace-id,omitempty"`
}

// GetPingJSONBody defines parameters for GetPing.
type GetPingJSONBody = map[string]interface{}

// MetricsDataJSONRequestBody defines body for MetricsData for application/json ContentType.
type MetricsDataJSONRequestBody = MetricsDataJSONBody

// PostUrlShortnerJSONRequestBody defines body for PostUrlShortner for application/json ContentType.
type PostUrlShortnerJSONRequestBody PostUrlShortnerJSONBody

// GetPingJSONRequestBody defines body for GetPing for application/json ContentType.
type GetPingJSONRequestBody = GetPingJSONBody
