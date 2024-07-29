package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/soulter/tickstats/config"
	"github.com/soulter/tickstats/controllers"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/services"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	var err error
	router := gin.Default()

	// CORS
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{
					"http://localhost:3000",
					"http://127.0.0.1:3000",
					"https://ts.lwl.lol",
				},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Cookie"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	router.StaticFile("/", "./frontend/tick-stats-fe/dist/index.html")
	router.Static("/assets", "./frontend/tick-stats-fe/dist/assets")

	// Initialize configuration
	config.LoadConfig()

	// from .env
	viper.SetConfigFile("config/.env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var db *gorm.DB
	var metricsDB *gorm.DB

	// SQLite database
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Postgres database
	dsn := "host=" + viper.GetString("PG_HOST") + " user=" + viper.GetString("PG_USER") + " password=" + viper.GetString("PG_PASSWORD") + " dbname=" + viper.GetString("PG_DBNAME") + " port=" + viper.GetString("PG_PORT") + " sslmode=disable TimeZone=" + viper.GetString("PG_TIMEZONE")
	metricsDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Chart{})
	metricsDB.AutoMigrate(&models.BasicMetricData{})

	// Create a hypertable for metrics
	metricsDB.Exec("SELECT create_hypertable('basic_metric_data', by_range('time'));")

	// Initialize components
	accountRepo := repositories.NewAccountRepository(db)
	applicationRepo := repositories.NewApplicationRepository(db)
	chartRepo := repositories.NewChartRepository(db)
	metricsRepo := repositories.NewMetricsRepository(metricsDB)
	accountService := services.NewAccountService(accountRepo, applicationRepo, chartRepo)
	metricsService := services.NewMetricsService(metricsRepo, chartRepo)
	accountController := controllers.NewAccountController(accountService)
	metricsController := controllers.NewMetricsController(metricsService)

	// Register routes
	RegisterAccountRoutes(router, accountController)
	RegisterMetricRoutes(router, metricsController)

	return router
}
