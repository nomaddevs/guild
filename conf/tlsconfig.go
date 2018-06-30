package conf

import (
	//"database/sql"
	//"errors"

	_ "github.com/go-sql-driver/mysql"
)

type TLSconfig struct {
	Addr     string
	CertFile string
	KeyFile  string
}
