package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx){
	err := recover()
	if err != nil{
		errorRollback := tx.Rollback()
		PanicError(errorRollback)
	}else{
		errorCommit := tx.Commit()
		PanicError(errorCommit)
	}
}
