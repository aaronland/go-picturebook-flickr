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

/*

{
    "photo": {
        "id": "8659998886",
        "secret": "69b0789db4",
        "server": "8124",
        "farm": 9,
        "dateuploaded": "1366257943",
        "isfavorite": 0,
        "license": "0",
        "safety_level": "0",
        "rotation": 0,
        "owner": {
            "nsid": "35034348999@N01",
            "username": "straup",
            "realname": "Aaron Straup Cope",
            "location": "",
            "iconserver": "1",
            "iconfarm": 1,
            "path_alias": "straup",
            "gift": {
                "gift_eligible": false,
                "eligible_durations": [],
                "new_flow": true
            }
        },
        "title": {
            "_content": ""
        },
        "description": {
            "_content": ""
        },
        "visibility": {
            "ispublic": 1,
            "isfriend": 0,
            "isfamily": 0
        },
        "dates": {
            "posted": "1366257943",
            "taken": "2013-04-17 21:00:44",
            "takengranularity": 0,
            "takenunknown": "0",
            "lastupdate": "1416944157"
        },
        "views": "120",
        "editability": {
            "cancomment": 0,
            "canaddmeta": 0
        },
        "publiceditability": {
            "cancomment": 1,
            "canaddmeta": 0
        },
        "usage": {
            "candownload": 0,
            "canblog": 0,
            "canprint": 0,
            "canshare": 0
        },
        "comments": {
            "_content": "0"
        },
        "notes": {
            "note": []
        },
        "people": {
            "haspeople": 0
        },
        "tags": {
            "tag": [
                {
                    "id": "6065-8659998886-101501429",
                    "author": "35034348999@N01",
                    "authorname": "straup",
                    "raw": "Aardvark filter",
                    "_content": "aardvarkfilter",
                    "machine_tag": 0
                },
                {
                    "id": "6065-8659998886-100995637",
                    "author": "35034348999@N01",
                    "authorname": "straup",
                    "raw": "flickriosapp:filter=Aardvark",
                    "_content": "flickriosapp:filter=aardvark",
                    "machine_tag": 1
                },
                {
                    "id": "6065-8659998886-97692397",
                    "author": "35034348999@N01",
                    "authorname": "straup",
                    "raw": "uploaded:by=flickr_mobile",
                    "_content": "uploaded:by=flickrmobile",
                    "machine_tag": 1
                }
            ]
        },
        "urls": {
            "url": [
                {
                    "type": "photopage",
                    "_content": "https:\/\/www.flickr.com\/photos\/straup\/8659998886\/"
                }
            ]
        },
        "media": "photo"
    },
    "stat": "ok"
}

*/
