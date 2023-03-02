package delivery

import (
	"fidibo/book/domain"
	"fidibo/helper"
	"github.com/gin-gonic/gin"
)

type searchHandler struct {
	uc domain.BookUc
}

func SearchBookDelivery(uc domain.BookUc) gin.HandlerFunc {
	h := &searchHandler{
		uc: uc,
	}
	return h.SearchBook
}

func (h *searchHandler) SearchBook(c *gin.Context) {
	var req domain.SearchRequest
	if err := c.BindQuery(&req); err != nil {
		helper.ErrorResponse(c, helper.BadRequest, err)
		return
	}
	res, err := h.uc.Search(req.Keyword)
	if err != nil {
		helper.ErrorResponse(c, helper.InternalServerError, err)
		return
	}
	helper.SuccessResponse(c, map[string]any{
		"items": res,
	})
}
