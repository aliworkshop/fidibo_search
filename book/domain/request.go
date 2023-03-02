package domain

type SearchRequest struct {
	Keyword string `json:"keyword" query:"keyword" form:"keyword" binding:"required"`
}
