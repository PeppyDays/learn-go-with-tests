package c10

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	ch := make(chan (checkResult))
	for i := 0; i < len(urls); i++ {
		go func(url string) {
			ch <- checkResult{url, wc(url)}
		}(urls[i])
	}

	results := make(map[string]bool, len(urls))
	for i := 0; i < len(urls); i++ {
		result := <-ch
		results[result.string] = result.bool
	}
	return results
}

type checkResult struct {
	string
	bool
}
