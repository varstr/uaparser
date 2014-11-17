package uaparser

import "testing"
import "fmt"

func TestParse(t *testing.T) {
	var expectedBrowserNames map[string][]string = make(map[string][]string)
	expectedBrowserNames["IE"] = []string{
		"Mozilla/5.0 (compatible; MSIE 6.0; Windows NT 5.1)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; FunWebProducts)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; WIndows NT 9.0; en-US))",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 7.1; Trident/5.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; Zune 4.0; InfoPath.3; MS-RTC LM 8; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; chromeframe/12.0.742.112)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 2.0.50727; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; Zune 4.0; Tablet PC 2.0; InfoPath.3; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; yie8)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.2; .NET CLR 1.1.4322; .NET4.0C; Tablet PC 2.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; FunWebProducts)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; chromeframe/13.0.782.215)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; chromeframe/11.0.696.57)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0) chromeframe/10.0.648.205",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/4.0; GTB7.4; InfoPath.1; SV1; .NET CLR 2.8.52393; WOW64; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0; chromeframe/11.0.696.57)",
		"Mozilla/4.0(compatible; MSIE 7.0b; Windows NT 6.0)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 6.0)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; Media Center PC 3.0; .NET CLR 1.0.3705; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; FDM; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; InfoPath.1; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; InfoPath.1)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; Alexa Toolbar; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; Alexa Toolbar)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.40607)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.1; .NET CLR 1.0.3705; Media Center PC 3.1; Alexa Toolbar; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.1; DigExt)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.1)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; YComp 5.0.2.6)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; YComp 5.0.0.0) (Compatible; ; ; Trident/4.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; YComp 5.0.0.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 4.0; .NET CLR 1.0.2914)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 4.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98; YComp 5.0.0.0)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98; Win 9x 4.90)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows 98)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.1)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; .NET CLR 1.0.3705)",
		"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 4.0)",
	}
	expectedBrowserNames["Chrome"] = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.0; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_5_8) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.66 Safari/535.11",
	}

	expectedBrowserNames["PlayStation"] = []string{
		"Mozilla/5.0 (PlayStation 4 2.02) AppleWebKit/537.73 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3; 4.50) AppleWebKit/531.22.8 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3 4.60) AppleWebKit/531.22.8 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3; 3.55)",
		"Mozilla/5.0 (PlayStation 4 2.02) AppleWebKit/537.73 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3 4.55) AppleWebKit/531.22.8 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3 4.65) AppleWebKit/531.22.8 (KHTML, like Gecko)",
		"Mozilla/5.0 (PLAYSTATION 3 4.60) AppleWebKit/531.22.8 (KHTML, like Gecko)",
	}

	var expectedOperatingSystems map[string][]string = make(map[string][]string)
	expectedOperatingSystems["Playstation OS"] = expectedBrowserNames["Playstation"]

	_checkExepectations(t, expectedOperatingSystems, "os")
	_checkExepectations(t, expectedBrowserNames, "browser")
}

func _checkExepectations(t *testing.T, expectations map[string][]string, testType string) {
	var uaParseResult *UAInfo
	var testResult bool

	for expectation := range expectations {
		for key := range expectations[expectation] {
			fmt.Printf("Testing: %s \n", expectations[expectation][key])
			uaParseResult = Parse(expectations[expectation][key])

			if testType == "browser" {
				testResult = _checkBrowser(uaParseResult, expectation)
			} else if testType == "os" {
				testResult = _checkOs(uaParseResult, expectation)
			}

			if !testResult {
				t.Fatalf("Expected: %s got nil or different value on: %s", expectation, expectations[expectation][key])
			}
		}
	}
}

func _checkBrowser(uainfo *UAInfo, expectation string) bool {
	if uainfo.Browser == nil {
		return false
	}
	return uainfo.Browser.Name == expectation
}

func _checkOs(uainfo *UAInfo, expectation string) bool {
	if uainfo.OS == nil {
		return false
	}
	return uainfo.OS.Name == expectation
}
