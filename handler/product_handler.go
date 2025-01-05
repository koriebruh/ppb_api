package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"koriebruh/uas-ppb/domain"
	"koriebruh/uas-ppb/dto"
	"koriebruh/uas-ppb/helper"
)

type ProductHandler struct {
	*gorm.DB
	*validator.Validate
}

func NewProductHandler(DB *gorm.DB, validate *validator.Validate) *ProductHandler {
	return &ProductHandler{DB: DB, Validate: validate}
}

type ProductHandlerImpl interface {
	CreateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	FindAllProduct(ctx *fiber.Ctx) error
	FindByIdProduct(ctx *fiber.Ctx) error
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var req dto.ProdukRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}
	//VALIDASI REQUEST
	if err := h.Validate.Struct(req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	newProduct := domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageUrl:    req.ImageUrl,
	}

	if err := h.Create(&newProduct).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success create product",
	})
}

func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	if err := h.DB.Where("id = ?", params).Delete(&domain.Product{}).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success delete product",
	})
}

func (h ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	var existingProduct domain.Product
	if err := h.DB.First(&existingProduct, params).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	var req dto.ProdukRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	if err := h.Validate.Struct(req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	existingProduct.Name = req.Name
	existingProduct.Description = req.Description
	existingProduct.Price = req.Price
	existingProduct.ImageUrl = req.ImageUrl

	if err := h.DB.Save(&existingProduct).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success update product",
	})
}

func (h ProductHandler) FindAllProduct(ctx *fiber.Ctx) error {
	var products []domain.Product

	if err := h.DB.Find(&products).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, products)
}

func (h ProductHandler) FindByIdProduct(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	var product domain.Product
	if err := h.DB.First(&product, params).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, product)
}
