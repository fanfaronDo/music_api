package entry

type Song struct {
	Id          string `json:"id"`
	Group       string `json:"group"`
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Patronymic  string `json:"patronymic"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
