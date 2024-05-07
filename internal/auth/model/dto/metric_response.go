package dto

type MetricsResponseV1 struct {
	Data  MetricsResponse `json:"data,omitempty"`
	Error ErrorResponse   `json:"error,omitempty"`
}

type MetricsData struct {
	URL string `json:"URL,omitempty"`
}

// MetricsResponse defines model for MetricsResponse.
type MetricsResponse struct {
	MetricsData []MetricsData `json:"metrics_data,omitempty"`
}
