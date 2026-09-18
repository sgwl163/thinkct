package site

import (
	"context"
	"fmt"
	"strings"

	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sitein"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sApiSite struct{}

func NewApiSite() *sApiSite {
	return &sApiSite{}
}

func init() {
	service.RegisterApiSite(NewApiSite())
}

// optimizeImage 提取图片显示优化逻辑
func optimizeImage(image, images string) string {
	if image == "" && images != "" {
		imagesArr := strings.Split(images, ",")
		if len(imagesArr) > 0 {
			return imagesArr[0]
		}
	}
	return image
}

// formatChannelUrl 格式化栏目URL
func formatChannelUrl(ch *entity.WebChannel) {
	if ch != nil && ch.Url == "" && ch.Diyname != "" {
		ch.Url = "/" + ch.Diyname
	}
}

// formatArticleUrl 格式化文章URL
func formatArticleUrl(article *entity.WebArticle, ch *entity.WebChannel) {
	if article != nil && article.Url == "" && ch != nil && ch.Diyname != "" {
		article.Url = "/" + ch.Diyname + "/detail/" + gconv.String(article.Id)
	}
}

// formatProductUrl 格式化产品URL
func formatProductUrl(product *entity.WebProducts, ch *entity.WebChannel) {
	if product != nil && product.Url == "" && ch != nil && ch.Diyname != "" {
		product.Url = "/" + ch.Diyname + "/detail/" + gconv.String(product.Id)
	}
}

// getChannelChildrenIds 递归获取频道的子频道ID
func getChannelChildrenIds(ctx context.Context, id uint) []uint {
	var ids []uint
	var children []*entity.WebChannel
	_ = dao.WebChannel.Ctx(ctx).Where("pid", id).Where("status", 1).Scan(&children)
	for _, child := range children {
		ids = append(ids, child.Id)
		ids = append(ids, getChannelChildrenIds(ctx, child.Id)...)
	}
	return ids
}

// Index 获取首页数据
func (s *sApiSite) Index(ctx context.Context, in sitein.IndexInp) (res *sitein.IndexModel, err error) {
	res = new(sitein.IndexModel)

	// 1. 获取网站基础配置与站点设置
	basicConfig, _ := service.SysConfig().GetBasic(ctx)
	siteConfig, _ := service.SysConfig().GetSite(ctx)

	siteMap := make(map[string]interface{})

	if basicConfig != nil {
		siteMap["sitelogo"] = basicConfig.Logo
		siteMap["siteSystemOpen"] = basicConfig.SystemOpen
		siteMap["siteCloseText"] = basicConfig.CloseText
	}

	if siteConfig != nil {
		siteMap["siteName"] = siteConfig.SiteName
		siteMap["siteTitle"] = siteConfig.SiteTitle
		siteMap["siteKeywords"] = siteConfig.SiteKeywords
		siteMap["siteDescription"] = siteConfig.SiteDescription
		siteMap["siteCopyRight"] = siteConfig.SiteCopyRight
		siteMap["siteIcp"] = siteConfig.SiteIcp
		siteMap["siteemail"] = siteConfig.Email
		siteMap["sitephone"] = siteConfig.Phone
		siteMap["sitewechatQrCode"] = siteConfig.WechatQrCode
		siteMap["sitempQrCode"] = siteConfig.MpQrCode
	}

	res.SiteConfig = siteMap

	return res, nil
}

// Home 获取首页各个栏目的聚合数据
func (s *sApiSite) Home(ctx context.Context, in sitein.HomeInp) (res *sitein.HomeModel, err error) {
	res = new(sitein.HomeModel)

	// 处理参数
	productsChannelIdsArr := []int{7, 8, 9, 10, 11} // 默认产品分类
	if in.ProductsChannelIds != "" {
		idStrings := gstr.SplitAndTrim(in.ProductsChannelIds, ",")
		if len(idStrings) > 0 {
			productsChannelIdsArr = make([]int, 0)
			for _, idStr := range idStrings {
				if gstr.IsNumeric(idStr) {
					productsChannelIdsArr = append(productsChannelIdsArr, gconv.Int(idStr))
				} else {
					var chId int
					_ = dao.WebChannel.Ctx(ctx).Where("diyname", idStr).Where("status", 1).Fields("id").Scan(&chId)
					if chId > 0 {
						productsChannelIdsArr = append(productsChannelIdsArr, chId)
					}
				}
			}
		}
	}
	productsLimit := 4 // 默认显示4条
	if in.ProductsLimit > 0 {
		productsLimit = in.ProductsLimit
	}

	casesChannelIdsArr := []int{16} // 默认应用案例/经典案例栏目ID为16 (diyname: cases)
	if in.CasesChannelIds != "" {
		idStrings := gstr.SplitAndTrim(in.CasesChannelIds, ",")
		if len(idStrings) > 0 {
			casesChannelIdsArr = make([]int, 0)
			for _, idStr := range idStrings {
				if gstr.IsNumeric(idStr) {
					casesChannelIdsArr = append(casesChannelIdsArr, gconv.Int(idStr))
				} else {
					var chId int
					_ = dao.WebChannel.Ctx(ctx).Where("diyname", idStr).Where("status", 1).Fields("id").Scan(&chId)
					if chId > 0 {
						casesChannelIdsArr = append(casesChannelIdsArr, chId)
					}
				}
			}
		}
	}
	casesLimit := 3
	if in.CasesLimit > 0 {
		casesLimit = in.CasesLimit
	}

	newsChannelIdsArr := []int{2} // 默认新闻中心根栏目ID为2
	if in.NewsChannelIds != "" {
		idStrings := gstr.SplitAndTrim(in.NewsChannelIds, ",")
		if len(idStrings) > 0 {
			newsChannelIdsArr = make([]int, 0)
			for _, idStr := range idStrings {
				if gstr.IsNumeric(idStr) {
					newsChannelIdsArr = append(newsChannelIdsArr, gconv.Int(idStr))
				} else {
					var chId int
					_ = dao.WebChannel.Ctx(ctx).Where("diyname", idStr).Where("status", 1).Fields("id").Scan(&chId)
					if chId > 0 {
						newsChannelIdsArr = append(newsChannelIdsArr, chId)
					}
				}
			}
		}
	}
	newsLimit := 3
	if in.NewsLimit > 0 {
		newsLimit = in.NewsLimit
	}

	// 1. 获取产品分类 (获取特定ID的栏目作为产品分类显示)
	var productsChannels []*entity.WebChannel
	err = dao.WebChannel.Ctx(ctx).Where("status", 1).WhereIn("id", productsChannelIdsArr).Order("weigh desc, id asc").Scan(&productsChannels)
	if err == nil && len(productsChannels) > 0 {
		res.ProductsChannels = make([]*sitein.ProductChannelModel, 0)
		for _, ch := range productsChannels {
			formatChannelUrl(ch)
			model := &sitein.ProductChannelModel{
				WebChannel: ch,
				List:       make([]*entity.WebProducts, 0),
			}
			// 获取频道及所有子频道ID
			channelIds := append(getChannelChildrenIds(ctx, ch.Id), ch.Id)

			// 建立 channelId -> channel 映射
			channelMap := make(map[uint]*entity.WebChannel)
			channelMap[ch.Id] = ch
			if len(channelIds) > 1 {
				var subChannels []*entity.WebChannel
				_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&subChannels)
				for _, subCh := range subChannels {
					formatChannelUrl(subCh)
					channelMap[subCh.Id] = subCh
				}
			}

			// 查询该栏目及子栏目下的产品内容
			var chProducts []*entity.WebProducts
			err = dao.WebProducts.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).Order("weigh desc, id desc").Limit(productsLimit).Scan(&chProducts)
			if err == nil && chProducts != nil {
				for _, p := range chProducts {
					p.Image = optimizeImage(p.Image, p.Images)
					itemChannel := channelMap[p.ChannelId]
					if itemChannel == nil {
						itemChannel = ch
					}
					formatProductUrl(p, itemChannel)
				}
				model.List = chProducts
			}
			res.ProductsChannels = append(res.ProductsChannels, model)
		}
	} else {
		res.ProductsChannels = make([]*sitein.ProductChannelModel, 0)
	}

	// 2. 获取经典案例 (查询特定栏目的文章)
	var casesChannels []*entity.WebChannel
	err = dao.WebChannel.Ctx(ctx).Where("status", 1).WhereIn("id", casesChannelIdsArr).Order("weigh desc, id asc").Scan(&casesChannels)
	if err == nil && len(casesChannels) > 0 {
		res.CasesChannels = make([]*sitein.ArticleChannelModel, 0)
		for _, ch := range casesChannels {
			formatChannelUrl(ch)
			model := &sitein.ArticleChannelModel{
				WebChannel: ch,
				List:       make([]*entity.WebArticle, 0),
			}
			// 获取频道及所有子频道ID
			channelIds := append(getChannelChildrenIds(ctx, ch.Id), ch.Id)

			channelMap := make(map[uint]*entity.WebChannel)
			channelMap[ch.Id] = ch
			if len(channelIds) > 1 {
				var subChannels []*entity.WebChannel
				_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&subChannels)
				for _, subCh := range subChannels {
					formatChannelUrl(subCh)
					channelMap[subCh.Id] = subCh
				}
			}

			var chCases []*entity.WebArticle
			// 经典案例一般带有图片 (数据库字段为 images)
			err = dao.WebArticle.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).WhereNot("images", "").Order("weigh desc, id desc").Limit(casesLimit).Scan(&chCases)
			if err == nil && chCases != nil {
				for _, item := range chCases {
					item.Image = optimizeImage(item.Image, item.Images)
					itemChannel := channelMap[item.ChannelId]
					if itemChannel == nil {
						itemChannel = ch
					}
					formatArticleUrl(item, itemChannel)
				}
				model.List = chCases
			}
			res.CasesChannels = append(res.CasesChannels, model)
		}
	} else {
		res.CasesChannels = make([]*sitein.ArticleChannelModel, 0)
	}

	// 3. 获取新闻中心 (查询特定栏目的文章)
	var newsChannels []*entity.WebChannel
	err = dao.WebChannel.Ctx(ctx).Where("status", 1).WhereIn("id", newsChannelIdsArr).Order("weigh desc, id asc").Scan(&newsChannels)
	if err == nil && len(newsChannels) > 0 {
		res.NewsChannels = make([]*sitein.ArticleChannelModel, 0)
		for _, ch := range newsChannels {
			formatChannelUrl(ch)
			model := &sitein.ArticleChannelModel{
				WebChannel: ch,
				List:       make([]*entity.WebArticle, 0),
			}
			// 获取频道及所有子频道ID
			channelIds := append(getChannelChildrenIds(ctx, ch.Id), ch.Id)

			channelMap := make(map[uint]*entity.WebChannel)
			channelMap[ch.Id] = ch
			if len(channelIds) > 1 {
				var subChannels []*entity.WebChannel
				_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&subChannels)
				for _, subCh := range subChannels {
					formatChannelUrl(subCh)
					channelMap[subCh.Id] = subCh
				}
			}

			var chNews []*entity.WebArticle
			err = dao.WebArticle.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).Order("publish_time desc, id desc").Limit(newsLimit).Scan(&chNews)
			if err == nil && chNews != nil {
				for _, item := range chNews {
					item.Image = optimizeImage(item.Image, item.Images)
					itemChannel := channelMap[item.ChannelId]
					if itemChannel == nil {
						itemChannel = ch
					}
					formatArticleUrl(item, itemChannel)
				}
				model.List = chNews
			}
			res.NewsChannels = append(res.NewsChannels, model)
		}
	} else {
		res.NewsChannels = make([]*sitein.ArticleChannelModel, 0)
	}

	// 4. 获取友情链接 (hg_web_links，状态为 enable，按 weigh 倒序、id 正序)
	var links []*entity.WebLinks
	err = dao.WebLinks.Ctx(ctx).Where("status", "enable").Order("weigh desc, id asc").Scan(&links)
	if err == nil && links != nil {
		res.Links = links
	} else {
		res.Links = make([]*entity.WebLinks, 0)
	}

	return res, nil
}

