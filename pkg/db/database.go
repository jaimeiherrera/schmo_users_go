package db

// mockgen -source=pkg/db/database.go -destination=pkg/db/mock/database.go -package=mock
type Database interface {
	Set(key string, value map[string]interface{}) error
	Get(key string) (map[string]interface{}, error)
	GetAll() ([]map[string]interface{}, error)
	Delete(key string) error
}
