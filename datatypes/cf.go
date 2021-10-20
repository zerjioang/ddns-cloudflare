package datatypes

// ErrorResponse Cloudflare API error response message format
type ErrorResponse struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   interface{}   `json:"result"`
}
