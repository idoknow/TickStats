package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/soulter/tickstats/config"
	"github.com/soulter/tickstats/controllers"
	"github.com/soulter/tickstats/models"
	"github.com/soulter/tickstats/repositories"
	"github.com/soulter/tickstats/services"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	var err error
	router := gin.Default()

	// Load .env file
	_ = godotenv.Load()

	// CORS
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
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

	// redirect to /
	router.GET("/dashboard", func(c *gin.Context) {
		// rewrite index.html to / so that the frontend router can handle the route
		c.File("./frontend/tick-stats-fe/dist/index.html")
	})
	// redirect /app and its subroutes to /
	router.GET("/app/*any", func(c *gin.Context) {
		c.File("./frontend/tick-stats-fe/dist/index.html")
	})
	router.GET("/help", func(c *gin.Context) {
		c.File("./frontend/tick-stats-fe/dist/index.html")
	})
	router.GET("/world", func(c *gin.Context) {
		c.File("./frontend/tick-stats-fe/dist/index.html")
	})

	// Initialize configuration
	config.LoadConfig()

	// from .env
	// viper.SetConfigFile("config/.env")
	viper.AutomaticEnv()

	var db *gorm.DB
	var metricsDB *gorm.DB

	// MySQL database
	dsn := viper.GetString("MYSQL_DSN")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect MySQL database")
	}

	// Postgres database
	dsn = viper.GetString("PG_DSN")
	metricsDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect PostgreSQL database")
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
	metricsService := services.NewMetricsService(metricsRepo, chartRepo, applicationRepo)
	statsService := services.NewStatsService(applicationRepo, chartRepo, accountRepo)
	accountController := controllers.NewAccountController(accountService, statsService)
	metricsController := controllers.NewMetricsController(metricsService)
	statsController := controllers.NewStatsController(statsService)

	// Register routes
	RegisterAccountRoutes(router, accountController)
	RegisterMetricRoutes(router, metricsController)
	RegisterStatsRoutes(router, statsController)

	return router
}
