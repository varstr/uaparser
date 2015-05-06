package uaparser

import (
	"strings"
	"unicode"
)

type itemSpec struct {
	name             string
	mustContains     []string
	mustNotContains  []string
	versionSplitters [][]string
}

type InfoItem struct {
	Name    string
	Version string
}

type UAInfo struct {
	Browser,
	Device,
	DeviceType,
	OS *InfoItem
}

func isEmptyString(str string) bool {
	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func contains(ua string, tokens []string) bool {
	for _, tk := range tokens {
		if strings.Contains(ua, tk) {
			return true
		}
	}
	return false
}

func matchSpec(ua string, spec *itemSpec) (info *InfoItem, ok bool) {
	if !contains(ua, spec.mustContains) {
		return
	}
	if contains(ua, spec.mustNotContains) {
		return
	}

	info = new(InfoItem)
	info.Name = spec.name
	ok = true

	for _, splitter := range spec.versionSplitters {
		if strings.Contains(ua, splitter[0]) {
			if rmLeft := strings.Split(ua, splitter[0])[1]; strings.Contains(rmLeft, splitter[1]) || isEmptyString(splitter[1]) {
				rmRight := strings.Split(rmLeft, splitter[1])[0]
				info.Version = strings.TrimSpace(rmRight)
				break
			}
		}
	}
	return
}

func searchIn(ua string, specs []*itemSpec) (info *InfoItem) {
	for _, spec := range specs {
		if result, ok := matchSpec(ua, spec); ok {
			info = result
			break
		}
	}
	return
}

func Parse(ua string) (info *UAInfo) {
	info = new(UAInfo)

	info.Browser = searchIn(ua, _BROWSERS)
	info.Device = searchIn(ua, _DEVICES)
	info.DeviceType = searchIn(ua, _DEVICETYPES)
	info.OS = searchIn(ua, _OS)

	return
}
