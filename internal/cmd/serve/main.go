package serve

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"student-app/internal/config"
	"student-app/internal/db"
	"student-app/internal/handler"
	"student-app/internal/store"
)

func main(logger *zap.Logger, cfg config.Config) {
	app := echo.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		logger.Named("db").Fatal("cannot create a db instance", zap.Error(err))
	}

	//var studentStore store.Students = store.NewStudentInMemory()
	var studentStore store.Students = store.NewStudentMongoDB(db, logger.Named("student"))

	ha := handler.Auth{
		Key:      []byte(cfg.Secret),
		Username: cfg.Admin.Username,
		Name:     cfg.Admin.Name,
		Password: cfg.Admin.Password,
		Logger:   logger.Named("auth"),
	}
	ha.Register(app.Group("/auth"))

	hs := handler.Student{
		Store:  studentStore,
		Logger: logger.Named("http").Named("student"),
	}
	app.GET("/health-check", handler.HealthCheck)
	hs.Register(app.Group("/api/students", ha.Auth))

	if err := app.Start(":8080"); err != nil {
		logger.Error("cannot start the http server", zap.Error(err))
	}
}

func New(logger *zap.Logger, cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "runs http server for students api",
		Run: func(cmd *cobra.Command, args []string) {
			main(logger, cfg)
		},
	}
}
