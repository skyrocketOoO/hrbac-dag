package main

import (
	"fmt"
	"rbac/internal/delivery"
	"rbac/internal/infra/sql"
	"rbac/internal/usecase"
	"strconv"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

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

	db, err := sql.InitDb()
	if err != nil {
		panic(err)
	}

	sqlRepo, err := sql.NewOrmRepository(db)
	if err != nil {
		panic(err)
	}

	usecaseRepo := usecase.NewUsecaseRepository(sqlRepo)

	handlerRepo := delivery.NewHandlerRepository(usecaseRepo)

	userApp := app.Group("/user")
	userHandler := handlerRepo.UserHandler
	userApp.Get("/", userHandler.ListUsers)
	userApp.Get("/:name", userHandler.GetUser)
	userApp.Delete("/:name", userHandler.DeleteUser)
	userApp.Post("/add-role", userHandler.AddRole)
	userApp.Post("/remove-role", userHandler.RemoveRole)
	userApp.Post("/find-all-object-relations", userHandler.FindAllObjectRelations)
	userApp.Post("/add-relation", userHandler.AddRelation)
	userApp.Post("/remove-relation", userHandler.RemoveRelation)
	userApp.Post("/check", userHandler.Check)

	roleApp := app.Group("/role")
	roleHandler := handlerRepo.RoleHandler
	roleApp.Get("/", roleHandler.ListRoles)
	roleApp.Get("/:name", roleHandler.GetRole)
	roleApp.Delete("/:name", roleHandler.DeleteRole)
	roleApp.Post("/add-relation", roleHandler.AddRelation)
	roleApp.Post("/remove-relation", roleHandler.RemoveRelation)
	roleApp.Post("/add-parent", roleHandler.AddParent)
	roleApp.Post("/remove-parent", roleHandler.RemoveParent)
	// roleApp.Get("/list-child-roles", roleHandler.ListChildRoles)
	roleApp.Post("/find-all-object-relations", roleHandler.FindAllObjectRelations)
	roleApp.Post("/get-members", roleHandler.GetMembers)
	roleApp.Post("/check", roleHandler.Check)

	objectApp := app.Group("/object")
	objectApp.Post("/list-user-has-relation", handlerRepo.ObjectHandler.ListUserHasRelationOnObject)
	objectApp.Post("/list-role-has-relation", handlerRepo.ObjectHandler.ListRoleHasWhatRelationOnObject)
	objectApp.Post("/list-user-or-role-has-relation", handlerRepo.ObjectHandler.ListUserOrRoleHasRelationOnObject)
	objectApp.Post("/list-relations", handlerRepo.ObjectHandler.ListRelations)

	relationApp := app.Group("/relation")
	relationHandler := handlerRepo.RelationHandler
	relationApp.Get("/", relationHandler.GetAllRelations)
	relationApp.Post("/link", relationHandler.Link)
	relationApp.Post("/check", relationHandler.Check)
	relationApp.Post("/path", relationHandler.Path) // to check how the subject obtain the relation on subject
	relationApp.Delete("/", relationHandler.ClearAllRelations)

	app.Listen(":3000")
}
