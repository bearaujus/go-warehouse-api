package model

import (
	"errors"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/errwrap"
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
// Handler Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var (
	ErrHOrderHTTPGetOrdersByUserIdAndStatus = errwrap.NewErrDef("h.order.http.001", "failed to get orders by user id and status in handler. %v")
	ErrHOrderHTTPCreateOrder                = errwrap.NewErrDef("h.order.http.002", "failed to create order in handler. %v")
	ErrHOrderHTTPCompleteOrder              = errwrap.NewErrDef("h.order.http.003", "failed to complete order in handler. %v")
)

// Product Service

var (
	ErrHProductHTTPGetSellerProductsByShopUserId = errwrap.NewErrDef("h.product.http.001", "failed to get seller products by shop user id in handler. %v")
	ErrHProductHTTPCreateProduct                 = errwrap.NewErrDef("h.product.http.003", "failed to create product in handler. %v")
)

// Shop Service

var (
	ErrHShopHTTPCreateShop = errwrap.NewErrDef("h.shop.http.001", "failed to create shop in handler. %v")
)

// User Service

var (
	ErrHUserHTTPGetUserById  = errwrap.NewErrDef("h.user.http.001", "failed to get user by id in handler. %v")
	ErrHUserHTTPRegisterUser = errwrap.NewErrDef("h.user.http.002", "failed to register user in handler. %v")
	ErrHUserHTTPLoginUser    = errwrap.NewErrDef("h.user.http.003", "failed to login user in handler. %v")
)

// Warehouse Service

var (
	ErrHWarehouseHTTPGetWarehousesByShopUserId     = errwrap.NewErrDef("h.warehouse.http.001", "failed to get warehouses by shop user id in handler. %v")
	ErrHWarehouseHTTPCreateWarehouse               = errwrap.NewErrDef("h.warehouse.http.002", "failed to create warehouse in handler. %v")
	ErrHWarehouseHTTPUpdateWarehouse               = errwrap.NewErrDef("h.warehouse.http.003", "failed to update warehouse in handler. %v")
	ErrHWarehouseHTTPAddWarehouseProductStock      = errwrap.NewErrDef("h.warehouse.http.004", "failed to add warehouse product stock in handler. %v")
	ErrHWarehouseHTTPTransferWarehouseProductStock = errwrap.NewErrDef("h.warehouse.http.005", "failed to transfer warehouse product stock in handler. %v")
)

// ---------------------------------------------------------------------------------------------------------------------
// Usecases Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var (
	ErrUOrderGetOrdersByUserIdAndStatus = errwrap.NewErrDef("h.order.001", "failed to get orders by user id and status in usecase. %v")
	ErrUOrderCreateOrder                = errwrap.NewErrDef("u.order.002", "failed to create order in usecase. %v")
)

// Product Service

var (
	ErrUProductGetProductsByShopUserIdAndWarehouseStatus = errwrap.NewErrDef("u.product.001", "failed to get products by shop user id and warehouse status in usecase. %v")
	ErrUProductCreateProduct                             = errwrap.NewErrDef("u.product.002", "failed to create product in usecase. %v")
)

// Shop Service

var (
	ErrUShopCreateShop = errwrap.NewErrDef("u.shop.001", "failed to create shop in usecase. %v")
)

// User Service

var (
	ErrUUserRegister = errwrap.NewErrDef("u.user.001", "failed to register user in usecase. %v")
	ErrUUserLogin    = errwrap.NewErrDef("u.user.002", "failed to login user in usecase. %v")
)

// Warehouse Service

var (
	ErrUWarehouseCreateWarehouse               = errwrap.NewErrDef("u.warehouse.001", "failed to create warehouse in usecase. %v")
	ErrUWarehouseUpdateWarehouse               = errwrap.NewErrDef("u.warehouse.002", "failed to update warehouse in usecase. %v")
	ErrUWarehouseAddWarehouseProductStock      = errwrap.NewErrDef("u.warehouse.003", "failed to add warehouse product stock in usecase. %v")
	ErrUWarehouseTransferWarehouseProductStock = errwrap.NewErrDef("u.warehouse.004", "failed to transfer warehouse product stock in usecase. %v")
)

// ---------------------------------------------------------------------------------------------------------------------
// Resources Errors
// ---------------------------------------------------------------------------------------------------------------------

// Order Service

var (
	// Postgres Errors

	ErrROrderPostgresGetOrdersByUserIdAndStatus = errwrap.NewErrDef("r.order.postgres.001", "failed to get orders by user id and status from postgres. %v")
	ErrROrderPostgresCreateOrder                = errwrap.NewErrDef("r.order.postgres.002", "failed to create order in postgres. %v")
	ErrROrderPostgresCompleteOrder              = errwrap.NewErrDef("r.order.postgres.003", "failed to complete order in postgres. %v")
	ErrROrderPostgresProcessExpiredOrders       = errwrap.NewErrDef("r.order.postgres.004", "failed to process expired orders in postgres. %v")

	// HTTP Client Errors

	ErrROrderHTTPClientGetOrdersByUserIdAndStatus = errwrap.NewErrDef("r.order.http_client.001", "failed to get orders by user id and status via http call. %v")
	ErrROrderHTTPClientCreateOrder                = errwrap.NewErrDef("r.order.http_client.002", "failed to create order via http call. %v")
	ErrROrderHTTPClientCompleteOrder              = errwrap.NewErrDef("r.order.http_client.003", "failed to complete order via http call. %v")
	ErrROrderHTTPClientProcessExpiredOrders       = errwrap.NewErrDef("r.order.http_client.004", "failed to process expired orders via http call. %v")
)

