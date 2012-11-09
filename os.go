package uaparser

var linux *itemSpec = &itemSpec{
    name: "Linux",
    mustContains: []string{"Linux"},
    mustNotContains: []string{"Android"},
}

var macOS *itemSpec = &itemSpec{
    name: "Mac OS",
    mustContains: []string{"Mac OS"},
    mustNotContains: []string{"iPad", "iPhone", "iPod"},
    versionSplitters: [][]string{
        []string{"Mac OS ", ")"},
        []string{"Mac OS ", ";"},
    },
}

var windows *itemSpec = &itemSpec{
    name: "Windows",
    mustContains: []string{"Windows"},
    mustNotContains: []string{"Windows Phone"},
    versionSplitters: [][]string{
        []string{"Windows ", ";"},
    },
}

var android *itemSpec = &itemSpec{
    name: "Android",
    mustContains: []string{"Android"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Android ", ";"},
        []string{"Android-", " "},
    },
}

var iOS *itemSpec = &itemSpec{
    name: "iOS",
    mustContains: []string{"CPU", "OS", "like Mac OS X"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"CPU iPhone OS ", " "},
        []string{"CPU OS ", " "},
    },
}

var wpOS *itemSpec = &itemSpec{
    name: "Windows Phone OS",
    mustContains: []string{"Windows Phone OS"},
    mustNotContains: []string{},
    versionSplitters: [][]string{
        []string{"Windows Phone OS ", ";"},
    },
}

var _OS []*itemSpec = []*itemSpec {
    linux,
    macOS,
    windows,
    android,
    iOS,
    wpOS,
}
