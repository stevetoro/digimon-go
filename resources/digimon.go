package resources

type Digimon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	XAntibody   bool   `json:"xAntibody"`
	ReleaseDate string `json:"releaseDate"`
	Images      []struct {
		Href        string `json:"href"`
		Transparent bool   `json:"transparent"`
	} `json:"images"`
	Levels []struct {
		ID    int    `json:"id"`
		Level string `json:"level"`
	} `json:"levels"`
	Types []struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"types"`
	Attributes []struct {
		ID        int    `json:"id"`
		Attribute string `json:"attribute"`
	} `json:"attributes"`
	Fields []struct {
		ID    int    `json:"id"`
		Field string `json:"field"`
		Image string `json:"image"`
	} `json:"fields"`
	Descriptions []struct {
		Origin      string `json:"origin"`
		Description string `json:"description"`
		Language    string `json:"language"`
	} `json:"descriptions"`
	Skills []struct {
		ID          int    `json:"id"`
		Skill       string `json:"skill"`
		Description string `json:"description"`
		Translation string `json:"translation"`
	} `json:"skills"`
	PriorEvolutions []struct {
		ID        int    `json:"id"`
		Digimon   string `json:"digimon"`
		Condition string `json:"condition"`
		Image     string `json:"image"`
		URL       string `json:"url"`
	} `json:"priorEvolutions"`
	NextEvolutions []struct {
		ID        int    `json:"id"`
		Digimon   string `json:"digimon"`
		Condition string `json:"condition"`
		Image     string `json:"image"`
		URL       string `json:"url"`
	} `json:"nextEvolutions"`
}

type DigimonPage struct {
	Content []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Href  string `json:"href"`
		Image string `json:"image"`
	} `json:"content"`
	Pageable Pageable `json:"pageable"`
}

type DigimonQueryParams struct {
	Name      string `url:"name"`
	Attribute string `url:"attribute"`
	XAntibody string `url:"xAntibody"`
	Level     string `url:"level"`
	Page      int    `url:"page"`
	PageSize  int    `url:"pageSize"`
}
