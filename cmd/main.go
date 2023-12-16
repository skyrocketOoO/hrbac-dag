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

	roleApp := app.Group("/role")
	roleApp.Post("/add-permission", handlerRepo.RoleHandler.AddPermissionToRole)
	roleApp.Post("/add-parent-role", handlerRepo.RoleHandler.AssignRoleUpRole)
	roleApp.Get("/list-child-roles", handlerRepo.RoleHandler.ListChildRoles)
	roleApp.Get("/list-role-permissions", handlerRepo.RoleHandler.ListRolePermissions)
	roleApp.Get("/list-roles", handlerRepo.RoleHandler.ListRoles)
	roleApp.Get("/get-role-members", handlerRepo.RoleHandler.GetRoleMembers)
	roleApp.Delete("/delete-role", handlerRepo.RoleHandler.DeleteRole)

	userApp := app.Group("/user")
	userApp.Post("/add-role", handlerRepo.UserHandler.AddUserToRole)
	userApp.Post("/remove-from-role", handlerRepo.UserHandler.RemoveUserFromRole)
	userApp.Get("/list-permissions", handlerRepo.UserHandler.ListUserPermissions)
	userApp.Post("/add-permission", handlerRepo.UserHandler.AddPermissionToUser)
	userApp.Get("/list-users", handlerRepo.UserHandler.ListUsers)

	objectApp := app.Group("/object")
	objectApp.Get("/list-who-has-permission", handlerRepo.ObjectHandler.ListWhoHasPermissionOnObject)
	objectApp.Get("/list-roles-has-permission", handlerRepo.ObjectHandler.ListRolesHasWhatPermissionOnObject)
	objectApp.Get("/list-who-or-role-has-permission", handlerRepo.ObjectHandler.ListWhoOrRoleHasPermissionOnObject)
	objectApp.Get("/list-all-permissions", handlerRepo.ObjectHandler.ListAllPermissions)

	permissionApp := app.Group("/permission")
	permissionApp.Post("/link", handlerRepo.ObjectHandler.LinkPermission)
	permissionApp.Post("/check-user-permission", handlerRepo.PermissionHandler.CheckUserPermission)

	app.Listen(":3000")
}
