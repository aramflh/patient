// Package initializers provides additional functions for initializing the main function
package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	//"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"os"
)

/*
DAO Documentation :https://gorm.io/gen/dao.html
*/

// DB is the database variable
var DB *gorm.DB

// ConnectToDB connect the server to a running instance of a postgresql database
func ConnectToDB() {
	var err error

	// Database connection string
	// values are imported from .env file
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"))

	// Initialize a *gorm.DB instance in DB var
	// err is set to nil (null) if the connection is successful
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Log an error if err is not nil (null)
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
}

/*
// GenStructFromDB Generates all the models from the DB tables
// See more: https://gorm.io/gen/database_to_structs.html & https://github.com/go-gorm/gen/tree/master/examples
func GenStructFromDB() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		// Generate global variable Q as DAO interface, then you can query data like: dal.Q.User.First()
		Mode: gen.WithDefaultQuery | gen.WithoutContext,
		// Output destination folder for the generator, default value: ./query
		OutPath: "dal",
		// Query code file name, default value: gen.go
		//OutFile: "queryCode.go",
		// if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		// if you want to assign field which has a default value in the `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// if you want to generate field with unsigned integer type, set FieldSignable true
		FieldSignable: true,
		// if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		// if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		// if you need unit tests for query code, set WithUnitTest true
		//WithUnitTest: true,
	})

	// Use the *gorm.DB instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(DB)

	// Generate structs from all tables of current database
	allTables := g.GenerateAllTable()
	// Changing  allTables []interface {} to a struct
	g.ApplyBasic(allTables...)

	// Execute the generator
	g.Execute()
}

*/
