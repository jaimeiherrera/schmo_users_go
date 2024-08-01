package db

import "errors"

type LocalDB struct {
	Data map[string]interface{}
}

func NewLocalDB() *LocalDB {
	return &LocalDB{
		Data: make(map[string]interface{}),
	}
}

func (ldb *LocalDB) Set(key string, value interface{}) error {
	ldb.Data[key] = value
	return nil
}

func (ldb *LocalDB) Get(key string) (interface{}, error) {
	if v, ok := ldb.Data[key]; !ok {
		return nil, errors.New("key not found")
	} else {
		return v, nil
	}
}

func (ldb *LocalDB) GetAll() (map[string]interface{}, error) {
	return ldb.Data, nil
}

func (ldb *LocalDB) Delete(key string) error {
	delete(ldb.Data, key)
	return nil
}
