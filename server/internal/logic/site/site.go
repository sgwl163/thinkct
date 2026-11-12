package site

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"hotgo/internal/consts"
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

// normalizeImages converts legacy comma-separated values and JSON arrays into image URLs.
func normalizeImages(images string) []string {
	images = strings.TrimSpace(images)
	if images == "" {
		return []string{}
	}

	var imageList []string
	if strings.HasPrefix(images, "[") && json.Unmarshal([]byte(images), &imageList) == nil {
		result := make([]string, 0, len(imageList))
		for _, image := range imageList {
			if image = strings.TrimSpace(image); image != "" {
				result = append(result, image)
			}
		}
		return result
	}

	imageList = strings.Split(images, ",")
	result := make([]string, 0, len(imageList))
	for _, image := range imageList {
		if image = strings.TrimSpace(image); image != "" {
			result = append(result, image)
		}
	}
	return result
}

// optimizeImage returns the configured cover image or the first image in the gallery.
func optimizeImage(image, images string) string {
	if imageList := normalizeImages(image); len(imageList) > 0 {
		return imageList[0]
	}
	if imageList := normalizeImages(images); len(imageList) > 0 {
		return imageList[0]
	}
	return ""
}

// normalizeImageFields standardizes image fields for the site API response.
func normalizeImageFields(image *string, images *string) {
	imageList := normalizeImages(*images)
	*images = strings.Join(imageList, ",")
	*image = optimizeImage(*image, *images)
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
		article.Url = "/" + ch.Diyname + "/" + gconv.String(article.Id)
	}
}

