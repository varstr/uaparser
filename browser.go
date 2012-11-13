package uaparser

var ie *itemSpec = &itemSpec {
    name: "IE",
    mustContains: []string{"MSIE"},
    mustNotContains: []string{
        "360SE",
        "Maxthon",
        "QIHU",
        "QQBrowser",
        "QQDownload",
        "SE ", "MetaSr",
        "TencentTraveler",
    },
    versionSplitters: [][]string{[]string{"MSIE ", ";"}},
}

var firefox *itemSpec = &itemSpec {
    name: "Firefox",
    mustContains: []string{"Firefox"},
    mustNotContains: []string{"Seamonkey"},
    versionSplitters: [][]string{[]string{"Firefox/", " "}},
}

var safari *itemSpec = &itemSpec {
    name: "Safari",
    mustContains: []string{"Safari"},
    mustNotContains: []string{
        "Chrome", "Chromium",
        "Maxthon",
        "LBBROWSER",
        "QIHU",
        "QQBrowser",
        "SE ", "MetaSr",
        "TaoBrowser",
    },
    versionSplitters: [][]string{
        []string{"Version/", " "},
    },
}

var chrome *itemSpec = &itemSpec {
    name: "Chrome",
    mustContains: []string{"Chrome"},
    mustNotContains: []string{
        "Chromium",
        "LBBROWSER",
        "Maxthon",
        "QIHU",
        "QQBrowser",
        "SE ", "MetaSr",
        "TaoBrowser",
    },
    versionSplitters: [][]string{[]string{"Chrome/", " "}},
}

var opera *itemSpec = &itemSpec {
    name: "Opera",
    mustContains: []string{"Opera"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Version/", " "},
        []string{"Opera/", " "},
    },
}

var _360se *itemSpec = &itemSpec {
    name: "360SE",
    mustContains: []string{"360SE", "QIHU"},
    mustNotContains: []string{},
}

var sougou *itemSpec = &itemSpec {
    name: "Sougou",
    mustContains: []string{"SE ", "MetaSr"},
    mustNotContains: []string{},
    versionSplitters: [][]string{[]string{"SE ", " "}},
}

var tencent *itemSpec = &itemSpec {
    name: "Tencent",
    mustContains: []string{"TencentTraveler"},
    mustNotContains: []string{"SE ", "MetaSr"},
    versionSplitters: [][]string{[]string{"TencentTraveler ", " "}},
}

var qq *itemSpec = &itemSpec {
    name: "QQ",
    mustContains: []string{"QQBrowser"},
    mustNotContains: []string{},
    versionSplitters: [][]string{[]string{"QQBrowser/", " "}},
}

var maxthon *itemSpec = &itemSpec {
    name: "Maxthon",
    mustContains: []string{"Maxthon"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Maxthon/", " "},
        []string{"Maxthon ", " "},
    },
}

var _BROWSERS []*itemSpec = []*itemSpec {
    ie,
    firefox,
    safari,
    chrome,
    opera,
    _360se,
    sougou,
    tencent,
    qq,
    maxthon,
}

