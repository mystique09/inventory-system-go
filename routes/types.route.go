package routes

type (
	Route interface {
		GetAll() []interface{}
		GetOne() interface{}
		CreateOne() error
		UpdateOne() error
		DeleteOne() error
	}
)
