package domain

type Url struct {
	ID      string
	Address string
}

const (
	ID = iota
	Address
)

type UrlCollection []Url
