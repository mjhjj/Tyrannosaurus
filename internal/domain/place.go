package domain

type Place struct {
	Id          string `json:"id"`
	PositionX   string `json:"mmx"`
	PositionY   string `json:"mmy"`
	Name        string `json:"placeName"`
	Address     string `json:"placeAddress"`
	About       string `json:"placeAbout"`
	Bio         string `json:"placeBio"`
	Image       string `json:"placeImage"`
	Sity        string `json:"sityName"`
	PanoramLink string `json:"panoramLink"`
	LinkName    string `json:"linkButtonName"`
}
