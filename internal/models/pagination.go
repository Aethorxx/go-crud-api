package models

// PaginationParams представляет параметры пагинации для списков
type PaginationParams struct {
	Page   int `form:"page" binding:"min=1"`
	Limit  int `form:"limit" binding:"min=1,max=100"`
	MinAge int `form:"min_age" binding:"min=0"`
	MaxAge int `form:"max_age" binding:"min=0"`
}

// PaginatedResponse представляет ответ с пагинацией
type PaginatedResponse struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	Users      interface{} `json:"users"`
}
