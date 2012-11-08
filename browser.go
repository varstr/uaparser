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
    mustNotContains: []string{"Maxthon"},
    versionSplitters: [][]string{[]string{"MSIE ", ";" }},
}

var firefox *browserSpec = &browserSpec {
    name: "Firefox",
    mustContains: []string{"Firefox"},
    mustNotContains: []string{"Seamonkey"},
    versionSplitters: [][]string{[]string{"Firefox/", " " }},
}

var safari *browserSpec = &browserSpec {
    name: "Safari",
    mustContains: []string{"Safari"},
    mustNotContains: []string{"Chrome", "Chromium", "OmniWeb", "wOSBrowser"},
    versionSplitters: [][]string{
        []string{"Version/", " "},
        []string{"Safari/", " "},
    },
}

var chrome *browserSpec = &browserSpec {
    name: "Chrome",
    mustContains: []string{"Chrome"},
    mustNotContains: []string{"Chromium"},
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


var _BROWSERS []*browserSpec = []*browserSpec {
    ie,
    firefox,
    safari,
    chrome,
    opera,
}

