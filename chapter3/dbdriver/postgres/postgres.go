package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"gitlab.wcxst.com/jormin/go-tools/log"
)

// Driver Postgres 驱动
type Driver struct {
}

// Open 打开Postgres连接
func (d Driver) Open(name string) (driver.Conn, error) {
	return nil, errors.New("unimplemented")
}

// 初始化注册驱动到sql中
func init() {
	sql.Register("postgres", &Driver{})
	log.Info("Register driver [%s] successful", "postgres")
}