// handleList 提取列表查询逻辑
func (s *sApiSite) handleList(ctx context.Context, m *gdb.Model, in sitein.ChannelInp, list interface{}) (totalCount int, err error) {
	totalCount, err = m.Count()
	if err == nil && totalCount > 0 {
		err = m.Page(in.Page, in.Limit).Order("weigh desc, id desc").Scan(list)
	}
	return
}

// handleHotList 提取热门列表查询逻辑
func (s *sApiSite) handleHotList(ctx context.Context, m *gdb.Model, hotList interface{}) (err error) {
	err = m.Order("views desc, id desc").Limit(5).Scan(hotList)
	return
}

// handlePrevNext 提取上一篇/下一篇查询逻辑
func (s *sApiSite) handlePrevNext(ctx context.Context, m *gdb.Model, channelId int64, currentId int64, prev interface{}, next interface{}) {
	_ = m.Clone().Where("status", "normal").Where("channel_id", channelId).Where("id <", currentId).Order("id desc").Scan(prev)
	_ = m.Clone().Where("status", "normal").Where("channel_id", channelId).Where("id >", currentId).Order("id asc").Scan(next)
}

// Channel 获取前台栏目列表与详情
func (s *sApiSite) Channel(ctx context.Context, in sitein.ChannelInp) (res *sitein.ChannelModel, err error) {
	res = new(sitein.ChannelModel)

	// 如果传入了特定栏目ID或自定义名称，则获取该栏目信息及其关联模型的数据列表
	if in.Channel != "" {
		mChannel := dao.WebChannel.Ctx(ctx).Where("status", 1)
		// 如果全为数字，优先视为ID查询，否则视为自定义名称(diyname)查询
		if gstr.IsNumeric(in.Channel) {
			mChannel = mChannel.Where("id", in.Channel)
		} else {
			mChannel = mChannel.Where("diyname", in.Channel)
		}
		err = mChannel.Scan(&res.Channel)
		if err != nil {
			return nil, err
		}

		// 若该栏目存在，根据它的模型(`model`)查询对应的列表
		if res.Channel != nil {
			formatChannelUrl(res.Channel)

			// 获取面包屑导航 (从根栏目到当前栏目)
			breadcrumb := make([]*entity.WebChannel, 0)
			curr := res.Channel
			for curr != nil {
				breadcrumb = append([]*entity.WebChannel{curr}, breadcrumb...)
				if curr.Pid > 0 {
					var parent *entity.WebChannel
					_ = dao.WebChannel.Ctx(ctx).Where("id", curr.Pid).Where("status", 1).Scan(&parent)
					if parent != nil {
						formatChannelUrl(parent)
					}
					curr = parent
				} else {
					break
				}
			}
			res.Breadcrumb = breadcrumb

			// 获取同级相关分类
			dao.WebChannel.Ctx(ctx).Where("pid", res.Channel.Pid).Where("status", 1).Order("weigh desc, id asc").Scan(&res.RelatedList)
			if res.RelatedList == nil {
				res.RelatedList = make([]*entity.WebChannel, 0)
			} else {
				for _, relCh := range res.RelatedList {
					formatChannelUrl(relCh)
				}
			}
			if in.Page <= 0 {
				in.Page = 1
			}
			if in.Limit <= 0 {
				in.Limit = 10
			}
			res.Page = in.Page

			// 获取频道及所有子频道ID
			channelIds := append(getChannelChildrenIds(ctx, res.Channel.Id), res.Channel.Id)

			// 建立 channelId -> channel 映射，方便生成子频道内容的准确 URL
			channelMap := make(map[uint]*entity.WebChannel)
			channelMap[res.Channel.Id] = res.Channel
			if len(channelIds) > 1 {
				var subChannels []*entity.WebChannel
				_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&subChannels)
				for _, ch := range subChannels {
					formatChannelUrl(ch)
					channelMap[ch.Id] = ch
				}
			}

			if res.Channel.Model == "article" {
				m := dao.WebArticle.Ctx(ctx).Where("status", 1).WhereIn("channel_id", channelIds)
				var list []*entity.WebArticle
				var lst interface{} = &list
				res.Total, err = s.handleList(ctx, m, in, lst)
				if list == nil {
					list = make([]*entity.WebArticle, 0)
				}
				for _, article := range list {
					article.Image = optimizeImage(article.Image, article.Images)
					itemChannel := channelMap[article.ChannelId]
					if itemChannel == nil {
						itemChannel = res.Channel
					}
					formatArticleUrl(article, itemChannel)
				}
				res.List = list

				// 获取热门文章 (根据 views 排序)
				mHot := dao.WebArticle.Ctx(ctx).Where("status", 1).WhereIn("channel_id", channelIds)
				var hotList []*entity.WebArticle
				var hotLst interface{} = &hotList
				err = s.handleHotList(ctx, mHot, hotLst)
				if hotList == nil {
					hotList = make([]*entity.WebArticle, 0)
				}
				for _, article := range hotList {
					article.Image = optimizeImage(article.Image, article.Images)
					itemChannel := channelMap[article.ChannelId]
					if itemChannel == nil {
						itemChannel = res.Channel
					}
					formatArticleUrl(article, itemChannel)
				}
				res.HotList = hotList
			} else if res.Channel.Model == "products" {
				m := dao.WebProducts.Ctx(ctx).Where("status", 1).WhereIn("channel_id", channelIds)
				var list []*entity.WebProducts
				var lst interface{} = &list
				res.Total, err = s.handleList(ctx, m, in, lst)
				if list == nil {
					list = make([]*entity.WebProducts, 0)
				}
				for _, product := range list {
					product.Image = optimizeImage(product.Image, product.Images)
					itemChannel := channelMap[product.ChannelId]
					if itemChannel == nil {
						itemChannel = res.Channel
					}
					formatProductUrl(product, itemChannel)
				}
				res.List = list

				// 获取热门产品 (根据 views 排序)
				mHot := dao.WebProducts.Ctx(ctx).Where("status", 1).WhereIn("channel_id", channelIds)
				var hotList []*entity.WebProducts
				var hotLst interface{} = &hotList
				err = s.handleHotList(ctx, mHot, hotLst)
				if hotList == nil {
					hotList = make([]*entity.WebProducts, 0)
				}
				for _, product := range hotList {
					product.Image = optimizeImage(product.Image, product.Images)
					itemChannel := channelMap[product.ChannelId]
					if itemChannel == nil {
						itemChannel = res.Channel
					}
					formatProductUrl(product, itemChannel)
				}
				res.HotList = hotList
			}
		}
	} else {
		// 没有传入栏目时，默认获取所有栏目列表(存入List以便导航使用)
		var list []*entity.WebChannel
		err = dao.WebChannel.Ctx(ctx).Where("status", 1).Order("weigh desc, id asc").Scan(&list)
		if err != nil {
			return nil, err
		}
		if list == nil {
			list = make([]*entity.WebChannel, 0)
		}
		for _, ch := range list {
			formatChannelUrl(ch)
		}
		res.List = list
	}

	return res, nil
}

