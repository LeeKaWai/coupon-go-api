package service

import (
	"coupon-go-api/config"
	"coupon-go-api/pkg/taobao"
	"coupon-go-api/src/models"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

var client = taobao.NewDefaultTopClient("", "", "", 5000, 6000)

// 通过淘口令获取商品信息
func TKLSearch(password string, relation_id string) (models.Result, error) {
	if relation_id == "" {
		relation_id = config.GLOBAL_CONF.Taobao.Relationid
	}
	query := url.Values{}
	query.Set("appkey", config.GLOBAL_CONF.Taobao.ZtkAppkey)
	query.Set("sid", config.GLOBAL_CONF.Taobao.ZtkSid)
	query.Set("pid", config.GLOBAL_CONF.Taobao.RelationPid)
	query.Set("tkl", password)
	query.Set("relation_id", relation_id)
	query.Set("signurl", "5")

	url := "https://api.zhetaoke.com:10001/api/open_gaoyongzhuanlian_tkl.ashx?" + query.Encode()
	// 发起 HTTP GET 请求
	request, err := http.Get(url)
	if err != nil {
		return models.Result{}, err
	}
	defer request.Body.Close()

	// 读取响应数据
	var content models.ResContent
	err = json.NewDecoder(request.Body).Decode(&content)
	if err != nil {
		return models.Result{}, err
	}

	contents := ConvertToMyResult(content.Content)
	result := models.Result{
		Status:  strconv.Itoa(content.Status),
		Message: "Success",
		Data:    contents,
	}

	return result, nil
}

func ConvertToMyResult(response []models.CouponResult) []models.ContentResult {
	var results []models.ContentResult
	for _, item := range response {
		results = append(results, models.ContentResult(item))
	}
	return results
}

// func TitleSearch(title string, relation_id string) (models.Result, error) {
// 	if relation_id == "" {
// 		relation_id = config.GLOBAL_CONF.Taobao.Relationid
// 	}
// 	var content models.CouponResult

// 	// 声明一个空的 map，键为 string 类型，值可以是任意类型
// 	var myMap map[string]interface{}

// 	// 初始化 map
// 	myMap = make(map[string]interface{})

// 	// 添加键值对
// 	myMap["q"] = title
// 	myMap["relation_id"] = relation_id
// 	myMap["adzone_id"] = config.GLOBAL_CONF.Taobao.Adzoneid
// 	fileMap := make(map[string]interface{})

// 	res, err := client.Execute("taobao.tbk.dg.material.optional.upgrade", myMap, fileMap)
// 	if err != nil {
// 		return content, err
// 	}
// 	fmt.Printf("%s", res)
// 	return content, nil
// }
