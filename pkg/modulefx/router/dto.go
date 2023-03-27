package router

// @Description Response when succeed
type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

// @Description Response when error hanppened
type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code       int            `json:"code"`
	Reason     string         `json:"reason"`
	Description string 		  `json:"description"`
}


