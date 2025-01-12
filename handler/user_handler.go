package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"koriebruh/uas-ppb/domain"
	"koriebruh/uas-ppb/dto"
	"koriebruh/uas-ppb/helper"
	"time"
)

type UserHandler struct {
	*gorm.DB
	*validator.Validate
}

func NewUserHandler(DB *gorm.DB, validate *validator.Validate) *UserHandler {
	return &UserHandler{DB: DB, Validate: validate}
}

type UserHandlerImpl interface {
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	AddProductToCart(ctx *fiber.Ctx) error
	GetCartItems(ctx *fiber.Ctx) error
	AddShippingAndGetTotal(ctx *fiber.Ctx) error
	CheckoutAndClearCart(ctx *fiber.Ctx) error
	RemoveUserById(ctx *fiber.Ctx) error
	FindAllUser(ctx *fiber.Ctx) error
	HistoryCheckout(ctx *fiber.Ctx) error
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var req dto.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}
	//VALIDASI REQUEST
	if err := h.Validate.Struct(req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	newUser := domain.User{
		User:     req.User,
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := h.Create(&newUser).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success create user",
	})

}

func (h UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req dto.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}
	//VALIDASI REQUEST
	if err := h.Validate.Struct(req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	id := ctx.Params("id")
	var user domain.User
	if err := h.First(&user, id).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Update user data
	user.User = req.User
	user.Username = req.Username
	user.Password = req.Password
	user.Role = req.Role

	// Save the updated user
	if err := h.Save(&user).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success update user",
	})

}

func (h UserHandler) Login(ctx *fiber.Ctx) error {
	var req dto.UserLoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	if err := h.Validate.Struct(req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	var user domain.User
	if err := h.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"data": user,
	})
}

func (h UserHandler) AddProductToCart(ctx *fiber.Ctx) error {
	var req dto.AddToCartRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Cek apakah user sudah memiliki keranjang aktif
	var cart domain.Cart
	if err := h.Where("user_id = ?", req.UserID).First(&cart).Error; err != nil {
		// Jika belum ada, buatkan keranjang baru
		cart = domain.Cart{
			UserID:  req.UserID,
			Status:  "-",
			Tanggal: time.Now(),
		}
		if err := h.Create(&cart).Error; err != nil {
			return helper.ErrResponse(ctx, err)
		}
	}

	// Cek apakah produk sudah ada dalam keranjang
	var cartItem domain.CartItem
	if err := h.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItem).Error; err == nil {
		// Jika ada, update jumlah dan subtotal
		cartItem.Jumlah += req.Jumlah
		cartItem.Subtotal = float64(cartItem.Jumlah) * cartItem.Harga
		if err := h.Save(&cartItem).Error; err != nil {
			return helper.ErrResponse(ctx, err)
		}
	} else {
		// Jika produk belum ada, tambahkan produk baru ke keranjang
		var product domain.Product
		if err := h.Where("id = ?", req.ProductID).First(&product).Error; err != nil {
			return helper.ErrResponse(ctx, fmt.Errorf("product not found"))
		}

		cartItem = domain.CartItem{
			CartID:    cart.ID,
			ProductID: req.ProductID,
			Jumlah:    req.Jumlah,
			Harga:     product.Price,
			Subtotal:  float64(req.Jumlah) * product.Price,
		}
		if err := h.Create(&cartItem).Error; err != nil {
			return helper.ErrResponse(ctx, err)
		}
	}

	var product domain.Product
	if err := h.Where("id = ?", req.ProductID).First(&product).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("product not found"))
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "Produk berhasil ditambahkan ke keranjang",
		"data": map[string]interface{}{
			"product_name": product.Name,
			"quantity":     cartItem.Jumlah,
			"price":        cartItem.Harga,
			"subtotal":     cartItem.Subtotal,
		},
	})
}
func (h UserHandler) GetCartItems(ctx *fiber.Ctx) error {
	var req dto.GetCartItemsRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Mencari keranjang aktif berdasarkan userID
	var cart domain.Cart
	if err := h.Where("user_id = ?", req.UserID).First(&cart).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("keranjang not found"))
	}

	// Ambil semua barang dalam keranjang
	var cartItems []domain.CartItem
	if err := h.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("keranjang not found"))
	}

	// Initialize a variable to accumulate the total price
	var totalPrice float64

	// Prepare the items data with only necessary fields (name, quantity, price, subtotal)
	var responseItems []map[string]interface{}
	for _, cartItem := range cartItems {
		var product domain.Product
		if err := h.Where("id = ?", cartItem.ProductID).First(&product).Error; err != nil {
			return helper.ErrResponse(ctx, fmt.Errorf("product not found"))
		}

		// Calculate the subtotal for the item
		cartItem.Subtotal = float64(cartItem.Jumlah) * cartItem.Harga

		// Add the subtotal to the total price
		totalPrice += cartItem.Subtotal

		// Append the formatted item response
		responseItems = append(responseItems, map[string]interface{}{
			"product_id":   product.ID,
			"product_name": product.Name,
			"quantity":     cartItem.Jumlah,
			"price":        cartItem.Harga,
			"subtotal":     cartItem.Subtotal,
		})
	}

	// Return the response with the total price
	return helper.SuccessResponse(ctx, map[string]interface{}{
		"items":       responseItems,
		"total_price": totalPrice, // Include total price of all items
	})
}
func (h UserHandler) AddShippingAndGetTotal(ctx *fiber.Ctx) error {
	var req dto.AddShippingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Cek apakah user sudah memiliki keranjang aktif
	var cart domain.Cart
	if err := h.Where("user_id = ?", req.UserID).Preload("CartItems").First(&cart).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("keranjang tidak ditemukan"))
	}

	// Menambahkan ongkir ke keranjang
	shipping := domain.Shipping{
		CartID:      cart.ID,
		KotaAsal:    req.KotaAsal,
		KotaTujuan:  req.KotaTujuan,
		BiayaOngkir: req.BiayaOngkir,
		Weight:      req.Weight,
	}

	// Simpan ongkir
	if err := h.Save(&shipping).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Menghitung total barang dalam keranjang
	var totalBarang float64
	var responseItems []map[string]interface{}
	for _, item := range cart.CartItems {
		var product domain.Product
		if err := h.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			return helper.ErrResponse(ctx, fmt.Errorf("produk tidak ditemukan"))
		}

		// Hitung subtotal per item
		item.Subtotal = float64(item.Jumlah) * item.Harga
		totalBarang += item.Subtotal

		// Menambahkan data barang ke responseItems
		responseItems = append(responseItems, map[string]interface{}{
			"product_name": product.Name,
			"quantity":     item.Jumlah,
			"price":        item.Harga,
			"subtotal":     item.Subtotal,
		})
	}

	// Menghitung total harga keranjang + ongkir
	cart.Total = totalBarang + shipping.BiayaOngkir
	if err := h.Save(&cart).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Mengembalikan hasil akhir dengan total, ongkir, dan informasi barang
	return helper.SuccessResponse(ctx, map[string]interface{}{
		"total":  cart.Total,           // Total harga barang + ongkir
		"ongkir": shipping.BiayaOngkir, // Biaya ongkir
		"items":  responseItems,        // Informasi barang dalam keranjang
		"shipping": map[string]interface{}{ //detail onkir
			"city_from":     shipping.KotaAsal,
			"city_to":       shipping.KotaTujuan,
			"shipping_cost": shipping.BiayaOngkir,
			"weight":        shipping.Weight,
		},
	})
}

