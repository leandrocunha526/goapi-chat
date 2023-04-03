package helper

import (
	"database/sql"
	"fmt"
)

func PanicError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func RollbackErr(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		fmt.Println(errRollback)
		PanicError(errRollback)
	} else {
		errCommit := tx.Commit()
		fmt.Println(errCommit)
		PanicError(errCommit)
	}
}
