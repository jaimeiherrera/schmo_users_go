package db

type Database interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	GetAll() (map[string]interface{}, error)
	Delete(key string) error
}
