package uaparser

var (
	desktop = &itemSpec{
		name:            "Desktop",
		mustContains:    []string{"Windows", "Linux", "Mac OS X", "CrOS", "Macintosh"},
		mustNotContains: []string{"Windows Phone", "Android", "ARM"},
	}

	tablet = &itemSpec{
		name:         "Tablet",
		mustContains: []string{"iPad", "Android", "ARM", "PlayBook"},
		mustNotContains: []string{
			"Mobile ",
			"C6802", // Xperia Z Android (which is a phone) does not include Mobile in UA-string so without this is seen as a tablet
			"Windows Phone",
			"model/Watch",
			"Ticwatch",
			"TicWatch",
		},
	}

	phone = &itemSpec{
		name: "Phone",
		mustContains: []string{
			"iPhone",
			"Android",
			"Windows Phone",
			"BB10",
			"BlackBerry",
			"nokia",
			"Nokia",
		},
		mustNotContains: []string{
			"model/Watch",
			"Ticwatch",
			"TicWatch",
		},
	}

	car = &itemSpec{
		name: "Car",
		mustContains: []string{
			"QtCarBrowser",
		},
		mustNotContains: []string{},
	}

	smartTv = &itemSpec{
		name: "SmartTV",
		mustContains: []string{
			"SMART-TV",
			"AppleTV",
			"CrKey",
			"Large Screen",
			"HbbTV",
			"LG Browser",
			"PhilipsTV",
			"Opera TV",
			"PlayStation",
		},
		mustNotContains: []string{},
	}

	watch = &itemSpec{
		name: "Watch",
		mustContains: []string{
			"watch",
			"Watch",
		},
	}

	_DEVICETYPES = []*itemSpec{
		tablet,
		phone,
		car,
		smartTv,
		desktop,
		watch,
	}
)
