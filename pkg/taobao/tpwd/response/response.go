package response

type TpwdQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   淘口令-文案
	*/
	Content string `json:"content,omitempty" `
	/*
	   淘口令-宝贝
	*/
	Title string `json:"title,omitempty" `
	/*
	   如果是宝贝，则为宝贝价格
	*/
	Price string `json:"price,omitempty" `
	/*
	   图片url
	*/
	PicUrl string `json:"pic_url,omitempty" `
	/*
	   是否成功
	*/
	Suc bool `json:"suc,omitempty" `
	/*
	   跳转url(长链)
	*/
	Url string `json:"url,omitempty" `
	/*
	   nativeUrl
	*/
	NativeUrl string `json:"native_url,omitempty" `
	/*
	   thumbPicUrl
	*/
	ThumbPicUrl string `json:"thumb_pic_url,omitempty" `
}
