package db

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicateKeyError(err error) bool {
	var mysqlErr = new(mysql.MySQLError)
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return false
}
