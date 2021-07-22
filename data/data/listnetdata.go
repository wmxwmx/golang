package data

import "time"

type ListNetData struct {
	Data struct {
		Locale struct {
			Tag struct {
				CacheKey            string      `json:"cacheKey"`
				ID                  string      `json:"id"`
				Slug                string      `json:"slug"`
				Avatar              string      `json:"avatar"`
				CreatedAt           time.Time   `json:"createdAt"`
				UpdatedAt           time.Time   `json:"updatedAt"`
				RedirectRelativeURL interface{} `json:"redirectRelativeUrl"`
				Alternates          []struct {
					CacheKey string `json:"cacheKey"`
					Short    string `json:"short"`
					Domain   string `json:"domain"`
					ID       string `json:"id"`
					Code     string `json:"code"`
					Typename string `json:"__typename"`
				} `json:"alternates"`
				TagTranslates []struct {
					CacheKey        string `json:"cacheKey"`
					ID              string `json:"id"`
					Title           string `json:"title"`
					MetaTitle       string `json:"metaTitle"`
					PageTitle       string `json:"pageTitle"`
					Description     string `json:"description"`
					MetaDescription string `json:"metaDescription"`
					Keywords        string `json:"keywords"`
					Typename        string `json:"__typename"`
				} `json:"tagTranslates"`
				Posts struct {
					Data []struct {
						CacheKey      string `json:"cacheKey"`
						ID            string `json:"id"`
						Slug          string `json:"slug"`
						Views         int    `json:"views"`
						PostTranslate struct {
							CacheKey             string    `json:"cacheKey"`
							ID                   string    `json:"id"`
							Title                string    `json:"title"`
							Avatar               string    `json:"avatar"`
							Published            time.Time `json:"published"`
							PublishedHumanFormat string    `json:"publishedHumanFormat"`
							LeadText             string    `json:"leadText"`
							Typename             string    `json:"__typename"`
						} `json:"postTranslate"`
						Category struct {
							CacheKey string `json:"cacheKey"`
							ID       string `json:"id"`
							Typename string `json:"__typename"`
						} `json:"category"`
						Author struct {
							CacheKey         string `json:"cacheKey"`
							ID               string `json:"id"`
							Slug             string `json:"slug"`
							AuthorTranslates []struct {
								CacheKey string `json:"cacheKey"`
								ID       string `json:"id"`
								Name     string `json:"name"`
								Typename string `json:"__typename"`
							} `json:"authorTranslates"`
							Typename string `json:"__typename"`
						} `json:"author"`
						PostBadge struct {
							CacheKey            string `json:"cacheKey"`
							ID                  string `json:"id"`
							Label               string `json:"label"`
							PostBadgeTranslates []struct {
								CacheKey string `json:"cacheKey"`
								ID       string `json:"id"`
								Title    string `json:"title"`
								Typename string `json:"__typename"`
							} `json:"postBadgeTranslates"`
							Typename string `json:"__typename"`
						} `json:"postBadge"`
						ShowShares bool   `json:"showShares"`
						ShowStats  bool   `json:"showStats"`
						Typename   string `json:"__typename"`
					} `json:"data"`
					PostsCount int    `json:"postsCount"`
					Typename   string `json:"__typename"`
				} `json:"posts"`
				Typename string `json:"__typename"`
			} `json:"tag"`
			Typename string `json:"__typename"`
		} `json:"locale"`
	} `json:"data"`
}
