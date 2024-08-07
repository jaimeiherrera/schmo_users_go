package db

import "errors"

type LocalDB struct {
	Data []map[string]interface{}
}

func NewLocalDB() *LocalDB {
	return &LocalDB{
		Data: []map[string]interface{}{},
	}
}

func (ldb *LocalDB) Set(key string, value map[string]interface{}) error {
	usr := map[string]interface{}{"id": key}
	for k, v := range value {
		usr[k] = v
	}
	ldb.Data = append(ldb.Data, usr)
	return nil
}

func (ldb *LocalDB) Get(key string) (map[string]interface{}, error) {
	for _, v := range ldb.Data {
		if v["id"] == key {
			return v, nil
		}
	}
	return nil, nil
}

func (ldb *LocalDB) GetAll() ([]map[string]interface{}, error) {
	return ldb.Data, nil
}

func (ldb *LocalDB) Delete(key string) error {
	for i, v := range ldb.Data {
		if v["id"] == key {
			ldb.Data = append(ldb.Data[:i], ldb.Data[i+1:]...)
			return nil
		}
	}
	return errors.New("key not found")
}
