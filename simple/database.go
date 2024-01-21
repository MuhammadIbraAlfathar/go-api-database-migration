package simple

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabasePostgreSQL Database

func NewDatabasePostgreSql() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgeSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgeSQL, DatabaseMongoDB: mongoDB}
}
