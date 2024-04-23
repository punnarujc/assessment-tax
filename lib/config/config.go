package config

import (
	"fmt"
	"os"
)

type Config interface {
	GetPort() string
	GetDb() string
	GetAdminCredential() AdminCfg
}

type configImpl struct {
	Port     string
	Db       string
	AdminCfg AdminCfg
}

type AdminCfg struct {
	AdminUsername string
	AdminPassword string
}

func New() Config {
	return &configImpl{
		Port: os.Getenv(PORT),
		Db:   os.Getenv(DB),
		AdminCfg: AdminCfg{
			AdminUsername: os.Getenv(ADMIN_USERNAME),
			AdminPassword: os.Getenv(ADMIN_PASSWORD),
		},
	}
}

func (c *configImpl) GetPort() string {
	return fmt.Sprintf(":%s", c.Port)
}

func (c *configImpl) GetDb() string {
	return c.Db
}

func (c *configImpl) GetAdminCredential() AdminCfg {
	return c.AdminCfg
}
