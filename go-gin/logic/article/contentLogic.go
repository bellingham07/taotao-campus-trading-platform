package articleLogic

import (
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	articleRepository "com.xpdj/go-gin/repository/article"
	"github.com/gin-gonic/gin"
	"time"
)

var ContentLogic = new(ArticleContentLogic)

type ArticleContentLogic struct {
}

func (cl *ArticleContentLogic) SavaOrPublish(contentDraft *model.ArticleContent, userId int64, isPublish bool) gin.H {
	articleContent := cl.copyDraftAttribute(contentDraft)
	articleContent.UserId = userId
	articleContent.CreateAt = time.Now()
	// 保存草稿
	if !isPublish {
		err := articleRepository.ArticleContent.Insert(articleContent)
		if err != nil {
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.OkData(gin.H{"id": articleContent.Id})
	}
	// 发布
	articleContent.Status = 2
	err := articleRepository.ArticleContent.Insert(articleContent)
	if err != nil {
		return response.ErrorMsg("操作失败，请重试！")
	}
	return response.OkData(gin.H{"id": articleContent.Id})
}

func (cl *ArticleContentLogic) Update(content *model.ArticleContent, isPublish bool) gin.H {
	articleContent := cl.copyDraftAttribute(content)
	if !isPublish {
		err := articleRepository.ArticleContent.UpdateById(articleContent)
		if err != nil {
			return response.ErrorMsg("操作失败，请重试！")
		}
		return response.Ok()
	}
	articleContent.Status = 2
	err := articleRepository.ArticleContent.UpdateById(articleContent)
	if err != nil {
		return response.ErrorMsg("操作失败，请重试！")
	}
	return response.Ok()
}

func (*ArticleContentLogic) copyDraftAttribute(draft *model.ArticleContent) *model.ArticleContent {
	articleContent := &model.ArticleContent{
		Title:    draft.Title,
		Content:  draft.Content,
		UpdateAt: time.Now(),
	}
	return articleContent
}