// Content 获取内容详情
func (s *sApiSite) Content(ctx context.Context, in sitein.ContentInp) (res *sitein.ContentModel, err error) {
	res = new(sitein.ContentModel)

	if in.Channel == "" {
		return nil, gerror.New("所属栏目不能为空")
	}

	// 先获取栏目信息
	mChannel := dao.WebChannel.Ctx(ctx).Where("status", 1)
	if gstr.IsNumeric(in.Channel) {
		mChannel = mChannel.Where("id", in.Channel)
	} else {
		mChannel = mChannel.Where("diyname", in.Channel)
	}
	err = mChannel.Scan(&res.Channel)
	if err != nil {
		return nil, err
	}
	if res.Channel == nil {
		return nil, gerror.New("所属栏目不存在或已禁用")
	}
	formatChannelUrl(res.Channel)

	// 获取面包屑导航 (从根栏目到当前栏目)
	breadcrumb := make([]*entity.WebChannel, 0)
	curr := res.Channel
	for curr != nil {
		breadcrumb = append([]*entity.WebChannel{curr}, breadcrumb...)
		if curr.Pid > 0 {
			var parent *entity.WebChannel
			_ = dao.WebChannel.Ctx(ctx).Where("id", curr.Pid).Where("status", 1).Scan(&parent)
			if parent != nil {
				formatChannelUrl(parent)
			}
			curr = parent
		} else {
			break
		}
	}
	res.Breadcrumb = breadcrumb

	if res.Channel.Model == "article" {
		var article *entity.WebArticle
		err = dao.WebArticle.Ctx(ctx).Where("id", in.Id).Where("status", 1).Scan(&article)
		if err != nil {
			return nil, err
		}
		if article != nil {
			// 优化 Image 字段：如果 article.Image 为空，但 Images 不为空，则取 Images 的第一张图
			article.Image = optimizeImage(article.Image, article.Images)
			formatArticleUrl(article, res.Channel)
			res.Content = article

			var prevItem *entity.WebArticle
			var nextItem *entity.WebArticle
			s.handlePrevNext(ctx, dao.WebArticle.Ctx(ctx), int64(res.Channel.Id), int64(in.Id), &prevItem, &nextItem)
			if prevItem != nil {
				prevItem.Image = optimizeImage(prevItem.Image, prevItem.Images)
				formatArticleUrl(prevItem, res.Channel)
				res.Prev = prevItem
			}
			if nextItem != nil {
				nextItem.Image = optimizeImage(nextItem.Image, nextItem.Images)
				formatArticleUrl(nextItem, res.Channel)
				res.Next = nextItem
			}

			// 每次访问增加一次浏览量
			dao.WebArticle.Ctx(ctx).Where("id", article.Id).Increment("views", 1)
		} else {
			return nil, gerror.New("内容不存在或已下架")
		}

		// 获取同级分类
		dao.WebChannel.Ctx(ctx).Where("pid", res.Channel.Pid).Where("status", 1).Order("weigh desc, id asc").Scan(&res.RelatedList)
		if res.RelatedList == nil {
			res.RelatedList = make([]*entity.WebChannel, 0)
		} else {
			for _, relCh := range res.RelatedList {
				formatChannelUrl(relCh)
			}
		}

		// 获取热门内容
		mHot := dao.WebArticle.Ctx(ctx).Where("status", 1).Where("channel_id", res.Channel.Id)
		var hotList []*entity.WebArticle
		var hotLst interface{} = &hotList
		_ = s.handleHotList(ctx, mHot, hotLst)
		if hotList == nil {
			hotList = make([]*entity.WebArticle, 0)
		}
		for _, v := range hotList {
			v.Image = optimizeImage(v.Image, v.Images)
			formatArticleUrl(v, res.Channel)
		}
		res.HotList = hotList

	} else if res.Channel.Model == "products" {
		var product *entity.WebProducts
		err = dao.WebProducts.Ctx(ctx).Where("id", in.Id).Where("status", 1).Scan(&product)
		if err != nil {
			return nil, err
		}
		if product != nil {
			// 优化 Image 字段：如果 product.Image 为空，但 Images 不为空，则取 Images 的第一张图
			product.Image = optimizeImage(product.Image, product.Images)
			formatProductUrl(product, res.Channel)
			res.Content = product

			var prevItem *entity.WebProducts
			var nextItem *entity.WebProducts
			s.handlePrevNext(ctx, dao.WebProducts.Ctx(ctx), int64(res.Channel.Id), int64(in.Id), &prevItem, &nextItem)
			if prevItem != nil {
				prevItem.Image = optimizeImage(prevItem.Image, prevItem.Images)
				formatProductUrl(prevItem, res.Channel)
				res.Prev = prevItem
			}
			if nextItem != nil {
				nextItem.Image = optimizeImage(nextItem.Image, nextItem.Images)
				formatProductUrl(nextItem, res.Channel)
				res.Next = nextItem
			}

			// 每次访问增加一次浏览量
			dao.WebProducts.Ctx(ctx).Where("id", product.Id).Increment("views", 1)
		} else {
			return nil, gerror.New("内容不存在或已下架")
		}

		// 获取同级分类
		dao.WebChannel.Ctx(ctx).Where("pid", res.Channel.Pid).Where("status", 1).Order("weigh desc, id asc").Scan(&res.RelatedList)
		if res.RelatedList == nil {
			res.RelatedList = make([]*entity.WebChannel, 0)
		} else {
			for _, relCh := range res.RelatedList {
				formatChannelUrl(relCh)
			}
		}

		// 获取热门内容
		mHot := dao.WebProducts.Ctx(ctx).Where("status", 1).Where("channel_id", res.Channel.Id)
		var hotList []*entity.WebProducts
		var hotLst interface{} = &hotList
		_ = s.handleHotList(ctx, mHot, hotLst)
		if hotList == nil {
			hotList = make([]*entity.WebProducts, 0)
		}
		for _, v := range hotList {
			v.Image = optimizeImage(v.Image, v.Images)
			formatProductUrl(v, res.Channel)
		}
		res.HotList = hotList

	} else {
		return nil, gerror.New("当前栏目不支持获取详情")
	}

	return res, nil
}

