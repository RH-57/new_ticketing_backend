package controllers

import (
	"net/http"
	"ticketing/backend-api/database"
	"ticketing/backend-api/helpers"
	"ticketing/backend-api/models"
	"ticketing/backend-api/structs"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// =======================
// GET LIST SUBCATEGORY BY CATEGORY
// =======================
func FindSubCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	// Cek category exist
	var category models.Category
	if err := database.DB.First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Ambil subcategory
	var subcategories []models.SubCategory
	database.DB.
		Preload("Category").
		Where("category_id = ?", categoryId).
		Find(&subcategories)

	var response []structs.SubCategoryResponse
	for _, sc := range subcategories {
		response = append(response, structs.SubCategoryResponse{
			Id:           sc.Id,
			Name:         sc.Name,
			Slug:         sc.Slug,
			CategoryId:   sc.CategoryId,
			CategoryName: sc.Category.Name,
			CreatedAt:    sc.CreatedAt,
			UpdatedAt:    sc.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Subcategories",
		Data:    response,
	})
}

// =======================
// GET SUBCATEGORY BY ID
// =======================
func FindSubCategoryById(c *gin.Context) {
	id := c.Param("subcategoryId")

	var subcategory models.SubCategory

	// preload category agar CategoryName bisa tampil
	if err := database.DB.Preload("Category").First(&subcategory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	resp := structs.SubCategoryResponse{
		Id:           subcategory.Id,
		Name:         subcategory.Name,
		Slug:         subcategory.Slug,
		CategoryId:   subcategory.CategoryId,
		CategoryName: subcategory.Category.Name,
		CreatedAt:    subcategory.CreatedAt,
		UpdatedAt:    subcategory.UpdatedAt,
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "SubCategory Found",
		Data:    resp,
	})
}

// =======================
// CREATE SUBCATEGORY
// =======================
func CreateSubCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	// Validasi category
	var category models.Category
	if err := database.DB.First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.SubCategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	sub := models.SubCategory{
		Name:       req.Name,
		Slug:       slug.Make(req.Name),
		CategoryId: category.Id,
	}

	if err := database.DB.Create(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Create SubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "SubCategory Created Successfully",
		Data:    sub,
	})
}

// =======================
// UPDATE SUBCATEGORY
// =======================
func UpdateSubCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")
	subcategoryId := c.Param("subcategoryId")

	// Validasi category
	var category models.Category
	if err := database.DB.First(&category, categoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Category Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Validasi subcategory
	var subcategory models.SubCategory
	if err := database.DB.First(&subcategory, subcategoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Validasi body
	var req structs.SubCategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Update name + slug
	subcategory.Name = req.Name
	subcategory.Slug = slug.Make(req.Name)

	// Jika user ingin pindahkan ke category lain, validasi dulu
	if req.CategoryId != 0 {
		var newCat models.Category
		if err := database.DB.First(&newCat, req.CategoryId).Error; err != nil {
			c.JSON(http.StatusBadRequest, structs.ErrorResponse{
				Success: false,
				Message: "New Category Not Found",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
		subcategory.CategoryId = req.CategoryId
	}

	// Simpan perubahan
	if err := database.DB.Save(&subcategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Update SubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Ambil lagi dengan preload agar CategoryName valid
	database.DB.Preload("Category").First(&subcategory)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "SubCategory Updated Successfully",
		Data: structs.SubCategoryResponse{
			Id:           subcategory.Id,
			Name:         subcategory.Name,
			Slug:         subcategory.Slug,
			CategoryId:   subcategory.CategoryId,
			CategoryName: subcategory.Category.Name,
			CreatedAt:    subcategory.CreatedAt,
			UpdatedAt:    subcategory.UpdatedAt,
		},
	})
}

// =======================
// DELETE SUBCATEGORY
// =======================
func DeleteSubCategory(c *gin.Context) {
	subcategoryId := c.Param("subcategoryId")

	var subcategory models.SubCategory

	if err := database.DB.First(&subcategory, subcategoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&subcategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed To Delete SubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Deleted SubCategory Successfully",
		Data:    nil,
	})
}