func (h UserHandler) CheckoutAndClearCart(ctx *fiber.Ctx) error {
	var req dto.CheckoutRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Mencari keranjang berdasarkan userID tanpa memperhatikan status 'aktif'
	var cart domain.Cart
	if err := h.Where("user_id = ?", req.UserID).First(&cart).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("keranjang tidak ditemukan"))
	}

	// Validasi pembayaran (misalnya memeriksa status pembayaran)
	if !req.IsPaid {
		return helper.ErrResponse(ctx, fmt.Errorf("pembayaran belum dilakukan"))
	}

	// Hapus barang dari keranjang
	if err := h.Where("cart_id = ?", cart.ID).Delete(&domain.CartItem{}).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Mengembalikan response bahwa transaksi selesai
	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "Pembayaran berhasil, keranjang telah selesai",
	})
}

func (h UserHandler) RemoveProductFromCart(ctx *fiber.Ctx) error {
	var req dto.RemoveFromCartRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Cek apakah user sudah memiliki keranjang aktif
	var cart domain.Cart
	if err := h.Where("user_id = ?", req.UserID).Preload("CartItems").First(&cart).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("keranjang tidak ditemukan"))
	}

	// Cek apakah produk ada dalam keranjang
	var cartItem domain.CartItem
	if err := h.Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).First(&cartItem).Error; err != nil {
		return helper.ErrResponse(ctx, fmt.Errorf("produk tidak ditemukan dalam keranjang"))
	}

	// Hapus barang dari keranjang
	if err := h.Delete(&cartItem).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	// Mengembalikan response sukses
	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "Produk berhasil dihapus dari keranjang",
	})
}

func (h UserHandler) RemoveUserById(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	if err := h.DB.Where("id = ?", params).Delete(&domain.User{}).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, map[string]interface{}{
		"message": "success delete consumer",
	})
}

func (h UserHandler) FindAllUser(ctx *fiber.Ctx) error {
	var users []domain.User

	if err := h.DB.Where("role != ?", "admin").Find(&users).Error; err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, users)
}

func (h UserHandler) HistoryCheckout(ctx *fiber.Ctx) error {
	var purchaseHistory []dto.PurchaseHistory

	err := h.DB.Table("cart_items").
		Select(`
			products.name AS name,
			cart_items.jumlah AS jumlah,
			products.price AS price,
			cart_items.subtotal AS subtotal,
			cart_items.deleted_at AS buy_time
		`).
		Joins("LEFT JOIN products ON cart_items.product_id = products.id").
		Where("cart_items.deleted_at IS NOT NULL"). // Hanya history yang memiliki buy_time
		Scan(&purchaseHistory).Error

	if err != nil {
		return helper.ErrResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, purchaseHistory)
}
