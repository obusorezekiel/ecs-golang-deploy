package response

// Response represents the structure of a response returned by the API.
// It typically includes fields for status, message, and data.
type Response struct {
 Code   int         `json:"code"`
 Status string      `json:"status"`
 Data   interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response with a message.
type ErrorResponse struct {
 Code    int         `json:"code"`
 Message string      `json:"message"`
}

// StudentResponse represents the response for a student entity.
type StudentResponse struct {
 ID     int    `json:"ID"`
 Name   string `json:"name"`
 Description string `json:"description"`
}