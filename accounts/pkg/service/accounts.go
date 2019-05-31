package service

import "context"

// Customer data type
type Customer struct {
	CustomerID string `json:"customerId,omitempty"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	CreatedAt  int64  `json:"createdAt"`
}

// Admin data type
type Admin struct {
	AdminID   string `json:"adminId,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	CreatedAt int64  `json:"createdAt"`
}

// Repository for both admin and customer models
type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	CreateAdmin(ctx context.Context, admin Admin) error
	GetCustomerByID(ctx context.Context, id string) (Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (Customer, error)
	GetAdminByID(ctx context.Context, id string) (Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (Admin, error)
	UpdateCustomer(ctx context.Context, customer Customer) error
	UpdateAdmin(ctx context.Context, admin Admin) error
}
