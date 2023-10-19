package server

func (s *Server) Routes() {
	group := s.e.Group("/api")
	group.GET("/customers", s.handler.GetCustomers)
	group.GET("/items", s.handler.GetItems)
	group.GET("/invoices", s.handler.GetInvoices)
	group.GET("/invoice/:id", s.handler.FindOneInvoice)
	group.PUT("/invoice/:id", s.handler.UpdateInvoice)
	group.POST("/invoice", s.handler.CreateInvoice)
}
