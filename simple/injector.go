//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleService, NewSimpleRepository)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSql,
		NewDatabaseRepository,
	)
	return nil
}
