package utilities_db

import (
    "database/sql"
    "fmt"
    _ "io/ioutil"
    "os"
    "encoding/json"
    _ "reflect"

    "bitbucket.org/pchandev/gamecatalogue-go/models"
    _ "bitbucket.org/pchandev/gamecatalogue-go/_vendor/github.com/go-sql-driver/mysql"
    "bitbucket.org/pchandev/gamecatalogue-go/_vendor/github.com/jmoiron/modl"
)

type database_config struct {
    DatabaseUser       string
    DatabasePassword   string
    DatabaseHost       string
    DatabaseName       string
}

// example dsn: testuser:password@localhost/test
func InitializeDbMap() *modl.DbMap {
    var dsn string
    config := database_config{}

    if config_file, err := os.Open( "../config/db.json" ); err != nil {
        panic( "Could not find db.json in config folder" )
    } else {
        decoder := json.NewDecoder( config_file )

        if err := decoder.Decode( &config ) ; err != nil {
            panic( "Could not decode config file: " + err.Error() )
        }

        dsn = fmt.Sprintf( "%s:%s@/%s", config.DatabaseUser, config.DatabasePassword, config.DatabaseName )
    }

    database_map := new_db_map( dsn )

    database_map.AddTableWithName( models_db.Mediatypes{}, "mediatypes" ).SetKeys( false, "mediatypeid" )

    for _, value := range models_db.GetAllDbStructs() {
        database_map.AddTable( value.DatabaseObject ).SetKeys( value.IsAutoIncrementKey, value.DatabaseKey )
    }

    return database_map
}

func new_db_map( dsn string ) *modl.DbMap {
    dialect, driver := get_dialect_and_driver( "mysql" )

    return modl.NewDbMap( connect( driver, dsn ), dialect )
}

func connect( driver string, dsn string ) *sql.DB {
    if dsn == "" {
        panic( "DSN is empty." )
    }

    if database_object, err := sql.Open( driver, dsn ); err != nil {
        panic( "Error connecting to db: " + err.Error() )
    } else {
        if err = database_object.Ping(); err != nil {
            panic( "Error connecting to db: " + err.Error() )
        }

        return database_object
    }
}

func get_dialect_and_driver( database_type string ) ( modl.Dialect, string ) {
    switch database_type {
        case "mysql":
            return modl.MySQLDialect{ "InnoDB", "UTF8" }, "mysql"
            /* TODO: FUTURE IMPLEMENTATION
               case "postgres":
                   return PostgresDialect{}, "postgres"
               case "sqlite":
                   return SqliteDialect{}, "sqlite3"
            */
    }

    panic( "database_type is invalid." )
}