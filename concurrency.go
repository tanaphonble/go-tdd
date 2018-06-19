package main

// WebsiteChecker - type WebsiteChecker
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites -  Check websites status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChanel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChanel
		results[result.string] = result.bool
	}

	return results
}
