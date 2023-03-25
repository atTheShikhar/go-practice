package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, websites []string) (result map[string]bool) {
	result = map[string]bool{}

	for _, website := range websites {
		result[website] = wc(website)
	}

	return
}
