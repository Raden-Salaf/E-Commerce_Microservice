package response

// Body adalah format response standar semua service.
type Body struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// OK membuat response sukses.
func OK(data any) Body {
	return Body{Success: true, Message: "ok", Data: data}
}

// Fail membuat response gagal.
func Fail(message string) Body {
	return Body{Success: false, Message: message, Data: nil}
}
