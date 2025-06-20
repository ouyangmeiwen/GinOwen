package global

import (
	"context"
	"database/sql"
	"sync"

	"GINOWEN/global/cfg"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OWEN_DB     *gorm.DB
	SQL_DB      *sql.DB
	OWEN_DBList map[string]*gorm.DB
	OWEN_REDIS  redis.UniversalClient
	OWEN_MONGO  *mongo.Client
	OWEN_CONFIG cfg.YarmConfig
	OWEN_LOCK   sync.RWMutex
	OWEN_LOG    *zap.Logger
	Ctx         = context.Background()
)
