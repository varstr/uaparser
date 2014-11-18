package uaparser

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	var expectedBrowserNames map[string][]string = GetBrowserNames()
	var expectedOperatingSystems map[string][]string = GetOSNames()

	_checkExepectations(t, expectedOperatingSystems, "os")
	_checkExepectations(t, expectedBrowserNames, "browser")
}

func _checkExepectations(t *testing.T, expectations map[string][]string, testType string) {
	var uaParseResult *UAInfo
	var testResult bool
	var comparedTo string

	for expectation := range expectations {
		for key := range expectations[expectation] {
			fmt.Printf("Testing: %s \n", expectations[expectation][key])
			uaParseResult = Parse(expectations[expectation][key])

			if testType == "browser" {
				testResult, comparedTo = _checkBrowser(uaParseResult, expectation)
			} else if testType == "os" {
				testResult, comparedTo = _checkOs(uaParseResult, expectation)
			}

			if !testResult {
				t.Fatalf("Expected: '%s' got '%s' on: '%s'", expectation, comparedTo, expectations[expectation][key])
			}
		}
	}
}

func _checkBrowser(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.Browser == nil {
		return false, ""
	}
	return (uainfo.Browser.Name == expectation), uainfo.Browser.Name
}

func _checkOs(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.OS == nil {
		return false, ""
	}
	return (uainfo.OS.Name == expectation), uainfo.OS.Name
}
