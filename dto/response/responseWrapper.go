package response

type ResponseWrapper[T any] struct {
	Result  *T             `json:"result,omitempty"`
	Success bool           `json:"success"`
	Meta    *MetaDto       `json:"meta,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

// ToSuccessResponseWrapper creates a successful response wrapper
func ToSuccessResponseWrapper[T any](result T) *ResponseWrapper[T] {
	return &ResponseWrapper[T]{
		Result:  &result,
		Success: true,
	}
}

// ToSuccessResponseWrapperWithMeta creates a successful response wrapper with metadata
func ToSuccessResponseWrapperWithMeta[T any](result T, meta MetaDto) *ResponseWrapper[T] {
	return &ResponseWrapper[T]{
		Result:  &result,
		Success: true,
		Meta:    &meta,
	}
}

// ToErrorResponseWrapper creates an error response wrapper
func ToErrorResponseWrapper[T any](errorResponse ErrorResponse) *ResponseWrapper[T] {
	return &ResponseWrapper[T]{
		Success: false,
		Error:   &errorResponse,
	}
}
