package dtypes

type TodoEntry struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  *string  `json:"description"`
	Done         bool     `json:"done"`
	CreationDate string   `json:"creationDate"`
	RelevantDate *string  `json:"relevantDate"`
	Link         *string  `json:"link"`
	Labels       *[]Label `json:"labels"`
}

type Label struct {
	Name  string `json:"labelName"`
	Color string `json:"labelColor"`
}
