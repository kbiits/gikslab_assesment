package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/kbiits/gikslab_assesment/internal/app/middleware"
)

func (h *httpServer) prepareRoutes() {
	h.fiber.Use(recover.New())
	h.fiber.Use(middleware.RequestIdMiddleware())
	h.fiber.Use(middleware.BasicAuthMiddleware(h.config.Http))
	router := h.fiber.Use(timeout.NewWithContext(func(c *fiber.Ctx) error {
		return c.Next()
	}, time.Millisecond*time.Duration(h.config.Http.TimeoutPerRequestInMs)))

	router.Post("/v1/auth/login", h.handlers.AuthHandler.Login)

	jwtProtectedRouter := router.Use(middleware.ValidateToken(h.jwt))
	jwtProtectedRouter.Get("/v1/auth/logout", h.handlers.AuthHandler.Logout)
	jwtProtectedRouter.Post("/v1/user", middleware.BoardUsersOnly(), h.handlers.AuthHandler.Register)

	activityRoutes := jwtProtectedRouter.Group("/v1/activity")
	activityRoutes.Get("/:skillId", h.handlers.ActivityHandler.ListActivities)
	activityRoutes.Post("/", middleware.ExpertUsersOnly(), h.handlers.ActivityHandler.UpsertActivity)
	activityRoutes.Put("/:activityId", middleware.ExpertUsersOnly(), h.handlers.ActivityHandler.UpsertActivity)
	activityRoutes.Delete("/:activityId", middleware.ExpertUsersOnly(), h.handlers.ActivityHandler.DeleteActivity)

	skillRoutes := jwtProtectedRouter.Group("/v1/skills")
	skillRoutes.Get("/", h.handlers.SkillHandler.ListSkill)
}
