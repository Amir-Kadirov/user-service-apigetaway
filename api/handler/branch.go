package handler

import (
	us "backend_course/branch_api_gateway/genproto/user_service"
	"backend_course/branch_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/branch/create [post]
// @Summary		Creates a Branch
// @Description	This api creates a Branch and returns its id
// @Tags		Branch
// @Accept		json
// @Produce		json
// @Param		Branch body user_service.CreateBranch true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateBranch(c *gin.Context) {
	var req us.CreateBranch

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "error while validating phone body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while create branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/getbyid/{id} [get]	
// @Summary		Get by id a Branch
// @Description	This api get by id a Branch
// @Tags		Branch
// @Produce		json
// @Param		id path string true "Branch id"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.grpcClient.BranchBranch().GetByID(c.Request.Context(), &us.BranchPrimaryKey{Id: id})
	if err != nil {
		log.Fatal("error while get by id", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while get by id branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/getlist [get]
// @Summary		Get list a branch
// @Description	This api get list a branch
// @Tags		Branch
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListBranch(c *gin.Context) {
	req := &us.GetListBranchRequest{}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	search := c.Query("search")

	if limitStr != "" {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			fmt.Errorf("error while parse limit", err)
			handleGrpcErrWithDescription(c, h.log, err, "error while parse limit")
			return
		}
		req.Limit = limit
	}

	if offsetStr != "" {
		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			fmt.Errorf("error while parse offset", err)
			handleGrpcErrWithDescription(c, h.log, err, "error while parse offset")
			return
		}
		req.Offset = offset
	}

	if search != "" {
		req.Search = search
	}

	resp, err := h.grpcClient.BranchBranch().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/updatebranch [PUT]
// @Summary		Update a Branch
// @Description	This API updates a Branch
// @Tags		Branch
// @Accept		json
// @Produce		json
// @Param		Branch body user_service.UpdateBranchRequest true "Branch object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) Update(c *gin.Context) {
	req := &us.UpdateBranchRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Update(c.Request.Context(), req)
	if err != nil {
		log.Fatal("error while update Branch", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/branch/delete [delete]
// @Summary		delete a Branch
// @Description	This api delete a Branch
// @Tags		Branch
// Accept		json
// @Produce		json
// @Param		Branch body user_service.BranchPrimaryKey true "Branch"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) Delete(c *gin.Context) {
	id := &us.BranchPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.BranchBranch().Delete(c.Request.Context(), id)
	if err != nil {
		log.Fatal("error while get delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting branch")
		return
	}

	c.JSON(http.StatusOK, resp)
}
