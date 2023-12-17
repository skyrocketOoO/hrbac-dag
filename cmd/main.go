package main

import (
	"fmt"
	"rbac/internal/delivery"
	"rbac/internal/infra/sql"
	"rbac/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/healthy", func(c *fiber.Ctx) error {
		fmt.Printf("Received request: %s %s\n", c.Method(), c.OriginalURL())
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
	userApp.Post("/list-relation", userHandler.ListRelations)
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
	roleApp.Get("/list-relation", roleHandler.ListRelations)
	roleApp.Get("/get-members", roleHandler.GetMembers)
	roleApp.Post("/check", roleHandler.Check)

	objectApp := app.Group("/object")
	objectApp.Get("/list-user-has-relation", handlerRepo.ObjectHandler.ListUserHasRelationOnObject)
	objectApp.Get("/list-role-has-relation", handlerRepo.ObjectHandler.ListRoleHasWhatRelationOnObject)
	objectApp.Get("/list-user-or-role-has-relation", handlerRepo.ObjectHandler.ListUserOrRoleHasRelationOnObject)
	objectApp.Get("/list-relations", handlerRepo.ObjectHandler.ListRelations)

	relationApp := app.Group("/relation")
	relationHandler := handlerRepo.RelationHandler
	relationApp.Get("/", relationHandler.ListRelations)
	relationApp.Post("/link", relationHandler.Link)
	relationApp.Post("/check", relationHandler.Check)
	relationApp.Post("/path", relationHandler.Path) // to check how the subject obtain the relation on subject

	app.Listen(":3000")
}
