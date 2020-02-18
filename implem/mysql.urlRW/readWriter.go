package urlRw

import (
	"fmt"
	"github.com/spf13/viper"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/saeidraei/go-realworld-clean/domain"
	"github.com/saeidraei/go-realworld-clean/uc"
)

type rw struct {
	db *sql.DB
}

func New() uc.UrlRW {
	db, err := sql.Open("mysql", viper.GetString("mysql.user")+":"+viper.GetString("mysql.password")+"@tcp("+viper.GetString("mysql.host")+":"+viper.GetString("mysql.port")+")/"+viper.GetString("mysql.database"))
	if err!=nil{
		fmt.Println(err)
		panic(err.Error())
	}
	return rw{
		db: db,
	}
}
func (rw rw) Create(url domain.Url) (*domain.Url, error) {

	ins, err := rw.db.Query("insert into url(ID,Address) values(?,?)",url.ID,url.Address)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer ins.Close()

	return rw.GetByID(url.ID)
}

func (rw rw) Save(url domain.Url) (*domain.Url, error) {
	//not used
	return rw.GetByID(url.ID)
}

func (rw rw) GetByID(id string) (*domain.Url, error) {

	var url domain.Url
	err := rw.db.QueryRow("select ID,Address from url where ID =? ",id).Scan(&url.ID,&url.Address)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return &url,nil
}

func (rw rw) Delete(slug string) error {
	//rw.store.Delete(slug)
	return nil
}
