package model

//  Metadata defines movie metadata
type Metadata struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Director string `json:"director"`
}