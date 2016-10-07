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
	Companies["Magnat"] = CompanyInfo{
		Icon: "/static/logo-magnat.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "5577",
			},
			Contact{
				Type:    "phone",
				Contact: "+996553895577",
			},
		},
	}
	Companies["Pegas"] = CompanyInfo{
		Icon: "/static/logo-pegas.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "1828",
			},
			Contact{
				Type:    "phone",
				Contact: "+996700941828",
			},
		},
	}
	Companies["Pelikan"] = CompanyInfo{
		Icon: "/static/logo-pelecan.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "+996558474747",
			},
			Contact{
				Type:    "phone",
				Contact: "+996312474747",
			},
		},
	}
	Companies["Super"] = CompanyInfo{
		Icon: "/static/logo-jorgo.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "xz",
			},
			Contact{
				Type:    "phone",
				Contact: "152",
			},
		},
	}
	Companies["Wifi"] = CompanyInfo{
		Icon: "/static/logo-wifi.png",
		Contacts: []Contact{
			Contact{
				Type:    "sms",
				Contact: "1422",
			},
			Contact{
				Type:    "phone",
				Contact: "+996552142200",
			},
		},
	}
}
