package walletobject

type TypedValue struct {
	Bool   bool         `json:"bool,omitempty"`
	Double float64      `json:"double,omitempty"`
	Image  *Image       `json:"image,omitempty"`
	Int    int          `json:"int,omitempty"`
	Kind   string       `json:"kind,omitempty"`
	List   []TypedValue `json:"list,omitempty"`
	String string       `json:"string,omitempty"`
	URI    *URI         `json:"uri,omitempty"`
}
