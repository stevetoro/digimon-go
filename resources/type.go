package resources

type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TypePage struct {
	Content struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Fields      []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"content"`
	Pageable Pageable `json:"pageable"`
}