// Search 全站内容搜索
func (s *sApiSite) Search(ctx context.Context, in sitein.SearchInp) (res *sitein.SearchModel, err error) {
	keywords := strings.TrimSpace(in.Keywords)
	if keywords == "" {
		return nil, gerror.New("请输入搜索关键词")
	}

	searchType := strings.TrimSpace(in.Type)
	if searchType == "" {
		searchType = "products"
	}

	if searchType != "products" && searchType != "article" && searchType != "news" && searchType != "cases" {
		return nil, gerror.New("非法的搜索类型")
	}

	if searchType == "news" || searchType == "cases" {
		searchType = "article"
	}

	page := in.Page
	if page <= 0 {
		page = 1
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 16
	}

	res = &sitein.SearchModel{
		Page:  page,
		Limit: limit,
	}

	likeKeywords := "%" + keywords + "%"

	if searchType == "products" {
		m := dao.WebProducts.Ctx(ctx).Where("status", 1).Where("(title LIKE ? OR description LIKE ?)", likeKeywords, likeKeywords)
		totalCount, countErr := m.Count()
		if countErr != nil {
			return nil, countErr
		}
		res.Total = totalCount

		var list []*entity.WebProducts
		if totalCount > 0 {
			err = m.Page(page, limit).Order("weigh desc, id desc").Scan(&list)
			if err != nil {
				return nil, err
			}
		}
		if list == nil {
			list = make([]*entity.WebProducts, 0)
		}

		if len(list) > 0 {
			// 收集所有的 channel_id，批量查询栏目以格式化 URL
			channelIdMap := make(map[uint]bool)
			for _, item := range list {
				channelIdMap[item.ChannelId] = true
			}
			var channelIds []uint
			for cid := range channelIdMap {
				channelIds = append(channelIds, cid)
			}
			var channels []*entity.WebChannel
			_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&channels)
			chMap := make(map[uint]*entity.WebChannel)
			for _, ch := range channels {
				formatChannelUrl(ch)
				chMap[ch.Id] = ch
			}

			for _, item := range list {
				item.Image = optimizeImage(item.Image, item.Images)
				formatProductUrl(item, chMap[item.ChannelId])
			}
		}
		res.List = list
	} else if searchType == "article" {
		m := dao.WebArticle.Ctx(ctx).Where("status", 1).Where("(title LIKE ? OR description LIKE ?)", likeKeywords, likeKeywords)
		totalCount, countErr := m.Count()
		if countErr != nil {
			return nil, countErr
		}
		res.Total = totalCount

		var list []*entity.WebArticle
		if totalCount > 0 {
			err = m.Page(page, limit).Order("weigh desc, id desc").Scan(&list)
			if err != nil {
				return nil, err
			}
		}
		if list == nil {
			list = make([]*entity.WebArticle, 0)
		}

		if len(list) > 0 {
			// 收集所有的 channel_id，批量查询栏目以格式化 URL
			channelIdMap := make(map[uint]bool)
			for _, item := range list {
				channelIdMap[item.ChannelId] = true
			}
			var channelIds []uint
			for cid := range channelIdMap {
				channelIds = append(channelIds, cid)
			}
			var channels []*entity.WebChannel
			_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&channels)
			chMap := make(map[uint]*entity.WebChannel)
			for _, ch := range channels {
				formatChannelUrl(ch)
				chMap[ch.Id] = ch
			}

			for _, item := range list {
				item.Image = optimizeImage(item.Image, item.Images)
				formatArticleUrl(item, chMap[item.ChannelId])
			}
		}
		res.List = list
	}

	// 尝试记录搜索词统计 (如果数据库存在 search_log / hg_web_search_log 相关表则自动计数，否则安全忽略)
	go func(kw string) {
		defer func() { _ = recover() }()
		// 容错检查 search_log 表是否存在并记录
		mLog := g.DB().Model("hg_web_search_log")
		record, _ := mLog.Where("search", kw).One()
		if !record.IsEmpty() {
			_, _ = mLog.Where("search", kw).Increment("count", 1)
		} else {
			_, _ = mLog.Insert(g.Map{
				"search":     kw,
				"count":      1,
				"hot":        0,
				"status":     0,
				"created_at": gtime.Now(),
				"updated_at": gtime.Now(),
			})
		}
	}(keywords)

	return res, nil
}

