package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"concerts/internal/config"
	"concerts/internal/handler"
	"concerts/internal/middleware"
	"concerts/internal/repository"
	"concerts/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	// เชื่อมต่อ Database
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Repository
	userRepo := repository.NewUserRepository(db)
	concertRepo := repository.NewConcertRepository(db)
	bannerRepo := repository.NewBannerRepository(db)
	contentTemplateRepo := repository.NewContentTemplateRepository(db)
	commissionRepo := repository.NewCommissionRepository(db)
	reportRepo := repository.NewReportRepository(db)
	partnerRepo := repository.NewPartnerRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	withdrawRepo := repository.NewWithdrawRepository(db)
	authRepo := repository.NewAuthRepository(db)

	// Service
	userService := service.NewUserService(userRepo)
	concertService := service.NewConcertService(concertRepo)
	bannerService := service.NewBannerService(bannerRepo)
	contentTemplateService := service.NewContentTemplateService(contentTemplateRepo)
	commissionService := service.NewCommissionService(commissionRepo)
	reportService := service.NewReportService(reportRepo)
	partnerService := service.NewPartnerService(partnerRepo, bookingRepo, withdrawRepo)
	authService := service.NewAuthService(authRepo, cfg.JWTSecret)

	// Handler
	userHandler := handler.NewUserHandler(userService)
	concertHandler := handler.NewConcertHandler(concertService)
	bannerHandler := handler.NewBannerHandler(bannerService)
	contentTemplateHandler := handler.NewContentTemplateHandler(contentTemplateService)
	commissionHandler := handler.NewCommissionHandler(commissionService)
	reportHandler := handler.NewReportHandler(reportService)
	partnerHandler := handler.NewPartnerHandler(partnerService, withdrawRepo)
	authHandler := handler.NewAuthHandler(authService)

	// สร้าง Router
	r := gin.Default()

	// Health Check ไม่ต้องใช้ Token
	r.GET("/health", func(c *gin.Context) {
		if err := repository.CheckDBConnection(db); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "Database connection failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "database": "connected"})
	})

	// Auth API (No Auth)
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/internal/register", authHandler.Register)
		auth.POST("/internal/login", authHandler.Login)
		auth.GET("/internal/validate-token", middleware.JWTAuth(authService), authHandler.ValidateToken)
		auth.POST("/internal/logout", middleware.JWTAuth(authService), authHandler.Logout)
	}

	// Protected routes ต้องใช้ Bearer Token
	protected := r.Group("/api/v1", middleware.BearerAuth(cfg.APIToken))
	{
		protected.GET("/internal/users", userHandler.GetAllUsers)       //
		protected.GET("/internal/users/:id", userHandler.GetUserByID)   //
		protected.POST("/internal/users", userHandler.CreateUser)       //
		protected.PUT("/internal/users/:id", userHandler.UpdateUser)    //
		protected.DELETE("/internal/users/:id", userHandler.DeleteUser) //

		protected.GET("/external/concerts", concertHandler.GetAllConcerts) //
		protected.GET("/external/search", concertHandler.SearchConcerts)   //

		protected.GET("/external/banners", bannerHandler.GetBanners) //

		protected.GET("/external/content-templates", contentTemplateHandler.GetContentTemplates) //

		protected.GET("/external/commissions/:partner_id", commissionHandler.GetCommissions) //

		protected.GET("/external/partner/rewards/:partner_id", partnerHandler.GetPartnerRewards) //
		protected.GET("/external/partner/bookings/:partner_id", partnerHandler.GetBookings)

		protected.GET("/external/reports/sales", reportHandler.GetSalesReport)
		protected.GET("/external/reports/sales-by-source", reportHandler.GetSalesBySource) //

		protected.GET("/external/partner/balance/:partner_id", partnerHandler.GetPartnerBalance)
		protected.POST("/external/partner/auto-withdraw/:partner_id", partnerHandler.SetAutoWithdraw)
		protected.POST("/external/partner/request-withdrawal/:partner_id", partnerHandler.CreateWithdrawRequest)

	}

	r.Run(":80")
}
