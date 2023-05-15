package articleLogic

import (
	mqLogic "com.xpdj/go-gin/logic/rabbitmq"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	articleRepository "com.xpdj/go-gin/repository/article"
	"com.xpdj/go-gin/utils/cache"
	"com.xpdj/go-gin/utils/jsonUtil"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"strconv"
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

func (*ArticleContentLogic) GetById(id, userId int64) gin.H {
	idStr := strconv.FormatInt(id, 10)
	key := cache.ArticleContent + idStr
	contentMap := cache.RedisUtil.HGETALL(key)
	// 数据库也没有数据，防止缓存穿透
	if contentMap["id"] == "nil" {
		_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
		return response.ErrorMsg("没有这篇文章！你不要乱来呀😡")
	}
	collectKey := cache.ArticleCollect + idStr
	var collectFlag int8 = 0
	var isMine int8 = 1
	// redis没有数据，就从数据库里查
	if contentMap["id"] == "" {
		content := articleRepository.ArticleContent.QueryById(id)
		// 数据无，设置空
		if content == nil {
			_ = cache.RedisUtil.HSET(key, map[string]string{"Id": "nil"})
			_ = cache.RedisUtil.EXPIRE(key, 30*time.Second)
			return response.ErrorMsg("没有这篇文章！你不要乱来呀😡")
		}
		// 数据有
		go updateArticleView(id)
		// jwt中存在用户，判断是在访问自己的商品还是别人的
		if userId != 0 {
			// 当文章不是自己的时候更新足迹，并且把标识也做好（是否是自己的，是为1，就不需要收藏按钮，否则为0）
			if content.UserId != userId {
				collectFlag, isMine = isCollected(userId, collectKey)
			}
		}
		return response.OkData(gin.H{"articleContent": content, "isCollected": collectFlag, "isMine": isMine})
	}
	// redis有数据
	if userId != 0 {
		if contentMap["userId"] != strconv.FormatInt(userId, 10) {
			collectFlag, isMine = isCollected(userId, collectKey)
			go updateArticleView(id)
		}
	}
	return response.OkData(gin.H{"commodityInfo": contentMap, "collectFlag": collectFlag, "isMine": isMine})
}

func isCollected(userId int64, collectKey string) (a1, a2 int8) {
	// 如果是别人的商品，就需要判断有没有收藏过
	if isMember := cache.RedisUtil.SISMEMBER(collectKey, userId); isMember {
		a1 = 1
	}
	return a1, 0
}

func updateArticleView(id int64) {
	key := cache.ArticleView + strconv.FormatInt(id, 10)
	// 1.如果hash的count字段设置成功，则说明，可以进行更新
	err := cache.RedisUtil.HSETNX(key, "count", 1)
	ticker := time.NewTicker(time.Second * 30)
	// 2.如果set失败，代表有点赞数，加1就好了
	if err != nil {
		_ = cache.RedisUtil.HINCRBY1(key, "count")
	}
	publisher := mqLogic.VPublisher()
	vMessage := &mqLogic.VMessage{
		RedisKey:    key,
		IsCommodity: false,
	}
	// 3.假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(vMessage)
			return
		}
	}
	body, _ := jsonUtil.Json.Marshal(vMessage)
	err = publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			go mqLogic.ViewCheckUpdate(vMessage)
			return
		}
	}
}

func (*ArticleContentLogic) Like(id int64, userId int64) interface{} {
	key := cache.ArticleLike + strconv.FormatInt(id, 10)
	affect := cache.RedisUtil.SADD(key, userId)
	if affect == 0 {
		return response.ErrorMsg("不能重复点赞哦😊")
	}
	go LikeUpdatePublisher(key, userId)
	return response.Ok()
}

func (*ArticleContentLogic) Unlike(id int64, userId int64) interface{} {
	key := cache.ArticleLike + strconv.FormatInt(id, 10)
	isMember := cache.RedisUtil.SISMEMBER(key, userId)
	if !isMember {
		return response.ErrorMsg("你本来就没有点赞🤡")
	}
	// 取消点赞，删去set中的member即可，没必要更改库
	go cache.RedisUtil.SREM(key, userId)
	return response.Ok()
}

func LikeUpdatePublisher(redisKey string, member int64) {
	now := time.Now()
	ticker := time.NewTicker(time.Second * 30)
	message := &mqLogic.LMessage{
		RedisKey:  redisKey,
		UserId:    member,
		Time:      now,
		IsArticle: true,
	}
	body, _ := jsonUtil.Json.Marshal(message)
	publisher := mqLogic.LPublisher()
	// 假如无法使用mq
	if publisher == nil {
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(message)
			return
		}
	}
	err := publisher.Channel.Publish(publisher.Exchange, publisher.Key, false, false,
		amqp.Publishing{DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("[RABBITMQ ERROR] ", err.Error())
		select {
		case <-ticker.C:
			go mqLogic.LikeCheckUpdate(message)
		}
	}
}