// Feedback 提交在线留言
func (s *sApiSite) Feedback(ctx context.Context, in sitein.FeedbackInp) (res *sitein.FeedbackModel, err error) {
	if in.Name == "" {
		return nil, gerror.New("姓名不能为空")
	}
	if in.Phone == "" {
		return nil, gerror.New("手机号不能为空")
	}
	if in.Content == "" {
		return nil, gerror.New("留言内容不能为空")
	}

	// 整合备注信息（如：留言主题、期望回复时间等扩展信息记录到 answer 备注字段）
	var answerParts []string
	if in.MessageTopic != "" {
		answerParts = append(answerParts, fmt.Sprintf("主题: %s", in.MessageTopic))
	}
	if in.ExpectedResponseTime != "" {
		answerParts = append(answerParts, fmt.Sprintf("期望回复时间: %s", in.ExpectedResponseTime))
	}
	answer := gstr.Join(answerParts, " | ")

	now := gtime.Now().Timestamp()
	data := g.Map{
		dao.WebFeedback.Columns().Name:       in.Name,
		dao.WebFeedback.Columns().Email:      in.Email,
		dao.WebFeedback.Columns().Phone:      in.Phone,
		dao.WebFeedback.Columns().Company:    in.Company,
		dao.WebFeedback.Columns().Content:    in.Content,
		dao.WebFeedback.Columns().Answer:     answer,
		dao.WebFeedback.Columns().Ip:         in.Ip,
		dao.WebFeedback.Columns().Status:     1, // 默认启用/正常
		dao.WebFeedback.Columns().CreateTime: now,
		dao.WebFeedback.Columns().UpdateTime: now,
	}

	if _, err = dao.WebFeedback.Ctx(ctx).Data(data).Insert(); err != nil {
		return nil, gerror.Wrap(err, "提交留言失败，请稍后重试")
	}

	return &sitein.FeedbackModel{}, nil
}
