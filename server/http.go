package server

import (
	"github.com/gin-gonic/gin"
)

func CreateServer() (*gin.Engine, error) {
	engine := gin.Default()
	setRoutes(engine)
	err := bootstrap()
	if err != nil {
		return nil, err
	}
	return engine, nil
}
