package data

type (
	//Book model struct
	Book struct {
		Model `json:",inline"` //inline tells the serializer to nest Model props at the root, not under Model
		Title string           `json:"title"` //tags are like attributes in languages like C#
		ISBN  string           `json:"isbn"`
	}
)
