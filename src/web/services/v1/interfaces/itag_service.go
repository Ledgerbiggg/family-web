package interfaces

import (
	tagDto "family-web-server/src/web/models/dto/tag"
	tagVo "family-web-server/src/web/models/vo/tag"
)

type ITagService interface {
	//
	// GetTagsByType
	//  @Description: 获取指定类型的标签
	//  @param tagType int
	//  @return []*tag.Tag 标签合集
	//  @return error 错误
	//
	GetTagsByType(tagType int) ([]*tagVo.TagVo, error)

	//
	// AddTag
	//  @Description:  添加标签
	//  @param v *tagDto.TagDto
	//  @return error 错误
	//
	AddTag(v *tagDto.TagDto) error

	//
	// UpdateTag
	//  @Description: 更新标签信息
	//  @param dto *tagDto.TagDto
	//  @return error 错误
	//
	UpdateTag(dto *tagDto.TagDto) error

	//
	// DeleteTag
	//  @Description:  删除标签
	//  @param id int
	//  @return error 错误
	//
	DeleteTag(id int) error
}
