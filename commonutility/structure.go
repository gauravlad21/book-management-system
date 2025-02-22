package commonutility

type ShortenerResponse struct {
	StatusCode int64  `json:"status,omitempty"`
	ShortUrl   string `json:"shortUrl,omitempty"`
	Msg        string `json:"message,omitempty"`
}

type RedirectResponse struct {
	StatusCode int64  `json:"status,omitempty"`
	LongUrl    string `json:"longUrl,omitempty"`
	Msg        string `json:"message,omitempty"`
}

type TopKResult struct {
	StatusCode int64 `json:"status,omitempty"`
	Details    []*NameCount
	Msg        string `json:"message,omitempty"`
}

type Req struct {
	Url string `json:"url,omitempty"`
}

type NameCount struct {
	Name  string
	Count int
}
