package uaparser

var (
    linux = &itemSpec{
        name: "Linux",
        mustContains: []string{"Linux"},
        mustNotContains: []string{"Android"},
    }

    macOS = &itemSpec{
        name: "Mac OS",
        mustContains: []string{"Mac OS"},
        mustNotContains: []string{"iPad", "iPhone", "iPod"},
        versionSplitters: [][]string{
            []string{"Mac OS ", ";"},
            []string{"Mac OS ", ")"},
        },
    }

    windows = &itemSpec{
        name: "Windows",
        mustContains: []string{"Windows"},
        mustNotContains: []string{"Windows Phone"},
        versionSplitters: [][]string{
            []string{"Windows ", ";"},
        },
    }

    android = &itemSpec{
        name: "Android",
        mustContains: []string{"Android"},
        mustNotContains: []string{},
        versionSplitters: [][]string{
            []string{"Android ", ";"},
            []string{"Android-", " "},
        },
    }
    
    iOS = &itemSpec{
        name: "iOS",
        mustContains: []string{"CPU", "OS", "like Mac OS X"},
        mustNotContains: []string{"Windows Phone OS"},
        versionSplitters: [][]string{
            []string{"CPU iPhone OS ", " "},
            []string{"CPU OS ", " "},
        },
    }

    wpOS = &itemSpec{
        name: "Windows Phone OS",
        mustContains: []string{"Windows Phone OS"},
        mustNotContains: []string{},
        versionSplitters: [][]string{
            []string{"Windows Phone OS ", ";"},
        },
    }

    _OS = []*itemSpec {
        linux,
        macOS,
        windows,
        android,
        iOS,
        wpOS,
    }
)
