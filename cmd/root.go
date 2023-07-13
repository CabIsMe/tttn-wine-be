/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/CabIsMe/tttn-wine-be/delivery"
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/repositories"
	"github.com/CabIsMe/tttn-wine-be/internal/services"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		schedule := gocron.NewScheduler(time.Local)
		fmt.Printf("Service starting at %s ...\n", carbon.Now().String())
		repos := repositories.NewRepos()
		services := services.NewServices(repos)
		handlers := delivery.NewHandlers(services)
		//
		AppServer := fiber.New()
		AppServer.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))

		// TODO: Upload file
		//Todo: APi portal call upload chưa gắn token

		AppServer.Use(logger.New(logger.Config{
			Format:     "${method} - ${path} - header:${reqHeaders} - body:${body} - resp-status:${status} - resp_body:${resBody}",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Ho_Chi_Minh",
		}))
		// Scheduler
		// Scan ready tasks 1 minute and trigger when received order
		schedule.StartAsync()

		// Inside
		AppServer.Post("/inside/employee-login", handlers.UserLoginHandler)
		AppServer.Post("/inside/create-promotion", handlers.VerifyTokenInside, handlers.CreatePromotionHandler)
		AppServer.Post("/inside/create-promotion-detail", handlers.VerifyTokenInside, handlers.CreatePromotionDetailHandler)

		// Client
		AppServer.Get("/client/list-products", handlers.AllProductsHandler)
		AppServer.Get("/client/list-promotional-products", handlers.PromotionalProductsHandler)
		AppServer.Get("/client/list-new-products", handlers.NewReleaseProductsHandler)
		AppServer.Post("/client/get-product", handlers.GetProductHandler)
		AppServer.Post("/client/sign-up", handlers.SignUpUserHandler)
		AppServer.Post("/client/customer-login", handlers.UserLoginHandler)

		if err := AppServer.Listen(":" + internal.Envs.ServicePort); err != nil {
			fmt.Println("Fiber server got error ", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}
