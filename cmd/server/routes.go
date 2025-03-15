package server

func SetupPublicRoutes(h *HttpServer) {
	ayman := h.PublicEngine.Group(BasePath)

	userGroup := ayman.Group("/users")

	setupUserRoutes(userGroup, h)
}