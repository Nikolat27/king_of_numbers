package types

type ErrorResponse struct {
	Type   string
	Detail string
}

func NewErrorResponse(typ, detail any) *ErrorResponse {
	var (
		typString    string
		detailString string
	)

	switch v := typ.(type) {
	case string:
		typString = v
	case error:
		typString = v.Error()
	default:
		typString = "unknown 'typ' data type"
	}

	switch v := detail.(type) {
	case string:
		detailString = v
	case error:
		detailString = v.Error()
	default:
		detailString = "unknown 'detail' data type"
	}

	return &ErrorResponse{
		Type:   typString,
		Detail: detailString,
	}
}
