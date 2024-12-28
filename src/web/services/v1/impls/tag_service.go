package impls

import (
	"family-web-server/src/log"
	"family-web-server/src/pkg/mysql"
	tagDto "family-web-server/src/web/models/dto/tag"
	tagEntity "family-web-server/src/web/models/eneity/tag"
	tagVo "family-web-server/src/web/models/vo/tag"
	"family-web-server/src/web/services/v1/interfaces"
	"family-web-server/src/web/utils"
)

type TagService struct {
	gorm *mysql.GormDb
	l    *log.ConsoleLogger
}

func (t *TagService) DeleteTag(id int) error {
	tx := t.gorm.GetDb().Where("id = ?", id).Delete(&tagEntity.Tag{})
	if tx.Error != nil {
		t.l.Error("删除标签失败:" + tx.Error.Error())
		return tx.Error
	}
	return nil
}

func (t *TagService) UpdateTag(dto *tagDto.TagDto) error {
	tx := t.gorm.GetDb().Where("id = ?", dto.Id).Updates(&tagEntity.Tag{
		Name: dto.Name,
		Type: dto.Type,
	})
	if tx.Error != nil {
		t.l.Error("更新标签失败:" + tx.Error.Error())
		return tx.Error
	}
	return nil
}

func (t *TagService) AddTag(v *tagDto.TagDto) error {
	tx := t.gorm.GetDb().Create(&tagEntity.Tag{
		Name:        v.Name,
		Description: v.Description,
		Type:        v.Type,
	})
	if tx.Error != nil {
		t.l.Error("添加标签失败:" + tx.Error.Error())
		return tx.Error
	}
	return nil
}

func (t *TagService) GetTagsByType(tagType int) ([]*tagVo.TagVo, error) {
	var tags []*tagEntity.Tag
	var tagVos []*tagVo.TagVo
	tx := t.gorm.GetDb().Raw(`
		SELECT t.*
		FROM tag t 
		LEFT JOIN tag_relation tr ON tr.tag_id = t.id
		WHERE tr.link_type = ?
	`, tagType).Scan(&tags)
	if tx.Error != nil {
		t.l.Error("获取标签失败:" + tx.Error.Error())
		return nil, tx.Error
	}
	err := utils.EntityToVO(&tags, &tagVos)
	if err != nil {
		t.l.Error("转换entity to vo失败:" + err.Error())
		return nil, err
	}
	return tagVos, nil
}

func NewTagService(gorm *mysql.GormDb, l *log.ConsoleLogger) interfaces.ITagService {
	return &TagService{gorm: gorm, l: l}
}
