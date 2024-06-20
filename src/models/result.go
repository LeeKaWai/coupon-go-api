package models

type Result struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResContent struct {
	Status  int            `json:"status"`
	Content []CouponResult `json:"content"`
}

type ContentResult struct {
	// 商品 id
	Id string
	// 商品标题
	Title string
	// 淘宝标题
	TaoTitle string
	// 商品简介
	Desc string
	// 商品主图
	PictUrl string
	// 是否天猫 0 淘宝 1 天猫
	UserType string
	// 卖家 id
	SellerId string
	// 折扣价
	Discount string
	// 券后价
	PostCouponPrice string
	// 佣金比例
	CommissionRate string
	// 优惠券开始时间
	CouponStartTime string
	// 优惠券结束时间
	CouponEndTime string
	// 优惠券金额
	CouponInfoMoney string
	// 优惠券总数量
	CouponTotalCount string
	// 优惠券剩余数量
	CouponRemainCount string
	// 优惠券信息 ex:满80.00元减10元
	CouponInfo string
	// 卖家昵称
	Nick string
	// 商品小图(多图，采用 ｜ 隔开)
	SmallImages string
	// 商品白底图
	WhiteImage string
	// 实时总销量
	SellCount string
	// 评论数量
	CommentCount string
	// 店铺 logo
	ShopIcon string
	// 商品 url
	TaobaoUrl string
	// 返佣金额
	CommissionFee string
	// 二合一推广连接
	CouponUrl string
	// 推广长链接
	ItemUrl string
	// 淘宝短连接
	ShortUrl string
	// 淘口令
	Tkl string
}

type CouponResult struct {
	// 商品 id
	Id string `json:"tao_id"`
	// 商品标题
	Title string `json:"title"`
	// 淘宝标题
	TaoTitle string `json:"tao_title"`
	// 商品简介
	Desc string `json:"jianjie"`
	// 商品主图
	PictUrl string `json:"pict_url"`
	// 是否天猫 0 淘宝 1 天猫
	UserType string `json:"user_type"`
	// 卖家 id
	SellerId string `json:"seller_id"`
	// 折扣价
	Discount string `json:"size"`
	// 券后价
	PostCouponPrice string `json:"quanhou_jiage"`
	// 佣金比例
	CommissionRate string `json:"tkrate3"`
	// 优惠券开始时间
	CouponStartTime string `json:"coupon_start_time"`
	// 优惠券结束时间
	CouponEndTime string `json:"coupon_end_time"`
	// 优惠券金额
	CouponInfoMoney string `json:"coupon_info_money"`
	// 优惠券总数量
	CouponTotalCount string `json:"coupon_total_count"`
	// 优惠券剩余数量
	CouponRemainCount string `json:"coupon_remain_count"`
	// 优惠券信息 ex:满80.00元减10元
	CouponInfo string `json:"coupon_info"`
	// 卖家昵称
	Nick string `json:"nick"`
	// 商品小图(多图，采用 ｜ 隔开)
	SmallImages string `json:"small_images"`
	// 商品白底图
	WhiteImage string `json:"white_image"`
	// 实时总销量
	SellCount string `json:"sell_count"`
	// 评论数量
	CommentCount string `json:"comment_count"`
	// 店铺 logo
	ShopIcon string `json:"shopIcon"`
	// 商品 url
	TaobaoUrl string `json:"taobao_url"`
	// 返佣金额
	CommissionFee string `json:"tkfee3"`
	// 二合一推广连接
	CouponUrl string `json:"coupon_click_url"`
	// 推广长链接
	ItemUrl string `json:"item_url"`
	// 淘宝短连接
	ShortUrl string `json:"shorturl"`
	// 淘口令
	Tkl string `json:"tkl"`
}
