package handler

import (
	"fmt"
	"math"
	"net/http"
	"project_pertama/formatter"
	"project_pertama/helper"
	"project_pertama/input"
	"project_pertama/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Save(c *gin.Context) {
	var input input.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("error", "Gagal menyimpan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdUser, err := h.userService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("error", "Gagal menyimpan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatterUser := formatter.FormatUser(createdUser)
	response := helper.ApiResponse("success", "Berhasil menyimpan data user", http.StatusOK, formatterUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadProfilePicture(c *gin.Context) {
	var inputID input.UserInputID
	id := c.PostForm("id")
	userID, err := strconv.Atoi(id)
	inputID.ID = userID
	if err != nil {
		errorMessage := gin.H{
			"is_uploaded": false,
		}
		response := helper.ApiResponse("error", err.Error(), http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("profile_picture")

	if err != nil {
		errorMessage := gin.H{
			"is_uploaded": false,
		}
		response := helper.ApiResponse("error", "Gagal mengupload profile picture user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	lokasiFile := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, lokasiFile)
	if err != nil {
		errorMessage := gin.H{
			"is_uploaded": false,
		}
		response := helper.ApiResponse("error", "Gagal mengupload profile picture user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedUser, err := h.userService.UploadProfilePicture(userID, lokasiFile)
	if err != nil {
		errorMessage := gin.H{
			"is_uploaded": false,
		}
		response := helper.ApiResponse("error", "Gagal mengupload profile picture user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("success", "Berhasil upload", http.StatusOK, formatter.FormatUser(updatedUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	// http://localhost:8080/users/1
	var inputID input.UserInputID
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mengupdate user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	existingUser, err := h.userService.FindByID(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mengupdate user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if existingUser.ID == 0 {
		errorMessage := gin.H{"error_message": "user tidak ditemukan di database anda"}
		response := helper.ApiResponse("error", "Gagal mengupdate user", http.StatusOK, errorMessage)
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	} else {
		var inputData input.UserInput
		err = c.ShouldBindJSON(&inputData)
		if err != nil {
			errorMessage := gin.H{"error_message": err.Error()}
			response := helper.ApiResponse("error", "Gagal mengupdate user", http.StatusBadRequest, errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		updatedUser, err := h.userService.Update(inputID, inputData)
		if err != nil {
			errorMessage := gin.H{"error_message": err.Error()}
			response := helper.ApiResponse("error", "Gagal mengupdate user", http.StatusBadRequest, errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ApiResponse("success", "Berhasil mengupdate data user", http.StatusOK, formatter.FormatUser(updatedUser))
		c.JSON(http.StatusOK, response)
	}
}

func (h *userHandler) GetAllUser(c *gin.Context) {
	search := c.Query("search")
	page := c.Query("page")
	size := c.Query("size")

	convertedPage, _ := strconv.Atoi(page)
	convertedSize, _ := strconv.Atoi(size)

	usersData, err := h.userService.FindAll(search, convertedPage, convertedSize)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mendapatkan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	totalData, err := h.userService.TotalFetchData(search, convertedPage, convertedSize)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mendapatkan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	totalPage := math.Ceil(float64(totalData) / float64(convertedSize))
	currentPage := convertedPage + 1

	response := helper.ServerSideResponses(totalData, int(totalPage), currentPage, formatter.FormatUsers(usersData))
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) GetByID(c *gin.Context) {
	var inputID input.UserInputID
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mendapatkan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userService.FindByID(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal mendapatkan data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("success", "Berhasil mendapatkan data user", http.StatusOK, formatter.FormatUser(user))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	var inputID input.UserInputID
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal menghapus data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedUser, err := h.userService.DeleteUser(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"error_message": err.Error()}
		response := helper.ApiResponse("error", "Gagal menghapus data user", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	deleteStatus := gin.H{"is_deleted": deletedUser}
	response := helper.ApiResponse("success", "Berhasil menghapus data user", http.StatusOK, deleteStatus)
	c.JSON(http.StatusOK, response)

}