// formatProductUrl 格式化产品URL
func formatProductUrl(product *entity.ShopProduct, ch *entity.WebChannel) {
	if product != nil && product.Url == "" && ch != nil && ch.Diyname != "" {
		product.Url = "/" + ch.Diyname + "/" + gconv.String(product.Id)
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

// resolveChannelIds resolves channel IDs or diyname values, falling back to defaults.
func resolveChannelIds(ctx context.Context, value string, defaults []int) []int {
	if value == "" {
		return defaults
	}

	idStrings := gstr.SplitAndTrim(value, ",")
	if len(idStrings) == 0 {
		return defaults
	}

	ids := make([]int, 0, len(idStrings))
	for _, idStr := range idStrings {
		if gstr.IsNumeric(idStr) {
			ids = append(ids, gconv.Int(idStr))
			continue
		}

		var channelId int
		_ = dao.WebChannel.Ctx(ctx).Where("diyname", idStr).Where("status", 1).Fields("id").Scan(&channelId)
		if channelId > 0 {
			ids = append(ids, channelId)
		}
	}
	return ids
}

// getChannelIdsAndMap returns a channel and all enabled descendants for content queries.
func getChannelIdsAndMap(ctx context.Context, channel *entity.WebChannel) ([]uint, map[uint]*entity.WebChannel) {
	channelIds := append(getChannelChildrenIds(ctx, channel.Id), channel.Id)
	channelMap := map[uint]*entity.WebChannel{channel.Id: channel}
	if len(channelIds) == 1 {
		return channelIds, channelMap
	}

	var subChannels []*entity.WebChannel
	_ = dao.WebChannel.Ctx(ctx).WhereIn("id", channelIds).Where("status", 1).Scan(&subChannels)
	for _, subChannel := range subChannels {
		formatChannelUrl(subChannel)
		channelMap[subChannel.Id] = subChannel
	}
	return channelIds, channelMap
}

// filterContentByChannelIds matches content assigned to a primary or secondary channel.
func filterContentByChannelIds(m *gdb.Model, channelIdColumn, channelIdsColumn string, channelIds []uint) *gdb.Model {
	if len(channelIds) == 0 {
		return m.Where("1 = 0")
	}

	conditions := make([]string, 0, len(channelIds)*2)
	args := make([]interface{}, 0, len(channelIds)*2)
	for _, channelId := range channelIds {
		conditions = append(conditions,
			fmt.Sprintf("%s = ?", channelIdColumn),
			fmt.Sprintf("FIND_IN_SET(?, %s) > 0", channelIdsColumn),
		)
		args = append(args, channelId, channelId)
	}
	return m.Where("("+strings.Join(conditions, " OR ")+")", args...)
}

func getContentChannelMap(ctx context.Context, contentChannelIds []uint) map[uint]*entity.WebChannel {
	channelMap := make(map[uint]*entity.WebChannel, len(contentChannelIds))
	if len(contentChannelIds) == 0 {
		return channelMap
	}

	var channels []*entity.WebChannel
	_ = dao.WebChannel.Ctx(ctx).WhereIn("id", contentChannelIds).Where("status", 1).Scan(&channels)
	for _, channel := range channels {
		formatChannelUrl(channel)
		channelMap[channel.Id] = channel
	}
	return channelMap
}

func getBreadcrumb(ctx context.Context, channel *entity.WebChannel) []*entity.WebChannel {
	breadcrumb := make([]*entity.WebChannel, 0)
	for current := channel; current != nil; {
		breadcrumb = append([]*entity.WebChannel{current}, breadcrumb...)
		if current.Pid == 0 {
			break
		}

		var parent *entity.WebChannel
		_ = dao.WebChannel.Ctx(ctx).Where("id", current.Pid).Where("status", 1).Scan(&parent)
		if parent != nil {
			formatChannelUrl(parent)
		}
		current = parent
	}
	return breadcrumb
}

func getRelatedChannels(ctx context.Context, pid uint) []*entity.WebChannel {
	var relatedList []*entity.WebChannel
	_ = dao.WebChannel.Ctx(ctx).Where("pid", pid).Where("status", 1).Order("weigh desc, id asc").Scan(&relatedList)
	if relatedList == nil {
		return make([]*entity.WebChannel, 0)
	}
	for _, channel := range relatedList {
		formatChannelUrl(channel)
	}
	return relatedList
}

func formatArticle(article *entity.WebArticle, channelMap map[uint]*entity.WebChannel, fallbackChannel *entity.WebChannel) {
	if article == nil {
		return
	}
	normalizeImageFields(&article.Image, &article.Images)
	channel := channelMap[article.ChannelId]
	if channel == nil {
		channel = fallbackChannel
	}
	formatArticleUrl(article, channel)
}

func formatArticles(articles []*entity.WebArticle, channelMap map[uint]*entity.WebChannel, fallbackChannel *entity.WebChannel) {
	for _, article := range articles {
		formatArticle(article, channelMap, fallbackChannel)
	}
}

type articleAuthorStats struct {
	ArticleCount int `json:"articleCount"`
	ViewCount    int `json:"viewCount"`
}

func (s *sApiSite) loadArticleAuthor(ctx context.Context, article *entity.WebArticle) (*sitein.ContentAuthor, error) {
	if article == nil {
		return nil, nil
	}

	if article.UserId > 0 {
		var user *entity.User
		cols := dao.User.Columns()
		if err := dao.User.Ctx(ctx).Fields(cols.Id, cols.Username, cols.Nickname, cols.Avatar, cols.Motto).WherePri(article.UserId).Scan(&user); err != nil {
			return nil, gerror.Wrap(err, "获取文章作者信息失败，请稍后重试")
		}
		if user != nil {
			stats, err := s.loadArticleAuthorStats(ctx, dao.WebArticle.Columns().UserId, article.UserId)
			if err != nil {
				return nil, err
			}
			return &sitein.ContentAuthor{
				Id:           user.Id,
				Username:     user.Username,
				Nickname:     user.Nickname,
				Avatar:       user.Avatar,
				Motto:        user.Motto,
				ArticleCount: stats.ArticleCount,
				ViewCount:    stats.ViewCount,
			}, nil
		}
	}

	if article.AdminId == 0 {
		return nil, nil
	}
	var admin *entity.AdminMember
	cols := dao.AdminMember.Columns()
	if err := dao.AdminMember.Ctx(ctx).Fields(cols.Id, cols.Username, cols.RealName, cols.Avatar).WherePri(article.AdminId).Scan(&admin); err != nil {
		return nil, gerror.Wrap(err, "获取文章作者信息失败，请稍后重试")
	}
	if admin == nil {
		return nil, nil
	}
	stats, err := s.loadArticleAuthorStats(ctx, dao.WebArticle.Columns().AdminId, article.AdminId)
	if err != nil {
		return nil, err
	}
	return &sitein.ContentAuthor{
		Id:           uint(admin.Id),
		Username:     admin.Username,
		Nickname:     admin.RealName,
		Avatar:       admin.Avatar,
		IsAdmin:      true,
		ArticleCount: stats.ArticleCount,
		ViewCount:    stats.ViewCount,
	}, nil
}

func (s *sApiSite) loadArticleAuthorStats(ctx context.Context, authorColumn string, authorID interface{}) (stats articleAuthorStats, err error) {
	cols := dao.WebArticle.Columns()
	err = dao.WebArticle.Ctx(ctx).
		Fields(fmt.Sprintf("COUNT(*) AS article_count, COALESCE(SUM(%s), 0) AS view_count", cols.Views)).
		Where(authorColumn, authorID).
		Where(cols.Status, 1).
		Scan(&stats)
	if err != nil {
		return stats, gerror.Wrap(err, "获取作者文章统计失败，请稍后重试")
	}
	return stats, nil
}

func formatProduct(product *entity.ShopProduct, channelMap map[uint]*entity.WebChannel, fallbackChannel *entity.WebChannel) {
	if product == nil {
		return
	}
	normalizeImageFields(&product.Image, &product.Images)
	channel := channelMap[product.ChannelId]
	if channel == nil {
		channel = fallbackChannel
	}
	formatProductUrl(product, channel)
}

func formatProducts(products []*entity.ShopProduct, channelMap map[uint]*entity.WebChannel, fallbackChannel *entity.WebChannel) {
	for _, product := range products {
		formatProduct(product, channelMap, fallbackChannel)
	}
}

func getProductChannelIds(products []*entity.ShopProduct) []uint {
	channelIds := make([]uint, 0)
	channelIdSet := make(map[uint]struct{})
	for _, product := range products {
		if _, exists := channelIdSet[product.ChannelId]; !exists {
			channelIds = append(channelIds, product.ChannelId)
			channelIdSet[product.ChannelId] = struct{}{}
		}
	}
	return channelIds
}

func getArticleChannelIds(articles []*entity.WebArticle) []uint {
	channelIds := make([]uint, 0)
	channelIdSet := make(map[uint]struct{})
	for _, article := range articles {
		if _, exists := channelIdSet[article.ChannelId]; !exists {
			channelIds = append(channelIds, article.ChannelId)
			channelIdSet[article.ChannelId] = struct{}{}
		}
	}
	return channelIds
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
		siteMap["siteFavicon"] = basicConfig.Favicon
		siteMap["siteDomain"] = basicConfig.Domain
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
		siteMap["sitePoliceRecord"] = siteConfig.PoliceRecord
		siteMap["siteemail"] = siteConfig.Email
		siteMap["sitephone"] = siteConfig.Phone
		siteMap["sitewechatQrCode"] = siteConfig.WechatQrCode
		siteMap["sitempQrCode"] = siteConfig.MpQrCode
	}

	res.SiteConfig = siteMap
	columns := dao.WebNavMenu.Columns()
	var menus []*entity.WebNavMenu
	if err = dao.WebNavMenu.Ctx(ctx).Where(columns.Status, consts.StatusEnabled).
		OrderAsc(columns.Sort).OrderAsc(columns.Id).Scan(&menus); err != nil {
		return nil, gerror.Wrap(err, "获取网站导航失败")
	}
	res.NavMenu = buildNavMenuTree(menus)

	return res, nil
}

// buildNavMenuTree 只遍历从根可达的启用节点，避免禁用父节点的子项被提升为顶级导航。
func buildNavMenuTree(menus []*entity.WebNavMenu) []*sitein.NavMenuItem {
	children := make(map[int64][]*entity.WebNavMenu)
	for _, menu := range menus {
		children[menu.Pid] = append(children[menu.Pid], menu)
	}
	visited := make(map[int64]bool)
	var build func(int64) []*sitein.NavMenuItem
	build = func(pid int64) []*sitein.NavMenuItem {
		items := make([]*sitein.NavMenuItem, 0, len(children[pid]))
		for _, menu := range children[pid] {
			if visited[menu.Id] {
				continue
			}
			visited[menu.Id] = true
			items = append(items, &sitein.NavMenuItem{
				Id: menu.Id, Name: menu.Name, Url: menu.Url,
				Image: menu.Image, ImageAlt: menu.ImageAlt, Children: build(menu.Id),
			})
		}
		return items
	}
	return build(0)
}

// Home 获取首页各个栏目的聚合数据
func (s *sApiSite) Home(ctx context.Context, in sitein.HomeInp) (res *sitein.HomeModel, err error) {
	res = new(sitein.HomeModel)

	// 处理参数
	productsChannelIdsArr := resolveChannelIds(ctx, in.ProductsChannelIds, []int{7, 8, 9, 10, 11}) // 默认产品分类
	productsLimit := 4                                                                             // 默认显示4条
	if in.ProductsLimit > 0 {
		productsLimit = in.ProductsLimit
	}

	casesChannelIdsArr := resolveChannelIds(ctx, in.CasesChannelIds, []int{16}) // 默认应用案例/经典案例栏目ID为16 (diyname: cases)
	casesLimit := 3
	if in.CasesLimit > 0 {
		casesLimit = in.CasesLimit
	}

	newsChannelIdsArr := resolveChannelIds(ctx, in.NewsChannelIds, []int{2}) // 默认新闻中心根栏目ID为2
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
				List:       make([]*entity.ShopProduct, 0),
			}
			channelIds, channelMap := getChannelIdsAndMap(ctx, ch)

			// 查询该栏目及子栏目下的产品内容
			var chProducts []*entity.ShopProduct
			err = dao.ShopProduct.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).Order("weigh desc, id desc").Limit(productsLimit).Scan(&chProducts)
			if err == nil && chProducts != nil {
				formatProducts(chProducts, channelMap, ch)
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
			channelIds, channelMap := getChannelIdsAndMap(ctx, ch)

			var chCases []*entity.WebArticle
			// 经典案例一般带有图片 (数据库字段为 images)
			err = dao.WebArticle.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).WhereNot("images", "").Order("weigh desc, id desc").Limit(casesLimit).Scan(&chCases)
			if err == nil && chCases != nil {
				formatArticles(chCases, channelMap, ch)
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
			channelIds, channelMap := getChannelIdsAndMap(ctx, ch)

			var chNews []*entity.WebArticle
			err = dao.WebArticle.Ctx(ctx).Where("status", "normal").WhereIn("channel_id", channelIds).Order("publish_time desc, id desc").Limit(newsLimit).Scan(&chNews)
			if err == nil && chNews != nil {
				formatArticles(chNews, channelMap, ch)
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

	columns := dao.WebCarousel.Columns()
	res.Carousels = make([]*sitein.CarouselItem, 0)
	if err = dao.WebCarousel.Ctx(ctx).Where(columns.Status, consts.StatusEnabled).
		OrderAsc(columns.Sort).OrderAsc(columns.Id).Scan(&res.Carousels); err != nil {
		return nil, gerror.Wrap(err, "获取首页轮播图失败")
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

			res.Breadcrumb = getBreadcrumb(ctx, res.Channel)
			res.RelatedList = getRelatedChannels(ctx, res.Channel.Pid)
			if in.Page <= 0 {
				in.Page = 1
			}
			if in.Limit <= 0 {
				in.Limit = 10
			}
			res.Page = in.Page

			channelIds, channelMap := getChannelIdsAndMap(ctx, res.Channel)

			if res.Channel.Model == "article" {
				m := filterContentByChannelIds(
					dao.WebArticle.Ctx(ctx).Where("status", 1),
					dao.WebArticle.Columns().ChannelId,
					dao.WebArticle.Columns().ChannelIds,
					channelIds,
				)
				var list []*entity.WebArticle
				var lst interface{} = &list
				res.Total, err = s.handleList(ctx, m, in, lst)
				if list == nil {
					list = make([]*entity.WebArticle, 0)
				}
				formatArticles(list, channelMap, res.Channel)
				res.List = list

				// 获取热门文章 (根据 views 排序)
				mHot := filterContentByChannelIds(
					dao.WebArticle.Ctx(ctx).Where("status", 1),
					dao.WebArticle.Columns().ChannelId,
					dao.WebArticle.Columns().ChannelIds,
					channelIds,
				)
				var hotList []*entity.WebArticle
				var hotLst interface{} = &hotList
				err = s.handleHotList(ctx, mHot, hotLst)
				if hotList == nil {
					hotList = make([]*entity.WebArticle, 0)
				}
				formatArticles(hotList, channelMap, res.Channel)
				res.HotList = hotList
			} else if res.Channel.Model == "products" {
				m := filterContentByChannelIds(
					dao.ShopProduct.Ctx(ctx).Where("status", 1),
					dao.ShopProduct.Columns().ChannelId,
					dao.ShopProduct.Columns().ChannelIds,
					channelIds,
				)
				var list []*entity.ShopProduct
				var lst interface{} = &list
				res.Total, err = s.handleList(ctx, m, in, lst)
				if list == nil {
					list = make([]*entity.ShopProduct, 0)
				}
				formatProducts(list, channelMap, res.Channel)
				res.List = list

				// 获取热门产品 (根据 views 排序)
				mHot := filterContentByChannelIds(
					dao.ShopProduct.Ctx(ctx).Where("status", 1),
					dao.ShopProduct.Columns().ChannelId,
					dao.ShopProduct.Columns().ChannelIds,
					channelIds,
				)
				var hotList []*entity.ShopProduct
				var hotLst interface{} = &hotList
				err = s.handleHotList(ctx, mHot, hotLst)
				if hotList == nil {
					hotList = make([]*entity.ShopProduct, 0)
				}
				formatProducts(hotList, channelMap, res.Channel)
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

	res.Breadcrumb = getBreadcrumb(ctx, res.Channel)

	if res.Channel.Model == "article" {
		var article *entity.WebArticle
		err = dao.WebArticle.Ctx(ctx).Where("id", in.Id).Where("status", 1).Scan(&article)
		if err != nil {
			return nil, err
		}
		if article != nil {
			formatArticle(article, nil, res.Channel)
			author, err := s.loadArticleAuthor(ctx, article)
			if err != nil {
				return nil, err
			}
			res.Content = article
			res.Author = author

			var prevItem *entity.WebArticle
			var nextItem *entity.WebArticle
			s.handlePrevNext(ctx, dao.WebArticle.Ctx(ctx), int64(res.Channel.Id), int64(in.Id), &prevItem, &nextItem)
			if prevItem != nil {
				formatArticle(prevItem, nil, res.Channel)
				res.Prev = prevItem
			}
			if nextItem != nil {
				formatArticle(nextItem, nil, res.Channel)
				res.Next = nextItem
			}

			// 每次访问增加一次浏览量
			dao.WebArticle.Ctx(ctx).Where("id", article.Id).Increment("views", 1)
		} else {
			return nil, gerror.New("内容不存在或已下架")
		}

		res.RelatedList = getRelatedChannels(ctx, res.Channel.Pid)

		// 获取热门内容
		mHot := dao.WebArticle.Ctx(ctx).Where("status", 1).Where("channel_id", res.Channel.Id)
		var hotList []*entity.WebArticle
		var hotLst interface{} = &hotList
		_ = s.handleHotList(ctx, mHot, hotLst)
		if hotList == nil {
			hotList = make([]*entity.WebArticle, 0)
		}
		formatArticles(hotList, nil, res.Channel)
		res.HotList = hotList

	} else if res.Channel.Model == "products" {
		var product *entity.ShopProduct
		err = dao.ShopProduct.Ctx(ctx).Where("id", in.Id).Where("status", 1).Scan(&product)
		if err != nil {
			return nil, err
		}
		if product != nil {
			formatProduct(product, nil, res.Channel)
			res.Content = product

			var prevItem *entity.ShopProduct
			var nextItem *entity.ShopProduct
			s.handlePrevNext(ctx, dao.ShopProduct.Ctx(ctx), int64(res.Channel.Id), int64(in.Id), &prevItem, &nextItem)
			if prevItem != nil {
				formatProduct(prevItem, nil, res.Channel)
				res.Prev = prevItem
			}
			if nextItem != nil {
				formatProduct(nextItem, nil, res.Channel)
				res.Next = nextItem
			}

			// 每次访问增加一次浏览量
			dao.ShopProduct.Ctx(ctx).Where("id", product.Id).Increment("views", 1)
		} else {
			return nil, gerror.New("内容不存在或已下架")
		}

		res.RelatedList = getRelatedChannels(ctx, res.Channel.Pid)

		// 获取热门内容
		mHot := dao.ShopProduct.Ctx(ctx).Where("status", 1).Where("channel_id", res.Channel.Id)
		var hotList []*entity.ShopProduct
		var hotLst interface{} = &hotList
		_ = s.handleHotList(ctx, mHot, hotLst)
		if hotList == nil {
			hotList = make([]*entity.ShopProduct, 0)
		}
		formatProducts(hotList, nil, res.Channel)
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
		m := dao.ShopProduct.Ctx(ctx).Where("status", 1).Where("(title LIKE ? OR description LIKE ?)", likeKeywords, likeKeywords)
		totalCount, countErr := m.Count()
		if countErr != nil {
			return nil, countErr
		}
		res.Total = totalCount

		var list []*entity.ShopProduct
		if totalCount > 0 {
			err = m.Page(page, limit).Order("weigh desc, id desc").Scan(&list)
			if err != nil {
				return nil, err
			}
		}
		if list == nil {
			list = make([]*entity.ShopProduct, 0)
		}

		if len(list) > 0 {
			formatProducts(list, getContentChannelMap(ctx, getProductChannelIds(list)), nil)
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
			formatArticles(list, getContentChannelMap(ctx, getArticleChannelIds(list)), nil)
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
