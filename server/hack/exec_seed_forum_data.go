package main

import (
	"context"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type seedUser struct {
	Id uint64 `orm:"id"`
}

type seedCategory struct {
	Id uint64 `orm:"id"`
}

type seedTopic struct {
	Id uint64 `orm:"id"`
}

func main() {
	ctx := gctx.New()
	var users []seedUser
	if err := g.DB().Model("hg_user").Fields("id").OrderAsc("id").Limit(2).Scan(&users); err != nil {
		fmt.Println("Error: query forum seed users:", err)
		return
	}
	if len(users) == 0 {
		fmt.Println("Error: no user records available for forum seed data")
		return
	}

	categoryID, err := ensureCategory(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	firstUserID := users[0].Id
	secondUserID := firstUserID
	if len(users) > 1 {
		secondUserID = users[1].Id
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		return seedTopics(ctx, tx, categoryID, firstUserID, secondUserID)
	}); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully inserted forum demo data.")
}

func ensureCategory(ctx context.Context) (uint64, error) {
	const name = "产品交流"
	count, err := g.DB().Model("hg_forum_category").Where("name", name).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		var category seedCategory
		if err = g.DB().Model("hg_forum_category").Fields("id").Where("name", name).Scan(&category); err != nil {
			return 0, err
		}
		return category.Id, nil
	}
	id, err := g.DB().Model("hg_forum_category").Data(g.Map{
		"pid": 0, "name": name, "description": "分享产品使用经验与建议", "image": "",
		"sort": 1, "status": 1, "topic_count": 0,
	}).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func seedTopics(ctx context.Context, tx gdb.TX, categoryID, firstUserID, secondUserID uint64) error {
	topicID, err := ensureTopic(ctx, tx, categoryID, firstUserID, "欢迎来到社区论坛", "社区,产品交流,新手", "社区交流区说明", "这里可以交流产品使用心得、功能建议和实际案例。欢迎发布你的第一个主题。")
	if err != nil {
		return err
	}
	firstCommentID, err := ensureComment(ctx, tx, topicID, secondUserID, 0, 0, "欢迎加入社区，期待看到更多经验分享。")
	if err != nil {
		return err
	}
	secondCommentID, err := ensureComment(ctx, tx, topicID, firstUserID, firstCommentID, secondUserID, "感谢支持，也欢迎提出改进建议。")
	if err != nil {
		return err
	}
	if _, err = ensureComment(ctx, tx, topicID, secondUserID, secondCommentID, firstUserID, "已经记录，后续会持续补充使用案例。"); err != nil {
		return err
	}
	if err = ensureInteraction(ctx, tx, topicID, secondUserID, "like"); err != nil {
		return err
	}
	if err = ensureInteraction(ctx, tx, topicID, secondUserID, "favorite"); err != nil {
		return err
	}

	if _, err = ensureTopic(ctx, tx, categoryID, secondUserID, "产品使用技巧征集", "技巧,效率,产品", "征集提升工作效率的产品使用技巧", "你最常用的功能是什么？请分享提高工作效率的小技巧。"); err != nil {
		return err
	}
	if _, err = ensureTopic(ctx, tx, categoryID, firstUserID, "社区功能建议收集", "建议,社区,功能", "收集社区产品功能建议", "欢迎提交你希望在社区中看到的功能，并说明实际使用场景。"); err != nil {
		return err
	}
	return refreshForumCounts(ctx, tx, categoryID)
}

func ensureTopic(ctx context.Context, tx gdb.TX, categoryID, userID uint64, title, keywords, description, content string) (uint64, error) {
	count, err := tx.Model("hg_forum_topic").Where("title", title).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		var topic seedTopic
		if err = tx.Model("hg_forum_topic").Fields("id").Where("title", title).Scan(&topic); err != nil {
			return 0, err
		}
		return topic.Id, nil
	}
	id, err := tx.Model("hg_forum_topic").Data(g.Map{
		"category_id": categoryID, "user_id": userID, "title": title, "keywords": keywords, "description": description,
		"content": content, "images": "",
		"status": "normal", "is_top": 0, "is_essence": 0,
	}).InsertAndGetId()
	return uint64(id), err
}

func ensureComment(ctx context.Context, tx gdb.TX, topicID, userID, parentID, replyUserID uint64, content string) (uint64, error) {
	count, err := tx.Model("hg_forum_comment").Where("topic_id", topicID).Where("content", content).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		var comment seedTopic
		if err = tx.Model("hg_forum_comment").Fields("id").Where("topic_id", topicID).Where("content", content).Scan(&comment); err != nil {
			return 0, err
		}
		return comment.Id, nil
	}
	id, err := tx.Model("hg_forum_comment").Data(g.Map{
		"topic_id": topicID, "user_id": userID, "parent_id": parentID, "reply_user_id": replyUserID,
		"content": content, "status": "normal",
	}).InsertAndGetId()
	return uint64(id), err
}

func ensureInteraction(ctx context.Context, tx gdb.TX, topicID, userID uint64, actionType string) error {
	count, err := tx.Model("hg_forum_interaction").
		Where("target_type", "topic").
		Where("target_id", topicID).
		Where("user_id", userID).
		Where("action_type", actionType).
		Count()
	if err != nil || count > 0 {
		return err
	}
	_, err = tx.Model("hg_forum_interaction").Data(g.Map{
		"target_type": "topic", "target_id": topicID, "user_id": userID, "action_type": actionType,
	}).Insert()
	return err
}

func refreshForumCounts(ctx context.Context, tx gdb.TX, categoryID uint64) error {
	var topics []seedTopic
	if err := tx.Model("hg_forum_topic").Fields("id").Where("category_id", categoryID).Where("status", "normal").Scan(&topics); err != nil {
		return err
	}
	if _, err := tx.Model("hg_forum_category").WherePri(categoryID).Data(g.Map{"topic_count": len(topics)}).Update(); err != nil {
		return err
	}
	for _, topic := range topics {
		comments, err := tx.Model("hg_forum_comment").Where("topic_id", topic.Id).Where("status", "normal").Count()
		if err != nil {
			return err
		}
		likes, err := tx.Model("hg_forum_interaction").Where("target_type", "topic").Where("target_id", topic.Id).Where("action_type", "like").Count()
		if err != nil {
			return err
		}
		favorites, err := tx.Model("hg_forum_interaction").Where("target_type", "topic").Where("target_id", topic.Id).Where("action_type", "favorite").Count()
		if err != nil {
			return err
		}
		if _, err = tx.Model("hg_forum_topic").WherePri(topic.Id).Data(g.Map{"comment_count": comments, "like_count": likes, "favorite_count": favorites}).Update(); err != nil {
			return err
		}
	}
	return nil
}
