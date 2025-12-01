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

// ========================
// FIND ALL BY SUBCATEGORY
// ========================
func FindSubSubCategory(c *gin.Context) {
	subcategoryId := c.Param("subcategoryId")

	// Validasi SubCategory exist
	var subcategory models.SubCategory
	if err := database.DB.First(&subcategory, subcategoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var subsubcategories []models.SubSubCategory
	database.DB.
		Preload("SubCategory").
		Where("sub_category_id = ?", subcategoryId).
		Find(&subsubcategories)

	var response []structs.SubSubCategoryResponse
	for _, ssc := range subsubcategories {
		response = append(response, structs.SubSubCategoryResponse{
			Id:              ssc.Id,
			Name:            ssc.Name,
			Slug:            ssc.Slug,
			SubCategoryId:   ssc.SubCategoryId,
			SubCategoryName: ssc.SubCategory.Name,
			CreatedAt:       ssc.CreatedAt.String(),
			UpdatedAt:       ssc.UpdatedAt.String(),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List SubSubCategory",
		Data:    response,
	})
}

// ========================
// FIND BY ID
// ========================
func FindSubSubCategoryById(c *gin.Context) {
	id := c.Param("subsubcategoryId")

	var ssc models.SubSubCategory
	if err := database.DB.Preload("SubCategory").First(&ssc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubSubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	resp := structs.SubSubCategoryResponse{
		Id:              ssc.Id,
		Name:            ssc.Name,
		Slug:            ssc.Slug,
		SubCategoryId:   ssc.SubCategoryId,
		SubCategoryName: ssc.SubCategory.Name,
		CreatedAt:       ssc.CreatedAt.String(),
		UpdatedAt:       ssc.UpdatedAt.String(),
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "SubSubCategory Found",
		Data:    resp,
	})
}

// ========================
// CREATE
// ========================
func CreateSubSubCategory(c *gin.Context) {
	subcategoryId := c.Param("subcategoryId")

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

	var req structs.SubSubCategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	data := models.SubSubCategory{
		Name:          req.Name,
		Slug:          slug.Make(req.Name),
		SubCategoryId: subcategory.Id,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Create SubSubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "SubSubCategory Created Successfully",
		Data:    data,
	})
}

// ========================
// UPDATE
// ========================
func UpdateSubSubCategory(c *gin.Context) {
	subcategoryId := c.Param("subcategoryId")
	subsubcategoryId := c.Param("subsubcategoryId")

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

	// Validasi subsubcategory
	var ssc models.SubSubCategory
	if err := database.DB.First(&ssc, subsubcategoryId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubSubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req structs.SubSubCategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Update nama + slug
	ssc.Name = req.Name
	ssc.Slug = slug.Make(req.Name)

	// Optional pindah subcategory
	if req.SubCategoryId != 0 {
		ssc.SubCategoryId = req.SubCategoryId
	}

	if err := database.DB.Save(&ssc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Update SubSubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "SubSubCategory Updated Successfully",
		Data:    ssc,
	})
}

// ========================
// DELETE
// ========================
func DeleteSubSubCategory(c *gin.Context) {
	id := c.Param("subsubcategoryId")

	var ssc models.SubSubCategory
	if err := database.DB.First(&ssc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "SubSubCategory Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&ssc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Delete SubSubCategory",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Deleted SubSubCategory Successfully",
		Data:    nil,
	})
}
