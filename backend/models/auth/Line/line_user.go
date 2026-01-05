package line_models

type LineUser struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureURL  string `json:"pictureUrl"`
	StatusMsg   string `json:"statusMessage"`
}