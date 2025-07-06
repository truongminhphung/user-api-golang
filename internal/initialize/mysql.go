package initialize

import (
	"fmt"
	"time"
	"user-api/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func panicIfError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// InitMySQL initializes the MySQL database connection using configuration from global.Config,
// sets up the connection pool, and assigns the database instance to global.Mdb.
// It panics if the connection cannot be established.
func InitMySQL() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.Dbname)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	panicIfError(err, "Failed to connect to database")
	global.Logger.Info("Connected to MySQL database successfully")
	global.Mdb = db

	SetPool()
}

// SetPool configures the database connection pool settings such as max idle connections,
// max open connections, and connection max lifetime based on the application's configuration.
func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	panicIfError(err, "Failed to get database instance")
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
