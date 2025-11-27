package controllers

import (
	"net/http"
	"ticketing/backend-api/database"
	"ticketing/backend-api/helpers"
	"ticketing/backend-api/models"
	"ticketing/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindBranch(c *gin.Context) {
	var branches []models.Branch

	database.DB.Find(&branches)
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Data Branch",
		Data:    branches,
	})
}

func CreateBranch(c *gin.Context) {
	var req = structs.BranchCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	branch := models.Branch{
		Code: req.Code,
		Name: req.Name,
	}

	if err := database.DB.Create(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Branch Created Successfully",
		Data: structs.BranchResponse{
			Id:        branch.Id,
			Code:      branch.Code,
			Name:      branch.Name,
			CreatedAt: branch.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: branch.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindBranchById(c *gin.Context) {
	id := c.Param("id")

	var branch models.Branch

	if err := database.DB.First(&branch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Branch Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Branch Found",
		Data: structs.BranchResponse{
			Id:        branch.Id,
			Code:      branch.Code,
			Name:      branch.Name,
			CreatedAt: branch.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: branch.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateBranch(c *gin.Context) {
	id := c.Param("id")

	var branch models.Branch

	if err := database.DB.First(&branch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Branch Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.BranchUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var existing models.Branch
	if err := database.DB.Where("code = ? AND id <> ?", req.Code, id).
		First(&existing).Error; err == nil {

		c.JSON(http.StatusConflict, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  map[string]string{"Code": "Code already exists"},
		})
		return
	}

	branch.Code = req.Code
	branch.Name = req.Name

	if err := database.DB.Save(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed To Update Branch",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Update Branch Successfully",
		Data: structs.BranchResponse{
			Id:        branch.Id,
			Code:      branch.Code,
			Name:      branch.Name,
			CreatedAt: branch.CreatedAt.Format("2006-01-02 15:05:05"),
			UpdatedAt: branch.UpdatedAt.Format("2006-01-02 15:05:05"),
		},
	})
}

func DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	var branch models.Branch

	if err := database.DB.First(&branch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Branch Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&branch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed TO Delete Branch",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Deleted Branch Successfully",
		Data:    nil,
	})
}
