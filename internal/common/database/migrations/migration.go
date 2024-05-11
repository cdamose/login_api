package migrations

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ExecMigration(user_name string, password string, host string, migration_db string,file_path string) {
	db, err := sql.Open("postgres", "postgres://postgres:Looser1997$@postgres:5432/auth?sslmode=disable")

	fmt.Println("debug")
	fmt.Println(err)
	//Config.MaxAllowedPacket
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	fmt.Println(err)
	fmt.Println("debug 2")
	file := file_path //"file://./db/migrations/auth"
	previousDir := filepath.Dir(filepath.Dir(file))
	print(previousDir)

	m, err := migrate.NewWithDatabaseInstance(
		file,
		"postgres",
		driver,
	)
	fmt.Println("debug 3")
	fmt.Println(err)
	err = m.Up()

	fmt.Println(err)

}
