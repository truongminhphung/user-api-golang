package initialize

import (
	"database/sql"
	"fmt"
	"time"
	"user-api/global"

	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
)

func panicIfError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// InitMySQL initializes the MySQL database connection using configuration from global.Config,
// sets up the connection pool, and assigns the database instance to global.MySQLDB.
// It panics if the connection cannot be established.
func InitMySQL() {
	global.Logger.Info("Initializing MySQL database connection...")
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.UserName, m.Password, m.Host, m.Port, m.Dbname)
	// var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.Dbname)

	db, err := sql.Open("mysql", dsn)
	panicIfError(err, "Failed to connect to database")
	// Test the connection to ensure it is valid
	if err := db.Ping(); err != nil {
		global.Logger.Error("Failed to ping MySQL database", zap.Error(err))
	}
	global.Logger.Info("Connected to MySQL database successfully")
	global.MySQLDB = db

	SetConnectionPool(db)

}

// SetConnectionPool configures the database connection pool settings such as max idle connections,
// max open connections, and connection max lifetime based on the application's configuration.
func SetConnectionPool(db *sql.DB) {
	m := global.Config.Mysql
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	global.Logger.Info("MySQL connection pool configured successfully")
}
