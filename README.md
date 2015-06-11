#uaparser#

HTTP User Agent parser in Golang

* browser part is mainly based on [Browser detection using the user agent | MDN](https://developer.mozilla.org/en-US/docs/Browser_detection_using_the_user_agent)
* deal with some weirdo chinese browsers that western guys need not to

##Example##

<pre><code>
package main

import(
    "github/varstr/uaparser"
    "fmt"
)

func main() {
    ua1 := "(Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16"
    rs1 := uaparser.Parse(ua1)

    ua2 := "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.33 Safari/534.3 SE 2.X MetaSr 1.0"
    rs2 := uaparser.Parse(ua2)

    fmt.Println("ua:", ua1,
                "\nbrowser:", rs1.Browser.Name, "version:", rs1.Browser.Version,
                "\ndevice:", rs1.Device.Name,
                "\nos:", rs1.OS.Name, "version:", rs1.OS.Version)

    fmt.Println()

    fmt.Println("ua:", ua2,
                "\nbrowser:", rs2.Browser.Name, "version:", rs2.Browser.Version,
                "\ndevice:", rs2.Device.Name,
                "\nos:", rs2.OS.Name, "version:", rs2.OS.Version)
}

/*
ua: (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16
browser: Safari version: 5.0
device: Macintosh
os: Mac OS version: X 10_6_3

ua: Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.33 Safari/534.3 SE 2.X MetaSr 1.0
browser: Sougou version: 2.X
device: PC
os: Windows version: NT 5.1
*/
</code></pre>

##License

[CC0](https://creativecommons.org/publicdomain/zero/1.0/)

