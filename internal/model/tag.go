package model

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

/**
 新增标签
 */
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

/**
 更新标签
 */
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&t).Where("is_del = ?", 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 删除标签
 */
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

/**
 计算标签总数
 */
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

/**
 获取数据列表
 */
func (t Tag) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}