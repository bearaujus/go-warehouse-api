package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds the environment variables
type Config struct {
	ServiceOrderContainerName              string
	ServiceOrderPort                       string
	ServiceOrderImageVersion               string
	ServiceOrderHTTPCallAuthKey            string
	ServiceOrderHTTPCallAuthTTL            time.Duration
	ServiceOrderHTTPCallTimeout            time.Duration
	ServiceOrderExpirationTTL              time.Duration
	ServiceOrderStockReservationGCCronSpec string

	ServiceProductContainerName   string
	ServiceProductPort            string
	ServiceProductImageVersion    string
	ServiceProductHTTPCallAuthKey string
	ServiceProductHTTPCallAuthTTL time.Duration
	ServiceProductHTTPCallTimeout time.Duration

	ServiceShopContainerName   string
	ServiceShopPort            string
	ServiceShopImageVersion    string
	ServiceShopHTTPCallAuthKey string
	ServiceShopHTTPCallAuthTTL time.Duration
	ServiceShopHTTPCallTimeout time.Duration

	ServiceUserContainerName   string
	ServiceUserPort            string
	ServiceUserImageVersion    string
	ServiceUserHTTPCallAuthKey string
	ServiceUserHTTPCallAuthTTL time.Duration
	ServiceUserHTTPCallTimeout time.Duration
	ServiceUserAuthSecretKey   string
	ServiceUserAuthTTL         time.Duration

	ServiceWarehouseContainerName   string
	ServiceWarehousePort            string
	ServiceWarehouseImageVersion    string
	ServiceWarehouseHTTPCallAuthKey string
	ServiceWarehouseHTTPCallAuthTTL time.Duration
	ServiceWarehouseHTTPCallTimeout time.Duration

	PostgresUser      string
	PostgresPassword  string
	PostgresDB        string
	PostgresInitRetry int
	PostgresInitDelay time.Duration

	PostgresMasterContainerName      string
	PostgresMasterPort               string
	PostgresMasterSchemaInitPath     string
	PostgresMasterWalLevel           string
	PostgresMasterHotStandby         string
	PostgresMasterMaxReplication     string
	PostgresMasterHotStandbyFeedback string
	PostgresMasterLoggingCollector   string
	PostgresMasterLogStatement       string

	PostgresReplicaInitScriptName   string
	PostgresReplicaLoggingCollector string
	PostgresReplicaLogStatement     string

	PostgresReplica1ContainerName string
	PostgresReplica1Port          string

	PostgresReplica2ContainerName string
	PostgresReplica2Port          string

	PostgresReplica3ContainerName string
	PostgresReplica3Port          string

	RedisContainerName string
	RedisPort          string
	RedisPass          string
	RedisBind          string
	RedisLogLevel      string
	RedisAppendFsync   string
	RedisTTL           time.Duration
	RedisInitRetry     int
	RedisInitDelay     time.Duration

	LocalEnvPath    string
	LocalLogPath    string
	LocalSchemaPath string
	LocalScriptPath string

	ContainerBinPath    string
	ContainerScriptPath string
	ContainerLogPath    string

	LogExtension string
	TZ           string
}

