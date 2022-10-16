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
	tagCond := make(map[string]interface{})
	tagsMap := make(map[int64][]string)
	//获取该用户的所有标签
	if len(req.Tags) > 0 {
		tagCond["tag_name"] = req.Tags
	}
	tagCond["user_id"] = userId
	tags, err := tagDao.GetAll(tagCond)
	if err != nil {
		log.Logger.Errorf("get tags info err:%+v", err)
		return nil, errors.New("服务器错误")
	}
	blogIds := make([]int64, 0)
	for _, v := range tags {
		blogIds = append(blogIds, v.BlogId)
		if tagArr, ok := tagsMap[v.BlogId]; ok {
			tagArr = append(tagArr, v.TagName)
			tagsMap[v.BlogId] = tagArr
		} else {
			tagArr = make([]string, 0)
			tagArr = append(tagArr, v.TagName)
			tagsMap[v.BlogId] = tagArr
		}
	}
	blogCond := make(map[string]interface{})
	blogCond["user_id"] = req.UserId
	if len(req.Tags) > 0 {
		blogCond["blog_id"] = blogIds
	}
	//获取博客
	blogs, err := blogDao.GetAll(blogCond)
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
	blogDao := blog_dao.BlogDaoNew()
	tagDao := tag_dao.TagDaoNew()
	cond := map[string]interface{}{
		"blog_id": req.BlogId,
		"user_id": userId,
	}
	err := blogDao.Update(cond, map[string]interface{}{
		"title":       req.Title,
		"content":     req.Content,
		"update_time": time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	err = tagDao.Delete(cond)
	if err != nil {
		return err
	}
	tagInfos := make([]*entity.Tag, 0)
	for _, v := range req.Tags {
		temp := entity.Tag{
			Userid:     userId,
			TagName:    v,
			BlogId:     req.BlogId,
			UpdateTime: time.Now().Unix(),
			CreateTime: time.Now().Unix(),
		}
		tagInfos = append(tagInfos, &temp)
	}
	//tags写入数据库
	err = tagDao.BatchAdd(tagInfos)
	return err
}

func DeleteBlog(userId int64, req *dto.BlogDeleteReq) error {
	blogDao := blog_dao.BlogDaoNew()
	tagDao := tag_dao.TagDaoNew()
	cond := map[string]interface{}{
		"user_id": userId,
		"blog_id": req.BlogId,
	}
	err := blogDao.Delete(cond)
	if err != nil {
		return err
	}
	err = tagDao.Delete(cond)
	return err
}

func BLogDetail(req *dto.BlogDetailReq) (*dto.BlogDetailResp, error) {
	blogDao := blog_dao.BlogDaoNew()
	tagDao := tag_dao.TagDaoNew()

	cond := map[string]interface{}{
		"blog_id": req.BlogId,
	}
	res, err := blogDao.GetOne(cond)
	if err != nil {
		log.Logger.Errorf("get blog data err: %+v", err)
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	data := &dto.BlogDetailResp{
		BlogId:     res.BlogId,
		Title:      res.Title,
		Content:    res.Content,
		UpdateTime: res.UpdateTime,
		CreateTime: res.CreateTime,
	}

	tagsRes, err := tagDao.GetAll(cond)
	if err != nil {
		log.Logger.Errorf("get tags data err: %+v", err)
		return nil, err
	}
	if tagsRes == nil {
		return data, nil
	}
	tags := make([]string, 0)
	for _, v := range tagsRes {
		tags = append(tags, v.TagName)
	}

	data.Tags = tags

	return data, nil
}
