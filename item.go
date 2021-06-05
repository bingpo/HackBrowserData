package hackbrowserdata

type item int

const (
	Password item = iota + 1
	Bookmark
	History
	Download
	Cookie
	CreditCard
)

type Itemer interface {
	Name() string
	FileName(browser Browser) string
	Data(browser Browser) BrowsingData
}

func GetItemer(browser Browser) Itemer {
	// switch browser {
	// case browser.(webkit):
	// case browser.(gecko):
	// }
	return nil
}

const unsupportedItem = "Unsupported Item"

func (i item) Name() string {
	switch i {
	case Password:
		return "Password"
	case Bookmark:
		return "Bookmark"
	case History:
		return "History"
	case Cookie:
		return "Cookie"
	case Download:
		return "Download"
	case CreditCard:
		return "Credit Card"
	default:
		return unsupportedItem
	}
}

// FileName return filename by browser type
func (i item) FileName(browser Browser) string {
	const (
		chromePasswordFile = "Login Data"
		chromeHistoryFile  = "History"
		chromeDownloadFile = "History"
		chromeCookieFile   = "Cookies"
		chromeBookmarkFile = "Bookmarks"
		chromeCreditFile   = "Web Data"
		firefoxCookieFile  = "cookies.sqlite"
		firefoxKey4File    = "key4.db"
		firefoxLoginFile   = "logins.json"
		firefoxDataFile    = "places.sqlite"
	)
	switch browser.(type) {
	case webkit:
		switch i {
		case Password:
			return chromePasswordFile
		case Bookmark:
			return chromeBookmarkFile
		case Cookie:
			return chromeCookieFile
		case History:
			return chromeHistoryFile
		case Download:
			return chromeDownloadFile
		case CreditCard:
			return chromeCreditFile
		}
	case gecko:
		switch i {
		case Password:
			return firefoxLoginFile + "|" + firefoxKey4File
		case Bookmark:
			return firefoxDataFile
		case Cookie:
			return firefoxCookieFile
		case History:
			return firefoxDataFile
		default:
			return unsupportedItem
		}
	default:
		return unsupported
	}
	return unsupported
}

func (i item) Data(browser Browser) BrowsingData {
	switch browser.(type) {
	case webkit:
		switch i {
		case Password:
			return &WebkitPassword{}
		case Cookie:
			return &WebkitCookie{}
		case Bookmark:
			return &WebkitBookmark{}
		case History:
			return &WebkitHistory{}
		case CreditCard:
			return &WebkitCreditCard{}
		case Download:
			return &WebkitDownload{}
		}
	case gecko:
	}
	return nil
}