// Product Service

var (
	// Postgres Errors

	ErrRProductPostgresGetProductsByShopUserIdAndWarehouseStatus = errwrap.NewErrDef("r.product.postgres.001", "failed to get products by shop user id and warehouse status from postgres. %v")
	ErrRProductPostgresCreateProduct                             = errwrap.NewErrDef("r.product.postgres.002", "failed to create product in postgres. %v")

	// HTTP Client Errors

	ErrRProductHTTPClientGetProductsByShopUserIdAndWarehouseStatus = errwrap.NewErrDef("r.product.http_client.001", "failed to get products by shop user id and warehouse status via http call. %v")
	ErrRProductHTTPClientCreateProduct                             = errwrap.NewErrDef("r.product.http_client.002", "failed to create product via http call. %v")
)

// Shop Service

var (
	// Postgres Errors

	ErrRShopPostgresCreateShop = errwrap.NewErrDef("r.shop.postgres.001", "failed to create shop in postgres. %v")

	// HTTP Client Errors

	ErrRShopHTTPClientCreateShop = errwrap.NewErrDef("r.shop.http_client.001", "failed to create shop via http call. %v")
)

// User Service

var (
	// Postgres Errors

	ErrRUserPostgresGetUserById    = errwrap.NewErrDef("r.user.postgres.001", "failed to get user by id from postgres. %v")
	ErrRUserPostgresGetUserByLogin = errwrap.NewErrDef("r.user.postgres.002", "failed to get user by login from postgres. %v")
	ErrRUserPostgresCreateUser     = errwrap.NewErrDef("r.user.postgres.003", "failed to create user in postgres. %v")
	ErrRUserPostgresDeleteUser     = errwrap.NewErrDef("r.user.postgres.004", "failed to delete user in postgres. %v")

	// HTTP Client Errors

	ErrRUserHTTPClientGetUserById    = errwrap.NewErrDef("r.user.http_client.001", "failed to get user by id via http call. %v")
	ErrRUserHTTPClientGetUserByLogin = errwrap.NewErrDef("r.user.http_client.002", "failed to get user by login via http call. %v")
	ErrRUserHTTPClientCreateUser     = errwrap.NewErrDef("r.user.http_client.003", "failed to create user via http call. %v")
	ErrRUserHTTPClientDeleteUser     = errwrap.NewErrDef("r.user.http_client.004", "failed to delete user via http call. %v")
)

// Warehouse Service

var (
	// Postgres Errors

	ErrRWarehousePostgresGetWarehousesByShopUserId                         = errwrap.NewErrDef("r.warehouse.postgres.001", "failed to get warehouses by shop user id from postgres. %v")
	ErrRWarehousePostgresGetActiveWarehouseProductStocksByProductId        = errwrap.NewErrDef("r.warehouse.postgres.002", "failed to get active warehouse product stocks by product id from postgres. %v")
	ErrRWarehousePostgresGetWarehouseProductStocksByShopUserIdAndProductId = errwrap.NewErrDef("r.warehouse.postgres.003", "failed to get warehouse product stocks by shop user id and product id from postgres. %v")
	ErrRWarehousePostgresCreateWarehouse                                   = errwrap.NewErrDef("r.warehouse.postgres.004", "failed to create warehouse in postgres. %v")
	ErrRWarehousePostgresUpdateWarehouse                                   = errwrap.NewErrDef("r.warehouse.postgres.005", "failed to update warehouse in postgres. %v")
	ErrRWarehousePostgresAddWarehouseProductStock                          = errwrap.NewErrDef("r.warehouse.postgres.006", "failed to add warehouse product stock from postgres. %v")
	ErrRWarehousePostgresTransferWarehouseProductStock                     = errwrap.NewErrDef("r.warehouse.postgres.007", "failed to transfer warehouse product stock from postgres. %v")

	// HTTP Client Errors

	ErrRWarehouseHTTPClientGetWarehousesByShopUserId                         = errwrap.NewErrDef("r.warehouse.http_client.001", "failed to get warehouses by shop user id via http call. %v")
	ErrRWarehouseHTTPClientGetActiveWarehouseProductStocksByProductId        = errwrap.NewErrDef("r.warehouse.http_client.002", "failed to get active warehouse product stocks by product id via http call. %v")
	ErrRWarehouseHTTPClientGetWarehouseProductStocksByShopUserIdAndProductId = errwrap.NewErrDef("r.warehouse.http_client.003", "failed to get warehouse product stocks by shop user id and product id via http call. %v")
	ErrRWarehouseHTTPClientCreateWarehouse                                   = errwrap.NewErrDef("r.warehouse.http_client.004", "failed to create warehouse via http call. %v")
	ErrRWarehouseHTTPClientUpdateWarehouse                                   = errwrap.NewErrDef("r.warehouse.http_client.005", "failed to update warehouse via http call. %v")
	ErrRWarehouseHTTPClientAddWarehouseProductStock                          = errwrap.NewErrDef("r.warehouse.http_client.006", "failed to add warehouse product stock via http call. %v")
	ErrRWarehouseHTTPClientTransferWarehouseProductStock                     = errwrap.NewErrDef("r.warehouse.http_client.007", "failed to transfer warehouse product stock via http call. %v")
)
