package dao

import (
	"github.com/EAHITechnology/raptor/emysql"
	"github.com/jinzhu/gorm"
)

func InsertDemoTask(d DemoModel) (int64, error) {
	client, err := emysql.GetClient("cage-mysql")
	if err != nil {
		return 0, err
	}

	if err := client.GetMaster().Table("cage_task").Create(&d).Error; err != nil {
		return 0, err
	}

	return d.Id, nil
}

func GetOneDemoInfo(id int64) (*DemoModel, error) {
	client, err := emysql.GetClient("cage-mysql")
	if err != nil {
		return nil, err
	}

	d := &DemoModel{}
	if err := client.GetMaster().Table("cage_task").Where("id = ?", id).Find(d).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return d, nil
}

func UpdateOneDemoInfo(d *DemoModel) (int64, error) {
	client, err := emysql.GetClient("cage-mysql")
	if err != nil {
		return 0, err
	}

	data := make(map[string]interface{})
	data["info"] = d.Info

	db := client.GetMaster()
	if err := db.Debug().Table("cage_task").
		Model(&DemoModel{}).
		Where("id = ?", d.Id).
		Updates(data).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return 0, nil
		}
		return 0, err
	}

	return db.RowsAffected, nil
}
