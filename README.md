## Installation and Usage

- To add all modules that use in this project and also remove unused modules, you can run "`go mod tidy`" command

- To run this project you can use this "`go run cmd/main.go`" command

- To try run some unit test you can move to helper directory and run the unit test with "`cd helper && go test -v`" command

## Available Endpoint

- **GET _/api/customers_**
  : This endpoint use to fetch all customer datas

- **GET _/api/items_**
  : This endpoint use to fetch all item datas

- **GET _/api/invoices_**
  : This endpoint use to fetch all invoice datas

- **GET _/api/invoice/:id_**
  : This endpoint use to find one invoice by its id

- **POST _/api/invoice_**
  : This endpoint use to create new invoice

- **POST _/api/invoice/:id_**
  : This endpoint use to update invoice by its id
