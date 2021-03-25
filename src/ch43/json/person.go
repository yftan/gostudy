package json

type BaseInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Skills struct {
	Technologies []string `json:"technologies"`
}

type Person struct {
	Base   BaseInfo `json:"base"`
	Skills Skills   `json:"skills"`
}
