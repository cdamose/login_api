package migrations

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ExecMigration(user_name string, password string, host string, migration_db string) {
	db, err := sql.Open("mysql", user_name+":"+password+"@tcp("+host+":3306)/"+migration_db+"?multiStatements=true")
	fmt.Println("debug")
	fmt.Println(err)
	//Config.MaxAllowedPacket
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	fmt.Println(err)
	fmt.Println("debug 2")
	file := "file://./db/migrations"
	previousDir := filepath.Dir(filepath.Dir(file))
	print(previousDir)

	m, err := migrate.NewWithDatabaseInstance(
		file,
		"mysql",
		driver,
	)
	fmt.Println("debug 3")
	fmt.Println(err)
	err = m.Up()

	fmt.Println(err)

}
