package caption

type GetInfoResponse struct {
	Photo *Photo `json:"photo"`
}

type Photo struct {
	Id          string       `json:"id"`
	Owner       *Owner       `json:"owner"`
	Title       *Title       `json:"title"`
	Description *Description `json:"description"`
	Dates       *Dates       `json:"dates"`
}

type Owner struct {
	NSID     string `json:"nsid"`
	UserName string `json:"username"`
	RealName string `json:"realname"`
}

type Title struct {
	Content string `json:"_content"`
}

type Description struct {
	Content string `json:"_content"`
}

type Dates struct {
	Posted           string `json:"posted"`
	Taken            string `json"taken"`
	TakenGranularity string `json:"takengranularity"`
	TakenUnknown     string `json:"takenunknown"`
	LastUpdate       string `json:"lastupdate"`
}
