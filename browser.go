package uaparser

type browserSpec struct {
    name string
    mustContains []string
    mustNotContains []string
    versionSplitters [][]string
}

var ie *browserSpec = &browserSpec {
    name: "IE",
    mustContains: []string{"MSIE"},
    mustNotContains: []string{
        "360SE",
        "Maxthon",
        "SE", "MetaSr",
        "TencentTraveler", "QQBrowser",
    },
    versionSplitters: [][]string{[]string{"MSIE ", ";"}},
}

var firefox *browserSpec = &browserSpec {
    name: "Firefox",
    mustContains: []string{"Firefox"},
    mustNotContains: []string{"Seamonkey"},
    versionSplitters: [][]string{[]string{"Firefox/", " "}},
}

var safari *browserSpec = &browserSpec {
    name: "Safari",
    mustContains: []string{"Safari"},
    mustNotContains: []string{
        "Chrome", "Chromium",
        "Maxthon",
        "QQBrowser",
        "SE", "MetaSr",
    },
    versionSplitters: [][]string{
        []string{"Version/", " "},
        []string{"Safari/", " "},
    },
}

var chrome *browserSpec = &browserSpec {
    name: "Chrome",
    mustContains: []string{"Chrome"},
    mustNotContains: []string{
        "Chromium",
        "Maxthon",
        "QQBrowser",
        "SE", "MetaSr",
    },
    versionSplitters: [][]string{[]string{"Chrome/", " "}},
}

var opera *browserSpec = &browserSpec {
    name: "Opera",
    mustContains: []string{"Opera"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Version/", " "},
        []string{"Opera/", " "},
    },
}

var _360se *browserSpec = &browserSpec {
    name: "360SE",
    mustContains: []string{"360SE"},
    mustNotContains: []string{},
}

var sougou *browserSpec = &browserSpec {
    name: "Sougou",
    mustContains: []string{"SE", "MetaSr"},
    mustNotContains: []string{},
    versionSplitters: [][]string{[]string{"SE ", " "}},
}

var tencent *browserSpec = &browserSpec {
    name: "Tencent",
    mustContains: []string{"TencentTraveler"},
    mustNotContains: []string{"SE", "MetaSr"},
    versionSplitters: [][]string{[]string{"TencentTraveler ", " "}},
}

var qq *browserSpec = &browserSpec {
    name: "QQ",
    mustContains: []string{"QQBrowser"},
    mustNotContains: []string{},
    versionSplitters: [][]string{[]string{"QQBrowser/", " "}},
}

var maxthon *browserSpec = &browserSpec {
    name: "Maxthon",
    mustContains: []string{"Maxthon"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Maxthon/", " "},
        []string{"Maxthon ", " "},
    },
}

var _BROWSERS []*browserSpec = []*browserSpec {
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

