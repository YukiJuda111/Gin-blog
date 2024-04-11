package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*
*

  - @Description: Select tags from current page

  - @param pageNum 当前页的startIndex

  - @param pageSize 每页的记录数

  - @param maps 查询条件

  - @return tags 查询到的标签
*/
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	// Limit specify the max number of records to retrieve
	// Offset specify the number of records to skip before starting to return the records
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return tags
}

/*
*

  - @Description: Get the total number of tags

  - @param maps 查询条件

  - @return count 查询到的标签总数
*/
func GetTagTotal(maps interface{}) (count int) {
	// SQL: SELECT count(*) FROM blog_tag WHERE maps
	db.Model(&Tag{}).Where(maps).Count(&count)

	return count
}

func ExistTagByName(name string) bool {
	var tag Tag
	// SELECT id FROM blog_tag WHERE name = ?
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}

// Hook参考：https://gorm.io/zh_CN/docs/hooks.html
