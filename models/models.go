package models

type RedditResponse struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
}
type Source struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Image struct {
	OriginalSource Source `json:"source"`
}
type Data struct {
	Url_overridden_by_dest string  `json:"url_overridden_by_dest"`
	Preview                []Image `json:"preview"`
}

type Listing struct {
	Kind string      `json:"kind"`
	Data ListingData `json:"data"`
}
type ListingData struct {
	After    string            `json:"after"`
	Children []ListingChildren `json:"children"`
}
type ListingChildren struct {
	Kind string              `json:"kind"`
	Data ListingChildrenData `json:"data"`
}
type ListingChildrenData struct {
	UrlOverriddenByDest string         `json:"url_overridden_by_dest"`
	Preview             ListingPreview `json:"preview"`
}
type ListingPreview struct {
	Images []ListingImage `json:"images"`
}
type ListingImage struct {
	Source ListingImageSource `json:"source"`
}
type ListingImageSource struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
