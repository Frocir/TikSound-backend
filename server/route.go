package server

import (
	"TikSound-backend/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(cfg *config.TkConfig) (*gin.Engine, error) {
	var r *gin.Engine
	if cfg.Mode == gin.ReleaseMode {
		r = gin.New()
		// 使用生产环境的中间件，日期记录器等
		r.Use()
	}

	r = gin.Default()
	if err := installMiddleware(r, cfg); err != nil {
		return nil, err
	}
	if err := installController(r, cfg); err != nil {
		return nil, err
	}
	return r, nil
}
func installMiddleware(r *gin.Engine, cfg *config.TkConfig) error {
	return nil
}
func installController(r *gin.Engine, cfg *config.TkConfig) error {
	//mysqlIns, err := mysql.GetMySQLFactory(cfg.Mysql)
	//if err != nil {
	//	return err
	//}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "I'm health!"})
	})
	return nil
}
