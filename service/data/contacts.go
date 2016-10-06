package data

type Contact struct {
	Type    string `json:"type"`
	Contact string `json:"contact"`
}
type CompanyInfo struct {
	Icon     string    `json:"icon"`
	Contacts []Contact `json:"contacts"`
}

var Companies map[string]CompanyInfo

func init() {
	Companies = make(map[string]CompanyInfo)
	Companies["NambaTaxi"] = CompanyInfo{
		Icon: "/static/namba.svg",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "9797",
			},
			Contact{
				Type:    "phone",
				Contact: "+996559976000",
			},
		},
	}
	Companies["SmsTaxi"] = CompanyInfo{
		Icon: "/static/logo-sms.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "1236",
			},
			Contact{
				Type:    "phone",
				Contact: "+996551061236",
			},
		},
	}
	Companies["Diesel"] = CompanyInfo{
		Icon: "/static/logo-diesel.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "1450",
			},
			Contact{
				Type:    "phone",
				Contact: "+996552145000",
			},
		},
	}
	Companies["Jorgo"] = CompanyInfo{
		Icon: "/static/logo-jorgo.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "2022",
			},
			Contact{
				Type:    "phone",
				Contact: "+996550662022",
			},
		},
	}
}
