package types

type SetMailInfoStruct struct {
	Title                string `form:"title"`
	Intro                string `form:"intro"`
	ShareImg             string `form:"shareImg"`
	ContactNumber        string `form:"contactNumber"`
	QrCode               string `form:"qrCode"`
	QrCodeDefault        int    `form:"qrCodeDefault"`
	ContactNumberDefault int    `form:"contactNumberDefault"`
}

type GoodsStruct struct {
	GoodsType          int     `form:"goodsType"`
	CategoryID         int     `form:"categoryId"`
	GoodsName          string  `form:"goodsName" binding:"required"`
	GoodsImgs          string  `form:"goodsImgs" binding:"required"`
	GoodsPrice         float64 `form:"goodsPrice"`
	GoodsShopPrice     float64 `form:"goodsShopPrice"`
	GoodsSourceContent string  `form:"goodsSourceContent"`
	GoodsSource        int     `form:"goodsSource"`
	GoodsStock         int     `form:"goodsStock"`
	GoodsDesc          string  `form:"goodsDesc"`
	LimitNum           int     `form:"limitNum"`
	GoodsSpec          string  `form:"goodsSpec"`
	GoodsSpecInfo      string  `form:"goodsSpecInfo"`
	FreightType        int     `form:"freightType"`
	Freight            float64 `form:"freight"`
	PickUpAddress      string  `form:"pickUpAddress"`
	PickUpDeadline     int     `form:"pickUpDeadline"`
	FreightTemplateID  int     `form:"freightTemplateId"`
	ContactNumber      string  `form:"contactNumber"`
	QrCode             string  `form:"qrCode"`
	GoodsContext       string  `form:"goodsContext"`
	GoodsID            int     `form:"goodsId"`
	GoodsDelImgs       string  `form:"goodsDelImgs"`
}

type GoodsSpecInfo struct {
	Specs    map[string]interface{} `json:"specs"`
	SpecsStr string                 `json:"specsStr"`
	ID       int                    `json:"id"`
	GoodsID  int                    `json:"goodsId"`
	Stock    int                    `json:"stock"`
	Sum      int                    `json:"sum"`
	Price    float64                `json:"price"`
	Img      string                 `json:"img"`
}

type ResponseGoodsSpecInfo struct {
	ID      int    `json:"id"`
	GoodsID int    `json:"goodsId"`
	Stock   int    `json:"stock"`
	Sum     int    `json:"sum"`
	Price   string `json:"price"`
	Img     string `json:"img"`
	Specs   string `json:"specs"`
}

type GoodsSpec struct {
	ID        int         `json:"id"`
	GoodsID   int         `json:"goodsId"`
	Name      string      `json:"name"`
	SpecValue []SpecValue `json:"spec_value"`
}

type SpecValue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
