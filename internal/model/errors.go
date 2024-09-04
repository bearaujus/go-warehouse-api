package model

import (
	"errors"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/errwrap"
)

// ---------------------------------------------------------------------------------------------------------------------
// Handler Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var ()

// Product Service

var (
	ErrHProductHTTPGetProductsWithStock = errwrap.NewErrDef("h.product.http.001", "failed to get products with stock in handler. %v")
	ErrHProductHTTPCreateProduct        = errwrap.NewErrDef("h.product.http.002", "failed to create product in handler. %v")
)

// Shop Service

var (
	ErrHShopHTTPGetShops   = errwrap.NewErrDef("h.shop.http.001", "failed to get shops in handler. %v")
	ErrHShopHTTPCreateShop = errwrap.NewErrDef("h.shop.http.002", "failed to create shop in handler. %v")
)

// User Service

var (
	ErrHUserHTTPGetUserById = errwrap.NewErrDef("h.user.http.001", "failed to get user by id in handler. %v")
)

// Warehouse Service

var (
	ErrHWarehouseHTTPGetWarehouses                     = errwrap.NewErrDef("h.warehouse.http.001", "failed to get warehouses in handler. %v")
	ErrHWarehouseHTTPCreateWarehouse                   = errwrap.NewErrDef("h.warehouse.http.002", "failed to create warehouse in handler. %v")
	ErrHWarehouseHTTPCreateWarehouseInboundTransaction = errwrap.NewErrDef("h.warehouse.http.003", "failed to create warehouse inbound transaction in handler. %v")
)

// ---------------------------------------------------------------------------------------------------------------------
// Usecases Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var ()

// Product Service

var (
	ErrUProductCreateProduct = errwrap.NewErrDef("u.product.001", "failed to create product in usecase. %v")
)

// Shop Service

var (
	ErrUShopCreateShop = errwrap.NewErrDef("u.shop.001", "failed to create shop in usecase. %v")
)

// User Service

var (
	ErrUUserRegister = errwrap.NewErrDef("u.user.001", "failed to register user in usecase. %v")
	ErrUUserLogin    = errwrap.NewErrDef("u.user.002", "failed to login in usecase. %v")
)

// Warehouse Service

var (
	ErrUWarehouseCreateWarehouse = errwrap.NewErrDef("u.warehouse.001", "failed to create warehouse in usecase. %v")
)

// ---------------------------------------------------------------------------------------------------------------------
// Resources Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var (
	// Postgres Errors

	ErrROrderPostgresGetOrdersByUser = errwrap.NewErrDef("r.order.postgres.001", "failed to get orders by user from postgres. %v")
	ErrROrderPostgresCreateOrder     = errwrap.NewErrDef("r.order.postgres.002", "failed to create order in postgres. %v")

	// HTTP Client Errors

	ErrROrderHTTPClientGetOrdersByUser = errwrap.NewErrDef("r.order.http_client.001", "failed to get orders by user via http call. %v")
	ErrROrderHTTPClientCreateOrder     = errwrap.NewErrDef("r.order.http_client.002", "failed to create order via http call. %v")
)

// Product Service

var (
	// Postgres Errors

	ErrRProductPostgresGetProductsWithStockByUser = errwrap.NewErrDef("r.product.postgres.001", "failed to get products with stock by user from postgres. %v")
	ErrRProductPostgresCreateProduct              = errwrap.NewErrDef("r.product.postgres.002", "failed to create product in postgres. %v")

	// HTTP Client Errors

	ErrRProductHTTPClientGetProductsWithStockByUser = errwrap.NewErrDef("r.product.http_client.001", "failed to get products with stock by user via http call. %v")
	ErrRProductHTTPClientCreateProduct              = errwrap.NewErrDef("r.product.http_client.002", "failed to create product via http call. %v")
)

// Shop Service

