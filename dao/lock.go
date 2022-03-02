package dao

import "github.com/EAHITechnology/raptor/eredis"

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
