package handler

import (
	minioHandler "github.com/1abobik1/Cloud-Storage/file_upload_service/internal/handler/minio_handler"
	"github.com/1abobik1/Cloud-Storage/file_upload_service/internal/minio"
	"github.com/gin-gonic/gin"
)

// Services структура всех сервисов, которые используются в хендлерах
// Это нужно чтобы мы могли использовать внутри хендлеров эти самые сервисы
type Services struct {
	minioService minio.Client // Сервис у нас только один - minio, мы планируем его использовать, поэтому передаем
}

// Handlers структура всех хендлеров, которые используются для обозначения действия в роутах
type Handlers struct {
	minioHandler minioHandler.Handler // Пока у нас только один роут
}

// NewHandler создает экземпляр Handler с предоставленными сервисами
func NewHandler(
	minioService minio.Client,
) (*Services, *Handlers) {
	return &Services{
			minioService: minioService,
		}, &Handlers{
			// инициируем Minio handler, который на вход получает minio service
			minioHandler: *minioHandler.NewMinioHandler(minioService),
		}
}

// RegisterRoutes - метод регистрации всех роутов в системе
func (h *Handlers) RegisterRoutes(router *gin.Engine) {

	// Здесь мы обозначили все эндпоинты системы с соответствующими хендлерами
	minioRoutes := router.Group("/files")
	{
		minioRoutes.POST("/one", h.minioHandler.CreateOne)
		minioRoutes.POST("/many", h.minioHandler.CreateMany)

		minioRoutes.GET("/one", h.minioHandler.GetOne)
		minioRoutes.GET("/many", h.minioHandler.GetMany)

		minioRoutes.DELETE("/one", h.minioHandler.DeleteOne)
		minioRoutes.DELETE("/many", h.minioHandler.DeleteMany)
	}

}