func Read() *Config {
	cfg := &Config{
		ServiceOrderContainerName:              os.Getenv("SERVICE_ORDER_CONTAINER_NAME"),
		ServiceOrderPort:                       os.Getenv("SERVICE_ORDER_PORT"),
		ServiceOrderImageVersion:               os.Getenv("SERVICE_ORDER_IMAGE_VERSION"),
		ServiceOrderHTTPCallAuthKey:            os.Getenv("SERVICE_ORDER_HTTP_CALL_AUTH_KEY"),
		ServiceOrderHTTPCallAuthTTL:            osGetenvToDuration("SERVICE_ORDER_HTTP_CALL_AUTH_TTL"),
		ServiceOrderHTTPCallTimeout:            osGetenvToDuration("SERVICE_ORDER_HTTP_CALL_TIMEOUT"),
		ServiceOrderExpirationTTL:              osGetenvToDuration("SERVICE_ORDER_EXPIRATION_TTL"),
		ServiceOrderStockReservationGCCronSpec: os.Getenv("SERVICE_ORDER_STOCK_RESERVATION_GC_CRON_SPEC"),

		ServiceProductContainerName:   os.Getenv("SERVICE_PRODUCT_CONTAINER_NAME"),
		ServiceProductPort:            os.Getenv("SERVICE_PRODUCT_PORT"),
		ServiceProductImageVersion:    os.Getenv("SERVICE_PRODUCT_IMAGE_VERSION"),
		ServiceProductHTTPCallAuthKey: os.Getenv("SERVICE_PRODUCT_HTTP_CALL_AUTH_KEY"),
		ServiceProductHTTPCallAuthTTL: osGetenvToDuration("SERVICE_PRODUCT_HTTP_CALL_AUTH_TTL"),
		ServiceProductHTTPCallTimeout: osGetenvToDuration("SERVICE_PRODUCT_HTTP_CALL_TIMEOUT"),

		ServiceShopContainerName:   os.Getenv("SERVICE_SHOP_CONTAINER_NAME"),
		ServiceShopPort:            os.Getenv("SERVICE_SHOP_PORT"),
		ServiceShopImageVersion:    os.Getenv("SERVICE_SHOP_IMAGE_VERSION"),
		ServiceShopHTTPCallAuthKey: os.Getenv("SERVICE_SHOP_HTTP_CALL_AUTH_KEY"),
		ServiceShopHTTPCallAuthTTL: osGetenvToDuration("SERVICE_SHOP_HTTP_CALL_AUTH_TTL"),
		ServiceShopHTTPCallTimeout: osGetenvToDuration("SERVICE_SHOP_HTTP_CALL_TIMEOUT"),

		ServiceUserContainerName:   os.Getenv("SERVICE_USER_CONTAINER_NAME"),
		ServiceUserPort:            os.Getenv("SERVICE_USER_PORT"),
		ServiceUserImageVersion:    os.Getenv("SERVICE_USER_IMAGE_VERSION"),
		ServiceUserHTTPCallAuthKey: os.Getenv("SERVICE_USER_HTTP_CALL_AUTH_KEY"),
		ServiceUserHTTPCallAuthTTL: osGetenvToDuration("SERVICE_USER_HTTP_CALL_AUTH_TTL"),
		ServiceUserHTTPCallTimeout: osGetenvToDuration("SERVICE_USER_HTTP_CALL_TIMEOUT"),
		ServiceUserAuthSecretKey:   os.Getenv("SERVICE_USER_AUTH_SECRET_KEY"),
		ServiceUserAuthTTL:         osGetenvToDuration("SERVICE_USER_AUTH_TTL"),

		ServiceWarehouseContainerName:   os.Getenv("SERVICE_WAREHOUSE_CONTAINER_NAME"),
		ServiceWarehousePort:            os.Getenv("SERVICE_WAREHOUSE_PORT"),
		ServiceWarehouseImageVersion:    os.Getenv("SERVICE_WAREHOUSE_IMAGE_VERSION"),
		ServiceWarehouseHTTPCallAuthKey: os.Getenv("SERVICE_WAREHOUSE_HTTP_CALL_AUTH_KEY"),
		ServiceWarehouseHTTPCallAuthTTL: osGetenvToDuration("SERVICE_WAREHOUSE_HTTP_CALL_AUTH_TTL"),
		ServiceWarehouseHTTPCallTimeout: osGetenvToDuration("SERVICE_WAREHOUSE_HTTP_CALL_TIMEOUT"),

		PostgresUser:      os.Getenv("POSTGRES_USER"),
		PostgresPassword:  os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:        os.Getenv("POSTGRES_DB"),
		PostgresInitRetry: osGetenvToInt("POSTGRES_INIT_RETRY"),
		PostgresInitDelay: osGetenvToDuration("POSTGRES_INIT_DELAY"),

		PostgresMasterContainerName:      os.Getenv("POSTGRES_MASTER_CONTAINER_NAME"),
		PostgresMasterPort:               os.Getenv("POSTGRES_MASTER_PORT"),
		PostgresMasterSchemaInitPath:     os.Getenv("POSTGRES_MASTER_SCHEMA_INIT_PATH"),
		PostgresMasterWalLevel:           os.Getenv("POSTGRES_MASTER_WAL_LEVEL"),
		PostgresMasterHotStandby:         os.Getenv("POSTGRES_MASTER_HOT_STANDBY"),
		PostgresMasterMaxReplication:     os.Getenv("POSTGRES_MASTER_MAX_REPLICATION"),
		PostgresMasterHotStandbyFeedback: os.Getenv("POSTGRES_MASTER_HOT_STANDBY_FEEDBACK"),
		PostgresMasterLoggingCollector:   os.Getenv("POSTGRES_MASTER_LOGGING_COLLECTOR"),
		PostgresMasterLogStatement:       os.Getenv("POSTGRES_MASTER_LOG_STATEMENT"),

		PostgresReplicaInitScriptName:   os.Getenv("POSTGRES_REPLICA_INIT_SCRIPT_NAME"),
		PostgresReplicaLoggingCollector: os.Getenv("POSTGRES_REPLICA_LOGGING_COLLECTOR"),
		PostgresReplicaLogStatement:     os.Getenv("POSTGRES_REPLICA_LOG_STATEMENT"),

		PostgresReplica1ContainerName: os.Getenv("POSTGRES_REPLICA_1_CONTAINER_NAME"),
		PostgresReplica1Port:          os.Getenv("POSTGRES_REPLICA_1_PORT"),

		PostgresReplica2ContainerName: os.Getenv("POSTGRES_REPLICA_2_CONTAINER_NAME"),
		PostgresReplica2Port:          os.Getenv("POSTGRES_REPLICA_2_PORT"),

		PostgresReplica3ContainerName: os.Getenv("POSTGRES_REPLICA_3_CONTAINER_NAME"),
		PostgresReplica3Port:          os.Getenv("POSTGRES_REPLICA_3_PORT"),

		RedisContainerName: os.Getenv("REDIS_CONTAINER_NAME"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPass:          os.Getenv("REDIS_PASS"),
		RedisBind:          os.Getenv("REDIS_BIND"),
		RedisLogLevel:      os.Getenv("REDIS_LOG_LEVEL"),
		RedisAppendFsync:   os.Getenv("REDIS_APPEND_FSYNC"),
		RedisTTL:           osGetenvToDuration("REDIS_TTL"),
		RedisInitRetry:     osGetenvToInt("REDIS_INIT_RETRY"),
		RedisInitDelay:     osGetenvToDuration("REDIS_INIT_DELAY"),

		LocalEnvPath:    os.Getenv("LOCAL_ENV_PATH"),
		LocalLogPath:    os.Getenv("LOCAL_LOG_PATH"),
		LocalSchemaPath: os.Getenv("LOCAL_SCHEMA_PATH"),
		LocalScriptPath: os.Getenv("LOCAL_SCRIPT_PATH"),

		ContainerBinPath:    os.Getenv("CONTAINER_BIN_PATH"),
		ContainerScriptPath: os.Getenv("CONTAINER_SCRIPT_PATH"),
		ContainerLogPath:    os.Getenv("CONTAINER_LOG_PATH"),

		LogExtension: os.Getenv("LOG_EXTENSION"),
		TZ:           os.Getenv("TZ"),
	}
	return cfg
}

func osGetenvToInt(k string) int {
	ret, _ := strconv.Atoi(os.Getenv(k))
	return ret
}

func osGetenvToDuration(k string) time.Duration {
	ret, _ := time.ParseDuration(os.Getenv(k))
	return ret
}
