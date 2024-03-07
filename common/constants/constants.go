package constants

import "time"

const (
	// auth
	JWTValue = "MTAxNTkwMTg1Mw=="
	StartID  = 10000

	// service name

	UserServiceName = "user"
	ApiServiceName  = "api"
	GameServiceName = "game"
	InteractService = "interact"

	// snowflake

	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// limit

	MaxConnections  = 1000
	MaxQPS          = 100
	MaxVideoSize    = 300000
	MaxListLength   = 100
	MaxIdleConn     = 10
	MaxGoroutines   = 10
	MaxOpenConn     = 100
	ConnMaxLifetime = 10 * time.Second

	// table

	UserTableName = "user"

	// page

	PageNum  = 1
	PageSize = 10

	// rpc

	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond

	// match

	MatchInitialDelay  = 50 * time.Millisecond
	MaxMatchRetryDelay = 20 * time.Second
)
