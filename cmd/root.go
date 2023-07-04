/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fc_optimal_assignment/delivery"
	"fc_optimal_assignment/internal"
	"fc_optimal_assignment/internal/repositories"
	"fc_optimal_assignment/internal/services"
	"fmt"
	"os"
	"time"

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
		services := services.NewFcServices(repos)
		handlers := delivery.NewHandlers(repos, services)
		//
		AppServer := fiber.New()
		AppServer.Use(cors.New(cors.Config{
			AllowOrigins: "*",
		}))

		pathPrefix := "/hi-ecom-fconnect-v2-api"
		// TODO: Upload file
		AppServer.Post(pathPrefix+"/v1/public/web/upload-file", handlers.UploadFile)
		//Todo: APi portal call upload chưa gắn token
		AppServer.Post(pathPrefix+"/v1/public/portal/upload-file", handlers.UploadFile)

		AppServer.Use(logger.New(logger.Config{
			Format:     "${method} - ${path} - header:${reqHeaders} - body:${body} - resp-status:${status} - resp_body:${resBody}",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Ho_Chi_Minh",
		}))
		// Scheduler
		// Scan ready tasks 1 minute and trigger when received order
		schedule.Every(10).Seconds().Do(services.ScanReadyTasksToAssignTenant)
		schedule.StartAsync()
		// API
		AppServer.Get(pathPrefix+"/health",
			handlers.Health)

		AppServer.Post(pathPrefix+"/scan-ready-tasks", handlers.ScanReadyTasks)

		// Staff by leader
		// AppServer.Post(pathPrefix+"/v1/public/staff-by-leader/web/get-all-leader", handlers.RequireTokenPortal, handlers.GetAllLeaderHandler)
		// AppServer.Post(pathPrefix+"/v1/public/staff-by-leader/web/get-available-users", handlers.RequireTokenPortal, handlers.GetAvailableUsersHandler)
		// AppServer.Post(pathPrefix+"/v1/public/staff-by-leader/web/add-staff", handlers.RequireTokenPortal, handlers.AddNewStaffHandler)
		// AppServer.Post(pathPrefix+"/v1/public/staff-by-leader/web/delete-staff", handlers.RequireTokenPortal, handlers.DeleteStaffHandler)

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
