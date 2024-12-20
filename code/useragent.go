package code

func DataUserAgent() []string {
	userAgents := []string{
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; de) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; de) Opera 8.02",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.02",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.53",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; pl) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; da) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.01",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.02",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 7.60",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.00",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.01",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.02",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.53",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.24",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.26",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; es-la) Opera 9.27",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; fr) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; IT) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; pl) Opera 8.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; pl) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ru) Opera 8.0",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ru) Opera 8.01",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ru) Opera 8.53",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ru) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ru) Opera 9.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; sv) Opera 8.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; sv) Opera 8.51",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; sv) Opera 8.53",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; tr) Opera 8.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; zh-cn) Opera 8.65",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; en) Opera 8.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; en) Opera 9.27",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; en) Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; ru) Opera 8.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 6.0; en) Opera 9.26",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 6.0; en) Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 6.0; tr) Opera 10.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; de) Opera 10.10",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 8.02",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 8.51",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 8.52",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 8.54",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 9.22",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 9.27",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; ru) Opera 8.51",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux x86_64; en) Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux x86_64; en) Opera 9.60",
		"Mozilla/4.0 (compatible; MSIE 8.0; Linux i686; en) Opera 10.51",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; ko) Opera 10.53",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; pl) Opera 11.00",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; en) Opera 11.00",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; en) Opera 9.50",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; en) Opera 9.52",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; en) Opera 9.60",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; ru) Opera 9.52",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; en) Opera 10.10",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; en) Opera 10.50",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; en) Opera 11.00",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; fr) Opera 9.80",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; ja) Opera 9.80",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; zh-cn) Opera 9.80",
		"Mozilla/4.0 (compatible; MSIE 8.0; X11; Linux x86_64; en) Opera 10.00",
		"Mozilla/4.0 (compatible; MSIE 8.0; X11; Linux x86_64; en) Opera 10.53",
		"Mozilla/4.0 (compatible; MSIE 8.0; X11; Linux x86_64; en) Opera 10.60",
		"Mozilla/4.0 (compatible; MSIE 8.0; X11; Linux x86_64; en) Opera 9.80",
		"Mozilla/4.0 (compatible; MSIE 8.0; X11; Linux x86_64; pl) Opera 10.10",
	}
	return userAgents
}