var (
	// Postgres Errors

	ErrRShopPostgresGetShopsByUser = errwrap.NewErrDef("r.shop.postgres.001", "failed to get shops by user from postgres. %v")
	ErrRShopPostgresCreateShop     = errwrap.NewErrDef("r.shop.postgres.002", "failed to create shop in postgres. %v")

	// HTTP Client Errors

	ErrRShopHTTPClientGetShopsByUser = errwrap.NewErrDef("r.shop.http_client.001", "failed to get shops by user via http call. %v")
	ErrRShopHTTPClientCreateShop     = errwrap.NewErrDef("r.shop.http_client.002", "failed to create shop via http call. %v")
)

// User Service

var (
	// Postgres Errors

	ErrRUserPostgresGetUserById    = errwrap.NewErrDef("r.user.postgres.001", "failed to get user by id from postgres. %v")
	ErrRUserPostgresGetUserByLogin = errwrap.NewErrDef("r.user.postgres.002", "failed to get user by login from postgres. %v")
	ErrRUserPostgresCreateUser     = errwrap.NewErrDef("r.user.postgres.003", "failed to create user in postgres. %v")

	// HTTP Client Errors

	ErrRUserHTTPClientGetUserById    = errwrap.NewErrDef("r.user.http_client.001", "failed to get user by id via http call. %v")
	ErrRUserHTTPClientGetUserByLogin = errwrap.NewErrDef("r.user.http_client.002", "failed to get user by login via http call. %v")
	ErrRUserHTTPClientCreateUser     = errwrap.NewErrDef("r.user.http_client.003", "failed to create user via http call. %v")
)

// Warehouse Service

var (
	// Postgres Errors

	ErrRWarehousePostgresGetWarehousesByUser               = errwrap.NewErrDef("r.warehouse.postgres.001", "failed to get warehouses by user from postgres. %v")
	ErrRWarehousePostgresGetWarehousesByUserAndShop        = errwrap.NewErrDef("r.warehouse.postgres.002", "failed to get warehouses by user and shop from postgres. %v")
	ErrRWarehousePostgresCreateWarehouse                   = errwrap.NewErrDef("r.warehouse.postgres.003", "failed to create warehouse in postgres. %v")
	ErrRWarehousePostgresCreateWarehouseInboundTransaction = errwrap.NewErrDef("r.warehouse.postgres.004", "failed to create warehouse inbound transaction from postgres. %v")

	// HTTP Client Errors

	ErrRWarehouseHTTPClientGetWarehousesByUser               = errwrap.NewErrDef("r.warehouse.http_client.001", "failed to get warehouses by user via http call. %v")
	ErrRWarehouseHTTPClientGetWarehousesByUserAndShop        = errwrap.NewErrDef("r.warehouse.http_client.002", "failed to get warehouses by user and shop via http call. %v")
	ErrRWarehouseHTTPClientCreateWarehouse                   = errwrap.NewErrDef("r.warehouse.http_client.003", "failed to create warehouse via http call. %v")
	ErrRWarehouseHTTPClientCreateWarehouseInboundTransaction = errwrap.NewErrDef("r.warehouse.http_client.004", "failed to create warehouse inbound transaction via http call. %v")
)

// ---------------------------------------------------------------------------------------------------------------------
// Middlewares Errors
// ---------------------------------------------------------------------------------------------------------------------

var (
	ErrMWAuthAuthenticateUser    = errwrap.NewErrDef("m.auth.001", "user is not authenticated")
	ErrMWAuthAuthenticateService = errwrap.NewErrDef("m.auth.002", "service is not authenticated")
	ErrMWAuthAuthorizeSeller     = errwrap.NewErrDef("m.auth.003", "not authorized")
	ErrMWAuthAuthorizeBuyer      = errwrap.NewErrDef("m.auth.004", "not authorized")
)

// ---------------------------------------------------------------------------------------------------------------------
// Common Errors
// ---------------------------------------------------------------------------------------------------------------------

var (
	ErrCommonNotImplemented     = errors.New("not implemented")
	ErrCommonInvalidRequestBody = errors.New("invalid request body")
	ErrCommonInvalidRequestURL  = errors.New("invalid request url")
	ErrCommonInvalidAuthToken   = errors.New("invalid auth token")
)
