package musicTracker

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    Locations
	ConcertDates ConcertDates
	Relations    Relations
}

type Locations struct {
	Locations []string
}

type ConcertDates struct {
	Dates []string
}

type Relations struct {
	DatesLocations map[string][]string
}

type ArtistLight struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
}
