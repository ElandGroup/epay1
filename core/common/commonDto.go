package common

type (
	APIResult struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Error   APIError    `json:"error"`
	}

	QueryResult struct {
		TotalCount int         `json:totalCount`
		Items      interface{} `items`
	}

	APIError struct {
		Code    int         `json:"code"`
		Details interface{} `json:"details"`
		Message string      `json:"message"`
	}
)

type (
	APIParam struct {
		SkinCount      int
		MaxResultCount int
		Fields         string
		Sort           string
	}
)
