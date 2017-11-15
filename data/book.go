package data

type (
	//Book model
	Book struct {
		Model `json:",inline"`
		Title string `json:"title"`
		ISBN  string `json:"isbn"`
	}
)
