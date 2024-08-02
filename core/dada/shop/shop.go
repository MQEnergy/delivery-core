package shop

import (
	"github.com/MQEnergy/delivery-core/core/dada"
	dada2 "github.com/MQEnergy/delivery-core/routers/dada"
)

type Shop struct {
	Base *dada.Base
}

func New(sourceID, appKey, appSecret string, online bool) *Shop {
	return &Shop{
		Base: dada.New(sourceID, appKey, appSecret, online),
	}
}

// CityCodeList is a list of city codes visit: https://newopen.imdada.cn/#/development/file/cityList
func (s *Shop) CityCodeList(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(dada2.CITY_CODE_LIST_URL, params).Result()
}

// QueryCates query cates visit: https://newopen.imdada.cn/#/development/file/cityList
func (s *Shop) QueryCates(params map[string]interface{}) string {
	return `[{"cate_id":1,"cate_name":"食品小吃"},{"cate_id":2,"cate_name":"饮料"},{"cate_id":3,"cate_name":"鲜花绿植"},{"cate_id":5,"cate_name":"其他"},{"cate_id":8,"cate_name":"文印票务"},{"cate_id":9,"cate_name":"便利店"},{"cate_id":13,"cate_name":"水果生鲜"},{"cate_id":19,"cate_name":"同城电商"},{"cate_id":20,"cate_name":"医药"},{"cate_id":21,"cate_name":"蛋糕"},{"cate_id":24,"cate_name":"酒品"},{"cate_id":25,"cate_name":"小商品市场"},{"cate_id":26,"cate_name":"服装"},{"cate_id":27,"cate_name":"汽修零配"},{"cate_id":28,"cate_name":"数码家电"},{"cate_id":29,"cate_name":"小龙虾/烧烤"},{"cate_id":31,"cate_name":"超市"},{"cate_id":51,"cate_name":"火锅"},{"cate_id":54,"cate_name":"个护美妆"},{"cate_id":55,"cate_name":"母婴"},{"cate_id":57,"cate_name":"家居家纺"},{"cate_id":59,"cate_name":"手机"},{"cate_id":61,"cate_name":"家装"},{"cate_id":63,"cate_name":"成人用品"},{"cate_id":65,"cate_name":"校园"},{"cate_id":66,"cate_name":"高端市场"}]`
}

// ShopAdd add shop visit: https://newopen.imdada.cn/#/development/file/shopAdd
func (s *Shop) ShopAdd(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(dada2.SHOP_ADD_URL, params).Result()
}

// ShopUpdate update shop visit: https://newopen.imdada.cn/#/development/file/shopUpdate
func (s *Shop) ShopUpdate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(dada2.SHOP_UPDATE_URL, params).Result()
}

// ShopDetail get shop detail visit: https://newopen.imdada.cn/#/development/file/shopDetail
func (s *Shop) ShopDetail(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(dada2.SHOP_DETAIL_URL, params).Result()
}
