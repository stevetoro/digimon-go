package resources

type Pageable struct {
	CurrentPage    int    `json:"currentPage"`
	ElementsOnPage int    `json:"elementsOnPage"`
	TotalElements  int    `json:"totalElements"`
	TotalPages     int    `json:"totalPages"`
	PreviousPage   string `json:"previousPage"`
	NextPage       string `json:"nextPage"`
}
