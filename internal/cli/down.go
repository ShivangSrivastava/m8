package cli

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ShivangSrivastava/m8/internal/app"
	"github.com/ShivangSrivastava/m8/internal/infra/db"
	"github.com/ShivangSrivastava/m8/internal/infra/fs"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

// Rolls back the latest applied migration by loading its .down.sql file.
// Ensures database connection and migration components are properly wired.
// Prints the reverted migration name for user feedback.
func Down(cmd *cobra.Command, args []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	connStr := "user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" host=" + os.Getenv("POSTGRES_HOST") +
		" sslmode=disable"

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	defer dbConn.Close()

	repo := &db.DBRepo{DB: dbConn}
	loader := &fs.FileLoader{Dir: "migrations", Direction: fs.Down}

	service := &app.ApplyService{
		Repo:   repo,
		Loader: loader,
	}

	if err := service.Revert(); err != nil {
		log.Fatalln("Error applying migrations:", err)
	}

	if name := service.RevertedMigration; name != "" {
		fmt.Println("Revert migrations:", name)
	} else {
		fmt.Println("No migrations revert.")
	}

	return nil
}
