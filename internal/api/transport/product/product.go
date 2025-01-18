package product

import (
	"github.com/bigxxby/digital-travel-test/internal/api/service/product"
	"github.com/bigxxby/digital-travel-test/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IProductController interface {
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetAllProducts(ctx *gin.Context)
}

type ProductController struct {
	ProductService product.IProductService
}

func NewProductController(productService product.IProductService) IProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//get user_id from token
	userId, _ := ctx.Get("user_id")
	if userId == nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	//parse
	userIdStr, ok := userId.(string)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Failed to parse user ID"})
		return
	}
	//create uuid
	userIdUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdProduct, status, err := pc.ProductService.CreateProduct(&userIdUUID, product)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"data": createdProduct})
}

func (pc ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//get user_id from token
	userId, _ := ctx.Get("user_id")
	if userId == nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	//parse
	userIdStr, ok := userId.(string)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Failed to parse user ID"})
		return
	}
	//create uuid
	userIdUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, status, err := pc.ProductService.UpdateProduct(&userIdUUID, product)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"data": updatedProduct})
}
func (pc ProductController) DeleteProduct(ctx *gin.Context) {
	productId, err := uuid.Parse(ctx.Param("productId"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//get user_id from token
	userId, _ := ctx.Get("user_id")
	if userId == nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	//parse
	userIdStr, ok := userId.(string)
	if !ok {
		ctx.JSON(400, gin.H{"error": "Failed to parse user ID"})
		return
	}
	//create uuid
	userIdUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	status, err := pc.ProductService.DeleteProduct(&productId, &userIdUUID)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"data": "Product deleted"})
}
func (pc ProductController) GetProductById(ctx *gin.Context) {
	productId, err := uuid.Parse(ctx.Param("productId"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, status, err := pc.ProductService.GetProductById(&productId)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"data": product})
}
func (pc ProductController) GetAllProducts(ctx *gin.Context) {
	products, status, err := pc.ProductService.GetAllProducts()
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, gin.H{"data": products})
}
