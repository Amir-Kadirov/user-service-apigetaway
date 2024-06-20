package handler

import (
	us "backend_course/branch_api_gateway/genproto/user_service"
	"backend_course/branch_api_gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/customer/create [post]
// @Summary		Creates a customer
// @Description	This api creates a customer and returns its id
// @Tags		Customer
// @Accept		json
// @Produce		json
// @Param		customer body user_service.CreateCustomer true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateCustomer(c *gin.Context) {
	req := &us.CreateCustomer{}

	if !validator.ValidateGender(req.Gender) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating gender"), "wrong gender,gender should be (male,female,other)")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidatePhone(req.Phone) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating phone"), "wrong phone")
		return
	}

	if !validator.ValidateGmail(req.Gmail) {
		handleGrpcErrWithDescription(c, h.log, errors.New("error while validating gmail"), "wrong gmail")
		return
	}

	resp, err := h.grpcClient.CustomerService().Create(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get create customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while creating customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router      /v1/customer/getbyid/{id} [get]
// @Summary     Get by id a customer
// @Description This api get by id a customer
// @Tags        Customer
// @Produce     json
// @Param       id path string true "Customer Id"
// @Success     200  {object}  models.ResponseSuccess
// @Failure     400  {object}  models.ResponseError
// @Failure     404  {object}  models.ResponseError
// @Failure     500  {object}  models.ResponseError
func (h *handler) GetByIdCustomer(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		h.log.Error("Customer ID is empty")
		handleGrpcErrWithDescription(c, h.log, errors.New("empty id"), "Error while getting customer by id")
		return
	}

	resp, err := h.grpcClient.CustomerService().GetByID(c.Request.Context(), &us.CustomerPrimaryKey{Id: id})
	if err != nil {
		fmt.Errorf("Error while getting customer by id: %v", err)
		handleGrpcErrWithDescription(c, h.log, err, "Error while getting customer by id")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/getlist [get]
// @Summary		Get list a customer
// @Description	This api get list a customer
// @Tags		Customer
// @Produce		json
// @Param       limit   query int64  false "Limit"
// @Param       offset  query int64  false "Offset"
// @Param       search  query string false "Search gender"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetListCustomer(c *gin.Context) {
	req := &us.GetListCustomerRequest{}

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

	resp, err := h.grpcClient.CustomerService().GetList(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while get list", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while getting list customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/update [PUT]
// @Summary		Update a customer
// @Description	This API updates a customer
// @Tags		Customer
// @Accept		json
// @Produce		json
// @Param		customer body user_service.UpdateCustomerRequest true "Customer object to update"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateCustomer(c *gin.Context) {
	req := &us.UpdateCustomerRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerService().Update(c.Request.Context(), req)
	if err != nil {
		fmt.Errorf("error while update customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while updating customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/customer/delete [delete]
// @Summary		delete a customer
// @Description	This api delete a customer
// @Tags		Customer
// Accept		json
// @Produce		json
// @Param		customer body user_service.CustomerPrimaryKey true "customer"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteCustomer(c *gin.Context) {
	id := &us.CustomerPrimaryKey{}

	if err := c.ShouldBindJSON(&id); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.CustomerService().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Errorf("error while delete", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}
