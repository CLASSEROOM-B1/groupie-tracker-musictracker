package musicTracker

type Artist struct {
	Id                 int
	Image              string
	Name               string
	Members            []string
	CreationDate       int
	FirstAlbum         string
	LocationsStruct    Locations
	ConcertDatesStruct ConcertDates
	RelationsStruct    Relations
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
	Id              int
	Image           string
	Name            string
	Members         []string
	CreationDate    int
	FirstAlbum      string
	LocationsStruct Locations
}
