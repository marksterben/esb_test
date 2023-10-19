package server

func (s *Server) Routes() {
	group := s.e.Group("/api")
	group.GET("/customers", s.handler.GetCustomers)
	group.GET("/items", s.handler.GetItems)
	group.GET("/invoices", s.handler.GetInvoices)
	group.POST("/invoice", s.handler.CreateInvoice)
}
