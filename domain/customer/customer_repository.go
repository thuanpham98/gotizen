package domain_customer

type CustomerRepository interface {
	GetCustomer(id int) (*Customer, error)
}