package uaparser

var (
	iPad = &itemSpec{
		name:         "iPad",
		mustContains: []string{"iPad"},
	}

	iPhone = &itemSpec{
		name:         "iPhone",
		mustContains: []string{"iPhone"},
	}

	iPod = &itemSpec{
		name:         "iPod",
		mustContains: []string{"iPod"},
	}

	mac = &itemSpec{
		name:         "Macintosh",
		mustContains: []string{"Macintosh"},
	}

	_DEVICES = []*itemSpec{
		iPad,
		iPhone,
		iPod,
		mac,
	}
)
