package ts_xorm

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
)

// conclusion: mysql orm insert multi records.
func TestMysqlORMInsert(t *testing.T) {
	// xorm原版标准方式创建引擎
	engine, err := xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			"root", "Teco@135", "100.10.22.82", "3306", "notebook"))
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	keys := []*Key{
		{
			Name:   "name",
			NbId:   sql.NullInt64{123, true},
			PubKey: "keeeee"},
		{
			Name:   "name2",
			NbId:   sql.NullInt64{123, true},
			PubKey: "keeeee"},
	}

	session.Insert(&keys)
	session.Commit()

	fmt.Printf("%s %d\n", "hello", 42)
}

// conclusion: not supported by pg to get last insert id by sql.
func TestPGInsertGetID(t *testing.T) {
	engine, err := xorm.NewPostgreSQL(
		"postgres://postgres:Teco@135@100.10.117.245:5432/notebook?sslmode=disable")
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	execute, err := session.SQL(
		"insert into notebook (display_name, namespace, name, nbc_id) values (?, ?, ?, ?) RETURNING id",
		"ut", "liuqw-test", "ut-name", 10000).Execute()
	if err != nil {
		panic(err)
	}
	id, _ := execute.LastInsertId()
	affected, _ := execute.RowsAffected()
	fmt.Printf("create notebook id:%v, affected:%v\n", id, affected)
	session.Commit()
}

// conclusion: pg insert to get id by ORM.
func TestPGInsertORMGetID(t *testing.T) {
	engine, err := xorm.NewPostgreSQL(
		"postgres://postgres:Teco@135@100.10.117.245:5432/notebook?sslmode=disable")
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	session := engine.NewSession()
	defer session.Close()
	session.Begin()

	notebook := &Notebook{
		DisplayName: "ut",
		Namespace:   "liuqw-test",
		Name:        "ut-name",
		NbcId:       10000,
	}

	res, err := session.InsertOne(notebook)
	if err != nil {
		panic(err)
	}
	session.Commit()
	fmt.Printf("create notebook id:%v, affected:%v\n", notebook.Id, res)
}

func TestPGQueryByField(t *testing.T) {
	engine, err := xorm.NewPostgreSQL(
		"postgres://postgres:Teco@135@100.10.14.220:5432/notebook?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	// engine.ShowSQL(true)

	for i := 0; i < 10; i++ {

		session := engine.NewSession()
		session.Begin()
		tempKeys := []*Key{
			{
				Name:   "test",
				NbId:   sql.NullInt64{10000, true},
				PubKey: "AAAAB3NzaC1yc2EAAAADAQABAAAAgQCTt8MMZ+4A463ZuVn2DOCT8mu3OKS3xmXWXmlAIeFYvvyXu2NuimWFunJqkHaw+QQ4n1GpqFAvBWi7ql6OAesB+AH7ERWgLbjZDGMI4oPeyGU+XOyxjbblo6ouRRhRNKYHcfueFmklNhwKjeU56dWOLDTjlgyRrdmI1ELhsHJwBQ==",
				UserId: "liuqw",
			},
		}
		num, err := session.Insert(&tempKeys)
		if err != nil {
			panic(err)
		}
		session.Close()
		fmt.Printf("affected: %v\n", num)

		var keys []*Key
		err = engine.Where("nb_id = ?", 10000).Find(&keys)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(keys))
	}
}
