// mysql db drives
package driver

//noinspection GoUnresolvedReference
import (
	"fmt"
	"gin/package/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"time"
)

//主库初始化
func initMasterDb() {
	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.ConfigParam.MasterDB.TablePrefix + defaultTableName
	}

	//链接配置
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.ConfigParam.MasterDB.DbUser,
		setting.ConfigParam.MasterDB.DbPwd,
		setting.ConfigParam.MasterDB.Host,
		setting.ConfigParam.MasterDB.Port,
		setting.ConfigParam.MasterDB.DbName,
		setting.ConfigParam.MasterDB.DbCharset,
	)

	// connect and open db connection
	MasterDb, DbErr = gorm.Open("mysql", dbDSN)

	//全局禁用表名复数
	MasterDb.SingularTable(true)

	if DbErr != nil {
		panic("database data source name error: " + DbErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(setting.ConfigParam.MasterDB.DbMaxIdleCconns)
	MasterDb.DB().SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(setting.ConfigParam.MasterDB.DbMaxIdleCconns)
	MasterDb.DB().SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(setting.ConfigParam.MasterDB.DbMaxLifetimeConns)
	MasterDb.DB().SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if DbErr = MasterDb.DB().Ping(); nil != DbErr {
		panic("database connect failed: " + DbErr.Error())
	}
}

//从库初始化
func initSlaveDb() {
	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.ConfigParam.SlaveDB.TablePrefix + defaultTableName
	}

	//链接配置
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.ConfigParam.SlaveDB.DbUser,
		setting.ConfigParam.SlaveDB.DbPwd,
		setting.ConfigParam.SlaveDB.Host,
		setting.ConfigParam.SlaveDB.Port,
		setting.ConfigParam.SlaveDB.DbName,
		setting.ConfigParam.SlaveDB.DbCharset,
	)

	// connect and open db connection
	SlaveDb, DbErr = gorm.Open("mysql", dbDSN)

	//全局禁用表名复数
	SlaveDb.SingularTable(true)

	if DbErr != nil {
		panic("database data source name error: " + DbErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(setting.ConfigParam.SlaveDB.DbMaxIdleCconns)
	SlaveDb.DB().SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(setting.ConfigParam.SlaveDB.DbMaxIdleCconns)
	SlaveDb.DB().SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(setting.ConfigParam.SlaveDB.DbMaxLifetimeConns)
	SlaveDb.DB().SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	// check db connection at once avoid connect failed
	// else error will be reported until db first sql operate
	if DbErr = SlaveDb.DB().Ping(); nil != DbErr {
		panic("database connect failed: " + DbErr.Error())
	}
}
