package controllers

import (
	"net/http"
	"ticketing/backend-api/database"
	"ticketing/backend-api/helpers"
	"ticketing/backend-api/models"
	"ticketing/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindDepartment(c *gin.Context) {
	var departments []models.Department

	database.DB.Preload("Division").Find(&departments)

	var response []structs.DepartmentResponse
	for _, d := range departments {
		response = append(response, structs.DepartmentResponse{
			Id:           d.Id,
			Name:         d.Name,
			DivisionId:   d.DivisionId,
			DivisionName: d.Division.Name,
			CreatedAt:    d.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    d.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Data Department",
		Data:    response,
	})
}

func CreateDepartment(c *gin.Context) {
	var req = structs.DepartmentCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	department := models.Department{
		Name:       req.Name,
		DivisionId: req.DivisionId,
	}

	var division models.Division
	if err := database.DB.First(&division, req.DivisionId).Error; err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Division Not Found",
			Errors:  map[string]string{"division_id": "Invalid Division Id"},
		})
		return
	}

	if err := database.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Internal Server Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Division").First(&department, department.Id)

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Department Created Successfully",
		Data: structs.DepartmentResponse{
			Id:           department.Id,
			Name:         department.Name,
			DivisionId:   department.DivisionId,
			DivisionName: department.Division.Name,
			CreatedAt:    department.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    department.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindDepartmentById(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Department Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Division").First(&department, department.Id)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Department Found",
		Data: structs.DepartmentResponse{
			Id:           department.Id,
			Name:         department.Name,
			DivisionId:   department.DivisionId,
			DivisionName: department.Division.Name,
			CreatedAt:    department.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    department.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Department Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.DepartmentUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	department.Name = req.Name
	department.DivisionId = req.DivisionId

	if err := database.DB.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed To Update Depertment",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	database.DB.Preload("Division").First(&department, department.Id)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Division Updated Successfully",
		Data: structs.DepartmentResponse{
			Id:           department.Id,
			Name:         department.Name,
			DivisionId:   department.DivisionId,
			DivisionName: department.Division.Name,
			CreatedAt:    department.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    department.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	if err := database.DB.First(&department, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Department Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed To Delete Department",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Delete Department Successfully",
		Data:    nil,
	})
}
