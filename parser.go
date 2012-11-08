package uaparser

import (
    "strings"
)

type InfoItem struct {
    Name string
    Version string
}

type UAInfo struct {
    Browser,
    Device,
    Os *InfoItem
}

func parseBrowser(ua string, spec *browserSpec) (info *InfoItem, ok bool){
    for _, mc := range spec.mustContains {
        if !strings.Contains(ua, mc) {
            return
        }
    }

    for _, mnc := range spec.mustNotContains {
        if strings.Contains(ua, mnc) {
            return
        }
    }

    info = new(InfoItem)
    info.Name = spec.name
    ok = true

    for _, splitter := range spec.versionSplitters {
        if strings.Contains(ua, splitter[0]) {
            info.Version = strings.TrimSpace(strings.Split(strings.Split(ua, splitter[0])[1], splitter[1])[0])
            break
        }
    }
    return
}

func Parse(ua string) (info *UAInfo) {
    info = new(UAInfo)
    for _, browser := range _BROWSERS {
        if browserInfo, ok := parseBrowser(ua, browser); ok {
            info.Browser = browserInfo
            break
        }
    }

    //TODO: parseDevice

    //TODO: parseOS

    return
}

