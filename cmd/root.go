/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CabIsMe/tttn-wine-be/delivery"
	"github.com/CabIsMe/tttn-wine-be/internal"
	"github.com/CabIsMe/tttn-wine-be/internal/models"
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
		AppServer.Post("/inside/update-customer-order", handlers.VerifyTokenInside, handlers.UpdateCustomerOrderHandler)
		AppServer.Post("/inside/update-status-customer-order", handlers.VerifyTokenInside, handlers.UpdateStatusCustomerOrderHandler)
		AppServer.Get("/inside/list-customer-orders", handlers.VerifyTokenInside, handlers.AllCustomerOrdersHandler)
		AppServer.Post("/inside/list-customer-orders-by-status", handlers.VerifyTokenInside, handlers.AllCustomerOrdersByStatusHandler)
		AppServer.Get("/inside/user-info", handlers.VerifyTokenInside, handlers.GetAccountInfoHandler)
		AppServer.Get("/inside/list-deliverers", handlers.VerifyTokenInside, handlers.AllDeliverersHandler)
		AppServer.Post("/inside/create-bill", handlers.VerifyTokenInside, handlers.CreateBillHandler)
		AppServer.Post("/inside/get-bill-by-order", handlers.VerifyTokenInside, handlers.GetBillByCustomerOrderIdHandler)
		AppServer.Post("inside/get-customer-order", handlers.GetCustomerOrderToCreateBillHandler)
		AppServer.Post("/inside/get-revenue", handlers.VerifyTokenInside, handlers.GetRevenueDateToDateHandler)
		AppServer.Get("/inside/list-status-customer-orders", handlers.VerifyTokenInside, func(ctx *fiber.Ctx) error {
			return ctx.Status(http.StatusOK).JSON(models.Cos)
		})
		AppServer.Post("/inside/add-product", handlers.VerifyTokenInside, handlers.AddNewProductHandler)
		AppServer.Post("/inside/update-product", handlers.VerifyTokenInside, handlers.UpdateProductHandler)

		// Client
		AppServer.Get("/client/list-products", handlers.AllProductsHandler)
		AppServer.Get("/client/list-top-products", handlers.TopSellingProductsHandler)
		AppServer.Get("/client/list-promotional-products", handlers.PromotionalProductsHandler)
		AppServer.Post("/client/list-relevant-products", handlers.GetProductsByTypeAndBrandHandler)
		AppServer.Get("/client/get-promotion-by-date", handlers.GetPromotionByDateHandler)
		AppServer.Get("/client/list-new-products", handlers.NewReleaseProductsHandler)
		AppServer.Post("/client/list-products-by-brand", handlers.GetProductsByBrandHandler)
		AppServer.Post("/client/list-products-by-category", handlers.GetProductsByCategoryHandler)
		AppServer.Post("/client/get-product", handlers.GetProductHandler)
		AppServer.Post("/client/get-product-by-name", handlers.GetProductByNameHandler)
		AppServer.Post("/client/sign-up", handlers.SignUpUserHandler)
		AppServer.Post("/client/customer-login", handlers.UserLoginHandler)
		AppServer.Post("/client/login-with-google", handlers.SignInWithGoogleHandler)
		AppServer.Post("/client/create-customer-order", handlers.VerifyTokenClient, handlers.CreateCustomerOrder)
		AppServer.Post("/client/add-cart", handlers.VerifyTokenClient, handlers.AddProductsToCartHandler)
		AppServer.Post("/client/remove-cart", handlers.VerifyTokenClient, handlers.RemoveProductsToCartHandler)
		AppServer.Get("/client/cart", handlers.VerifyTokenClient, handlers.AllProductsInCartHandler)
		AppServer.Get("/client/customer-info", handlers.VerifyTokenClient, handlers.GetCustomerInfoHandler)
		AppServer.Post("/client/update-customer", handlers.VerifyTokenClient, handlers.UpdateCustomerHandler)
		AppServer.Post("/client/payment/success", handlers.VerifyTokenClient, handlers.UpdatePaymentStatusHandler)
		AppServer.Post("/client/list-ordered", handlers.VerifyTokenClient, handlers.AllCustomerOrdersByCustomerHandler)
		if err := AppServer.Listen(":" + internal.Envs.ServicePort); err != nil {
			fmt.Println("Fiber server got error ", err)
		}
	},
}

func Execute() {
	os.Setenv("RAPYD_ACCESS_KEY", "rak_232FCC9417B6261038A9")
	os.Setenv("RAPYD_SECRET_KEY", "rsk_cb8692a765a4372bf4101a68fd57e45f52f3148a07d5807ca892476be0e289a889347d9fb9014386")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}
