//go:build wireinject
// +build wireinject
package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializeService(isError bool) (*SimpleService, error){
	wire.Build(
		NewSimpleService, NewSimpleRepository,
		)
	return nil, nil
}

func InitalizeDatabase() *DatabaseRepository{
	wire.Build(
		NewDatabaseMysql, NewDatabaseMonggo, NewDatabaseRepository,
		)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService{
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}
//salah
//func InitializeHelloService() *HelloService{
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
	)

func InitializeHelloService() *HelloService{
	wire.Build(helloSet, NewHelloService)
	return  nil
}

func InitializeFooBar() *FooBar{
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar),"*"))
	return nil
}


func IntializeFooBarWithValue() *FooBar{

	return nil
}


var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar  {
	wire.Build(wire.Value(fooValue),wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}


func InitializeReader() io.Reader{
	wire.Build(wire.InterfaceValue(new(io.Reader),os.Stdin))
	return nil
}

func InitializeConfiguration() *Configuration{
	 //application := NewApplication()
	 //configuration := application.Configuration
	 //return configuration
	wire.Build(NewApplication,wire.FieldsOf(new(*Application),"Configuration"))
	return nil
}

func IntiliazeConnection(name string) (*Connection, func()){
	wire.Build(NewConnection, NewFile)
	return nil, nil
}