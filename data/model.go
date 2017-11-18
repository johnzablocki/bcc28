package data

type (
	//Model base for db models
	//This struct is meant to be embedded into other structs (see Book)
	//An embedded struct is created by having a nameless property on a struct
	//of some other type
	Model struct {
		ID string `json:"-"` //The - will cause ID not to be serialized as a JSON property
	}
)
