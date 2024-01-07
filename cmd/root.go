package cmd

import (
	"fmt"
	"rbac/config"
	"rbac/internal/delivery"
	"rbac/internal/infra"
	"rbac/internal/usecase"
	"strconv"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "The command-line tool for hrbac system",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.ReadConfig(); err != nil {
				panic(err.Error())
			}

			app := fiber.New()

			app.Get("/swagger/*", swagger.HandlerDefault)

			app.Use(func(c *fiber.Ctx) error {
				// Continue processing the request
				err := c.Next()

				// Log the result and final status code
				blue := color.New(color.FgBlue).SprintFunc()
				yellow := color.New(color.FgYellow).SprintFunc()
				green := color.New(color.FgGreen).SprintFunc()
				status_code := blue(strconv.Itoa(c.Response().StatusCode()))
				method := yellow(c.Method())
				path := green(c.OriginalURL())
				fmt.Printf("%s %s %s %s\n", method, path, c.Response().Body(), status_code)
				if err != nil {
					// If there was an error processing the request, log the error
					fmt.Printf("Error processing request: %v\n", err)
					fmt.Println()
					return err
				}
				// Return nil to indicate that the middleware has completed processing
				return nil
			})

			app.Get("/healthy", func(c *fiber.Ctx) error {
				return nil
			})

			infraRepo := infra.NewInfraRepository()

			usecaseRepo := usecase.NewUsecaseRepository(infraRepo)

			handlerRepo := delivery.NewHandlerRepository(usecaseRepo)

			userApp := app.Group("/user")
			userHandler := handlerRepo.UserHandler
			userApp.Get("/", userHandler.GetAll)
			userApp.Delete("/:name", userHandler.Delete)
			userApp.Post("/add-role", userHandler.AddRole)
			userApp.Post("/remove-role", userHandler.RemoveRole)
			userApp.Post("/get-all-object-relations", userHandler.GetAllObjectRelations)
			userApp.Post("/add-relation", userHandler.AddRelation)
			userApp.Post("/remove-relation", userHandler.RemoveRelation)
			userApp.Post("/check", userHandler.Check)

			roleApp := app.Group("/role")
			roleHandler := handlerRepo.RoleHandler
			roleApp.Get("/", roleHandler.GetAll)
			roleApp.Delete("/:name", roleHandler.Delete)
			roleApp.Post("/add-relation", roleHandler.AddRelation)
			roleApp.Post("/remove-relation", roleHandler.RemoveRelation)
			roleApp.Post("/add-parent", roleHandler.AddParent)
			roleApp.Post("/remove-parent", roleHandler.RemoveParent)
			roleApp.Get("/get-child-roles", roleHandler.GetChildRoles)
			roleApp.Post("/get-all-object-relations", roleHandler.GetAllObjectRelations)
			roleApp.Post("/get-members", roleHandler.GetMembers)
			roleApp.Post("/check", roleHandler.Check)

			objectApp := app.Group("/object")
			objectApp.Post("/get-user-relations", handlerRepo.ObjectHandler.GetUserRelations)
			objectApp.Post("/get-role-relations", handlerRepo.ObjectHandler.GetRoleRelations)

			relationApp := app.Group("/relation")
			relationHandler := handlerRepo.RelationHandler
			relationApp.Get("/", relationHandler.GetAll)
			relationApp.Post("/query", relationHandler.Query)
			relationApp.Post("/", relationHandler.Create)
			relationApp.Delete("/", relationHandler.Delete)
			relationApp.Post("/check", relationHandler.Check)
			relationApp.Post("/clear-all-relations", relationHandler.ClearAllRelations)

			app.Listen(":3000")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
