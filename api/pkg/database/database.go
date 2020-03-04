package database

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}
