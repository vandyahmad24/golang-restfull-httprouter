package simple

type Database struct {
	Name string
}
type DatabaseMysql Database
type DatabaseMonggo Database


func NewDatabaseMonggo() *DatabaseMonggo {
	return (*DatabaseMonggo)(&Database{Name: "monggo"})
}

func NewDatabaseMysql() *DatabaseMysql {
	return (*DatabaseMysql)(&Database{Name:"mysql"})
}


type DatabaseRepository struct {
	DatabaseMysql *DatabaseMysql
	DatabaseMonggo *DatabaseMonggo
}

func NewDatabaseRepository(databaseMysql *DatabaseMysql, databaseMonggo *DatabaseMonggo) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMysql: databaseMysql, DatabaseMonggo: databaseMonggo}
}
