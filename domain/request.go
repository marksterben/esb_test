package domain

type InvoiceRequest struct {
	InvoiceID string `query:"invoice_id"`
	IssueDate string `query:"issue_date"`
	Subject   string `query:"subject"`
	TotalItem *int   `query:"total_item"`
	Customer  string `query:"customer"`
	DueDate   string `query:"due_date"`
	Status    string `query:"status"`
	Limit     int    `query:"limit"`
	Page      int    `query:"page"`
}

type PostInvoiceRequest struct {
	ID         string
	Subject    string `form:"subject"`
	IssueDate  string `form:"issueDate"`
	DueDate    string `form:"dueDate"`
	TotalItem  int    `form:"totalItem"`
	Subtotal   int    `form:"subtotal"`
	Tax        int    `form:"tax"`
	Grandtotal int    `form:"grandtotal"`
	CustomerID int    `form:"customerId"`
	Details    []struct {
		ItemID int `form:"itemId"`
		Qty    int `form:"qty"`
		Price  int `form:"price"`
		Amount int `form:"amount"`
	} `form:"details"`
}
