package domain

type (
	// SakeEntity is sake struct
	SakeEntity struct {
		Name              string `json:"name" csv:"商品名"`
		Category          string `json:"category" csv:"種類"`
		Brewer            string `json:"brewer" csv:"蔵元"`
		Location          string `json:"location" csv:"蔵元所在地"`
		Alcohol           string `json:"alcohol" csv:"アルコール分"`
		Capacity          string `json:"capacity" csv:"内容量"`
		RawRice           string `json:"raw_rice" csv:"原料米"`
		RicePolishingRate string `json:"rice_polishing_rate" csv:"精米歩合"`
		Burning           string `json:"burning" csv:"火入"`
		SakeDegree        string `json:"sake_degree" csv:"日本酒度"`
		Acidity           string `json:"acidity" csv:"酸度"`
		AminoAcidity      string `json:"amino_acidity" csv:"アミノ酸度"`
		Yeast             string `json:"yeast" csv:"酵母"`
		StorageMethod     string `json:"storage_method" csv:"保管方法"`
		Price             string `json:"price" csv:"価格"`
		Remarks           string `json:"remarks" csv:"備考"`
		URL               string `json:"url" csv:"参照サイトURL"`
	}
)
