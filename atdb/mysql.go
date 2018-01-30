package atdb
import (
	"database/sql"
	"io/ioutil"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql(config DbConfig)(*sql.DB,error) {
	db, err := sql.Open("mysql", config.Name+":"+config.Pwd+"@tcp("+config.Ip+")/"+config.DbName+"?charset=utf8")
	return db,err
}

/**
	读取数据库的配置文件
 */
func ReadDbConfig(path string) (*DbConfig,error) {
	res,err := ioutil.ReadFile(path)
	if err!= nil {
		return nil,err
	}
	var config DbConfig

	err = json.Unmarshal(res,&config)
	if err != nil {
		return nil,err
	}
	return &config,nil
}
type DbConfig struct{
	Name string
	Pwd string
	Ip string
	DbName string
}
