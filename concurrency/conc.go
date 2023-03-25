package main

type WebsiteChecker func(string) bool

type result struct {
	website string
	result  bool
}

func CheckWebsites(wc WebsiteChecker, websites []string) (results map[string]bool) {
	results = map[string]bool{}
	resultChannel := make(chan result)

	for _, website := range websites {
		go func(w string) {
			resultChannel <- result{website: w, result: wc(w)}
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChannel
		results[r.website] = r.result
	}
	return
}
