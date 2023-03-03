package musicTracker

import (
	"musicTracker/musicTracker"
)

func tri(localise []musicTracker.Locations, Localised string) []string {
	var fill []string

	//i := musicTracker.GetAllArtists()

	for _, dom := range localise {
		for i := 0; i < len(dom.Locations); i++ {
			if dom.Locations[i] == Localised {
				fill = append(fill, dom.Locations[i])
			}

		}
	}
	return fill
}
