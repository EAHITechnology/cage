package dao

import "github.com/EAHITechnology/raptor/eredis"

func Set(key string, value interface{}) error {
	rds, err := eredis.GetClient("cage-redis")
	if err != nil {
		return err
	}
	return rds.Set(key, value)
}

func Get(key string) (string, error) {
	rds, err := eredis.GetClient("cage-redis")
	if err != nil {
		return "", err
	}
	return rds.GetString(key)
}

func SetRedisLock(key string, ex int64) (int64, bool, error) {
	rds, err := eredis.GetClient("cage-redis")
	if err != nil {
		return 0, false, err
	}
	version, ok, err := rds.Lock(key, ex)
	if err != nil {
		return 0, false, err
	}
	if !ok {
		return 0, false, nil
	}
	return version, true, nil
}

func DelRedisLock(key string, version int64) error {
	rds, err := eredis.GetClient("cage-redis")
	if err != nil {
		return err
	}
	_, err = rds.Unlock(key, version)
	return err
}
