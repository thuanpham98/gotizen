package infra_database

import (
	domain_customer "gotizen/domain/customer"
)


func GetCustomer(id int) (*domain_customer.Customer, error){
	return &domain_customer.Customer{ID: id,Name: 
		"2343",Email: "5322432",Address: "656",Phone: "5354",},nil
}
