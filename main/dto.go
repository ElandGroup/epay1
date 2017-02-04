package main

type (
	APIResult struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Error   APIError    `json:"error"`
	}

	APIError struct {
		Code    string      `json:"code"`
		Details interface{} `json:"details"`
		Message string      `json:"message"`
	}

	// ServiceResult struct {
	// 	Result  interface{} `json:"result"`
	// 	Success string      `json:"success"`
	// 	Error   APIError    `json:"error"`
	// }
)
