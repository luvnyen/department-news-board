package controller

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luvnyen/department-news-board/pkg/dto"
	"github.com/luvnyen/department-news-board/pkg/utils"
	"github.com/luvnyen/department-news-board/service"
)

type NewsController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	InsertNews(ctx *gin.Context)
	UpdateNews(ctx *gin.Context)
	DeleteNews(ctx *gin.Context)
	DownloadNews(ctx *gin.Context)
}

type newsController struct {
	newsService service.NewsService
	jwtService  service.JWTService
}

func NewNewsController(newsService service.NewsService, jwtService service.JWTService) NewsController {
	return &newsController{newsService: newsService, jwtService: jwtService}
}

func (db *newsController) All(ctx *gin.Context) {
	_, err := db.newsService.ArchiveNews()
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	news, err := db.newsService.All()
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	response := utils.BuildResponse(true, "Successfully get all news", news)
	ctx.JSON(200, response)
}

func (db *newsController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := db.newsService.FindByID(id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	response := utils.BuildResponse(true, "Successfully get news", res)
	ctx.JSON(200, response)
}

func (db *newsController) InsertNews(ctx *gin.Context) {
	var newsDTO dto.NewsDTO
	errDTO := ctx.ShouldBind(&newsDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request DTO", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	_, err := db.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", "Invalid token", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	// fileBytes, err := file.Open()
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }
	// defer fileBytes.Close()
	// fileBase64, err := ioutil.ReadAll(fileBytes)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }

	// fileBase64String := base64.StdEncoding.EncodeToString(fileBase64)

	extension := filepath.Ext(file.Filename)
	if extension != ".pdf" {
		response := utils.BuildErrorResponse("Failed to process request", "Invalid file extension", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	newFileName := uuid.New().String() + extension

	absPath, err := filepath.Abs("./cdn/news/" + newFileName)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	res, err := db.newsService.InsertNews(newsDTO, newFileName)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := ctx.SaveUploadedFile(file, absPath); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	response := utils.BuildResponse(true, "Successfully insert news", res)
	ctx.JSON(200, response)
}

func (db *newsController) UpdateNews(ctx *gin.Context) {
	id := ctx.Param("id")

	var newsDTO dto.NewsDTO
	errDTO := ctx.ShouldBind(&newsDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	_, err := db.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", "Invalid token", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	// fileBytes, err := file.Open()
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }
	// defer fileBytes.Close()
	// fileBase64, err := ioutil.ReadAll(fileBytes)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }

	// fileBase64String := base64.StdEncoding.EncodeToString(fileBase64)

	extension := filepath.Ext(file.Filename)
	if extension != ".pdf" {
		response := utils.BuildErrorResponse("Failed to process request", "Invalid file extension", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	newFileName := uuid.New().String() + extension

	absPath, err := filepath.Abs("./cdn/news/" + newFileName)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	oldRes, err := db.newsService.FindByID(id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := os.Remove("./cdn/news/" + oldRes.File); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	res, err := db.newsService.UpdateNews(id, newsDTO, newFileName)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if err := ctx.SaveUploadedFile(file, absPath); err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	response := utils.BuildResponse(true, "Successfully update news", res)
	ctx.JSON(200, response)
}

func (db *newsController) DeleteNews(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, errValidate := db.jwtService.ValidateToken(authHeader)
	if errValidate != nil {
		response := utils.BuildErrorResponse("Failed to process request", "Invalid token", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	id := ctx.Param("id")

	oldRes, errFetch := db.newsService.FindByID(id)
	if errFetch != nil {
		response := utils.BuildErrorResponse("Failed to process request", errFetch.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	if errRemove := os.Remove("./cdn/news/" + oldRes.File); errRemove != nil {
		response := utils.BuildErrorResponse("Failed to process request", errRemove.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	errDelete := db.newsService.DeleteNews(id)
	if errDelete != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDelete.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	response := utils.BuildResponse(true, "Successfully delete news", utils.EmptyObj{})
	ctx.JSON(200, response)
}

func (db *newsController) DownloadNews(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := db.newsService.FindByID(id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	absPath, err := filepath.Abs("./cdn/news/" + res.File)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(400, response)
		return
	}

	ctx.File(absPath)

	// fileBase64, err := base64.RawStdEncoding.DecodeString(res.File)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }

	// fileName := uuid.New().String() + ".pdf"
	// absPath, err := filepath.Abs("./cdn/news/" + fileName)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }

	// file, err := os.Create(absPath)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }
	// defer file.Close()

	// _, err = file.Write(fileBase64)
	// if err != nil {
	// 	response := utils.BuildErrorResponse("Failed to process request", err.Error(), utils.EmptyObj{})
	// 	ctx.AbortWithStatusJSON(400, response)
	// 	return
	// }

	// ctx.Header("Content-Description", "File Transfer")
	// ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	// ctx.Header("Content-Type", "application/pdf")
	// ctx.Header("Content-Transfer-Encoding", "binary")
	// ctx.File(absPath)
}
