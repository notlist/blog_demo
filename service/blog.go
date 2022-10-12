package service

import (
	"errors"
	"goadmin/common/log"
	"goadmin/dao/blog_dao"
	"goadmin/dao/tag_dao"
	"goadmin/dto"
	"goadmin/entity"
	"time"
)

// TODO 更改根据标签的查询查询
func BlogList(userId int64, req *dto.BlogListReq) ([]dto.BlogListResp, error) {
	blogDao := blog_dao.BlogDaoNew()
	tagDao := tag_dao.TagDaoNew()
	if userId == 0 {
		return nil, errors.New("没有传入uid")
	}
	cond := make(map[string]interface{})
	tagsMap := make(map[int64][]string)
	//获取该用户的所有标签
	if len(req.Tags) > 0 {
		tags, err := tagDao.GetAll(map[string]interface{}{
			"tag_name": req.Tags,
			"user_id":  userId,
		})
		if err != nil {
			log.Logger.Errorf("get tags info err:%+v", err)
			return nil, errors.New("服务器错误")
		}
		tagNames := make([]string, 0)
		for _, v := range tags {
			tagNames = append(tagNames, v.TagName)

			if tagArr, ok := tagsMap[v.BlogId]; ok {
				tagArr = append(tagArr, v.TagName)
				tagsMap[v.BlogId] = tagArr
			} else {
				tagArr = make([]string, 0)
				tagsMap[v.BlogId] = tagArr
			}
		}
		cond["tag"] = tagNames
	}

	blogs, err := blogDao.GetAll(cond)

	if err != nil {
		log.Logger.Errorf("get blogs info err:%+v", err)
		return nil, errors.New("服务器错误")
	}
	blogInfo := make([]dto.BlogListResp, 0)
	for _, v := range blogs {
		temp := dto.BlogListResp{
			BlogId:     v.BlogId,
			Title:      v.Title,
			Content:    v.Content,
			UpdateTime: v.UpdateTime,
			CreateTime: v.CreateTime,
			Tags:       tagsMap[v.BlogId],
		}
		blogInfo = append(blogInfo, temp)
	}
	return blogInfo, nil
}

func CreateBlog(userId int64, req *dto.BlogCreateReq) error {
	blogDao := blog_dao.BlogDaoNew()
	tagDao := tag_dao.TagDaoNew()

	blogId := time.Now().Unix() + userId - 10000000
	blogInfo := &entity.Blog{
		BlogId:     blogId,
		Userid:     userId,
		Title:      req.Title,
		Content:    req.Content,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}
	//博客写入数据库
	blogDao.Add(blogInfo)
	tagInfos := make([]*entity.Tag, 0)
	for _, v := range req.Tags {
		temp := entity.Tag{
			Userid:     userId,
			TagName:    v,
			BlogId:     blogId,
			UpdateTime: time.Now().Unix(),
			CreateTime: time.Now().Unix(),
		}
		tagInfos = append(tagInfos, &temp)
	}
	//tags写入数据库
	tagDao.BatchAdd(tagInfos)
	return nil
}

func EditBlog(userId int64, req *dto.BlogEditReq) error {

	return nil
}
