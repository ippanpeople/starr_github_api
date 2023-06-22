package git

import (
	"math"
)

type Repository struct {
	Name map[string]int
}

func User_lang() map[string]float64 {
	totalCount := 0
	username := ""
	stocks := map[string]float64{}
	//repos.goからユーザのリポジトリを取得
	re := Repo_search(username)
	// fmt.Println(re)

	for _, repo := range re {
		url := "https://api.github.com/repos/" + username + "/" + repo.Name + "/languages"
		// var c []Repository
		//リクエスト作成

		languag := map[string]int{}
		Http_request(url, &languag)

		for languages, count := range languag {
			totalCount += count
			// fmt.Println(count)
			_, ok := stocks[languages]
			if !ok {
				stocks[languages] = float64(count)
			} else {
				stocks[languages] += float64(count)
			}
		}
	}
	// // fmt.Println(totalCount)
	for languages, count := range stocks {
		percentage := math.Round((float64(count)/float64(totalCount)*100)*10) / 10
		stocks[languages] = percentage
	}

	return stocks
}
