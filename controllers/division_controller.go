package controllers

import (
	"net/http"
	"ticketing/backend-api/database"
	"ticketing/backend-api/helpers"
	"ticketing/backend-api/models"
	"ticketing/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindDivision(c *gin.Context) {
	var divisions []models.Division

	database.DB.Preload("Branch").Find(&divisions)

	var response []structs.DivisionResponse
	for _, d := range divisions {
		response = append(response, structs.DivisionResponse{
			Id:         d.Id,
			Name:       d.Name,
			BranchId:   d.BranchId,
			BranchCode: d.Branch.Code,
			CreatedAt:  d.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  d.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Data Division",
		Data:    response,
	})
}

func CreateDivision(c *gin.Context) {
	var req = structs.DivisionCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	division := models.Division{
		Name:     req.Name,
		BranchId: req.BranchId,
	}

	// Validate if branch exists
	var branch models.Branch
	if err := database.DB.First(&branch, req.BranchId).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Branch Not Found",
			Errors:  map[string]string{"branch_id": "Invalid branch_id"},
		})
		return
	}

	if err := database.DB.Create(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Branch").First(&division, division.Id)

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Division Created Successfully",
		Data: structs.DivisionResponse{
			Id:         division.Id,
			Name:       division.Name,
			BranchId:   division.BranchId,
			BranchCode: division.Branch.Code,
			CreatedAt:  division.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  division.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindDivisionByid(c *gin.Context) {
	id := c.Param("id")

	var division models.Division

	if err := database.DB.First(&division, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Division Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Branch").First(&division, division.Id)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Division Found",
		Data: structs.DivisionResponse{
			Id:         division.Id,
			Name:       division.Name,
			BranchId:   division.BranchId,
			BranchCode: division.Branch.Code,
		},
	})
}

func UpdateDivision(c *gin.Context) {
	id := c.Param("id")

	var division models.Division

	if err := database.DB.First(&division, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Division Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.DivisionUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	division.Name = req.Name
	division.BranchId = req.BranchId

	if err := database.DB.Save(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to Update Division",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Branch").First(&division, division.Id)

	// Response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Division Updated Successfully",
		Data: structs.DivisionResponse{
			Id:         division.Id,
			Name:       division.Name,
			BranchId:   division.BranchId,
			BranchCode: division.Branch.Code,
			CreatedAt:  division.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  division.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteDivision(c *gin.Context) {
	id := c.Param("id")

	var division models.Division

	if err := database.DB.First(&division, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Division Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&division).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed To Delete Division",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Delete Division Successfully",
		Data:    nil,
	})
}
