package router

import (
	"tpserver/internal/boilerplate"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	ctrl boilerplate.ControllerImplementation
	app  *fiber.App
}

func New(ctrl boilerplate.ControllerImplementation) *Router {
	return &Router{
		ctrl: ctrl,
		app:  fiber.New(fiber.Config{}),
	}
}

func (router *Router) RegisterRoutes() {
	// TODO auth := router.app.Group("/auth", handleToken) // ? авторизация
	// basic := router.app.Group("/core") // ? медиафайлы
	// basic.Get("/thumb", GetThumb)
	// basic.Get("/thumbs", GetThumbs)
	// basic.Get("/video", handleToken, GetVideoStream)

	// utility := router.app.Group("/util") // ? Утилитарка(запросы на получение N из БД, прооведение транзакций тдтп)
	// utility.Get("/objects", GetObjects)
	// utility.Get("/reviews", GetReviews)
	// utility.Post("/reviews", PostReview)
	// utility.Post("/buy", handleToken, BuyFilm)
	// utility.Get("/profile", handleToken, GetProfile)

}
