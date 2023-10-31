package http

func (s Server) Route() {
	v1 := s.app.Group("v1")
	v1.Post("/", s.CreateOrder)
	v1.Get("/", s.ListOrderRedis)
}
