package response

type AcceptResponse struct {
	HttpCode int
	Message  string `json:"message"`
}
