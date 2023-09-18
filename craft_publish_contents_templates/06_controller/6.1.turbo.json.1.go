package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *publishcontentstemplates) DeleteCraftV1PublishContentsTemplatesId(ctx *gin.Context, s1 string) {
	tenantID := ctx.GetHeader("Tenant-ID")

	err := c.u.Delete(ctx, tenantID, s1)
	if err != nil {
		logger.Error(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete publish contents templates")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Publish contents templates deleted successfully",
	})
}

func (c *publishcontentstemplates) DeleteCraftV1PublishContentsTemplatesId_want(ctx *gin.Context, s1 string) {
	au, err := common.GetAuthenticatedUser(ctx)
	if err != nil {
		detail := err.Error()
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
			Msg:    "Get Authenticated User failed",
			Code:   "internal_server_error",
			Detail: &detail,
		})
		return
	}

	err = c.u.Delete(ctx, au.TenantId, s1)
	if err != nil {
		logger.Error(err)
		if err == usecase.ErrNotFoundPublishContentsTemplatesID {
			ctx.JSON(
				http.StatusNotFound,
				common.NewBadRequest(
					fmt.Sprintf("Not Found Publish Contents Templates ID Error: %s", s1),
				),
			)
			return
		}
		errorMsg := err.Error()
		ctx.JSON(http.StatusInternalServerError, common.InternalServerError{
			Code:   "internal_server_error",
			Msg:    "Internal Server Error",
			Detail: &errorMsg,
		})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
