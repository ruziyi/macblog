package models

import(
    "strings"
    "os"
    "fmt"
    "path"
    
	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    "github.com/weisd/macblog/modules/setting"
)

var (
    Engine *xorm.Engine

    DbCfg struct {
        Type, Host, Name, User, Pwd, Path, SslMode string
    }

    tables []interface{}

    UserSQLite3 bool
)


func init() {
    tables = append(tables, new(User))
}

func LoadModelsConfig() {
    DbCfg.Type = setting.Cfg.MustValue("database", "DB_TYPE")
    if DbCfg.Type == "sqlite3" {
        UserSQLite3 = true
    }
    DbCfg.Host = setting.Cfg.MustValue("database", "HOST")
    DbCfg.Name = setting.Cfg.MustValue("database", "NAME")
    DbCfg.User = setting.Cfg.MustValue("database", "USER")
    DbCfg.Pwd = setting.Cfg.MustValue("database", "PASSWD")
    if len(DbCfg.Pwd) == 0 {
        DbCfg.Pwd = setting.Cfg.MustValue("database", "PASSWD")
    }
    DbCfg.SslMode = setting.Cfg.MustValue("database", "SSL_MODE")
    // sqlite3
    DbCfg.Path = setting.Cfg.MustValue("database", "PATH", "data/macblog.db")
}

func SetEngine() (err error) {
    switch DbCfg.Type {
    case "mysql":
        Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", DbCfg.User, DbCfg.Pwd, DbCfg.Host, DbCfg.Name))
	case "postgres":
		var host, port = "127.0.0.1", "5432"
		fields := strings.Split(DbCfg.Host, ":")
		if len(fields) > 0 && len(strings.TrimSpace(fields[0])) > 0 {
			host = fields[0]
		}
		if len(fields) > 1 && len(strings.TrimSpace(fields[1])) > 0 {
			port = fields[1]
		}
		Engine, err = xorm.NewEngine("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			DbCfg.User, DbCfg.Pwd, host, port, DbCfg.Name, DbCfg.SslMode))
	case "sqlite3":
		os.MkdirAll(path.Dir(DbCfg.Path), os.ModePerm)
		Engine, err = xorm.NewEngine("sqlite3", DbCfg.Path)
	default:
		return fmt.Errorf("Unknown database type: %s", DbCfg.Type)
	}
	if err != nil {
		return fmt.Errorf("models.init(fail to conntect database): %v", err)
	}

	// WARNNING: for serv command, MUST remove the output to os.stdout,
	// so use log file to instead print to stdout.
	logPath := path.Join(setting.LogRootPath, "xorm.log")
	os.MkdirAll(path.Dir(logPath), os.ModePerm)

	f, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("models.init(fail to create xorm.log): %v", err)
	}
	Engine.Logger = xorm.NewSimpleLogger(f)

	Engine.ShowSQL = true
	Engine.ShowDebug = true
	Engine.ShowErr = true
	return nil
}

func NewEngine() (err error) {
    if err = SetEngine(); err != nil {
        return err
    }

    if err = Engine.Sync2(tables...); err != nil {
        return fmt.Errorf("sync database struct error: %v\n", err)
    }

    return nil
}

func DumpDatabase(filePath string) error {
    return Engine.DumpAllToFile(filePath)
}
