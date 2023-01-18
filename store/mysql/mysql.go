package mysql

import (
	"TikSound-backend/config"
	"TikSound-backend/store"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type dataStore struct {
	db *gorm.DB
}

func (ds *dataStore) User() store.UserStore {
	return newUser(ds)
}

func (ds *dataStore) Video() store.VideoStore {
	return newVideo(ds)
}

func (ds *dataStore) Comment() store.CommentStore {
	return newComment(ds)
}

func (ds *dataStore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.New("get gorm db instance failed")
	}
	return db.Close()
}

var (
	mysqlFactory *dataStore
	once         sync.Once
)

// GetMySQLFactory 单例模式创建存储实例
func GetMySQLFactory(cfg *config.Mysql) (store.Factory, error) {
	var errOr error
	once.Do(func() {
		dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Database,
			true,
			"local",
		)
		dbIns, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			errOr = err
			return
		}
		sqlDB, err := dbIns.DB()
		if err != nil {
			errOr = err
			return
		}
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(cfg.MaxConnectionLifeTime)

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
		mysqlFactory = &dataStore{db: dbIns}
	})

	return mysqlFactory, errOr
}
