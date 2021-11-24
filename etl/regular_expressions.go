package etl

import "regexp"

var (
	regexList = map[string]*regexp.Regexp{
		"date":      regexp.MustCompile(`\d\d/\d\d/\d\d\d\d`),
		"duration":  regexp.MustCompile(`\d\d:\d\d:\d\d`),
		"place":     regexp.MustCompile(`INAOE|UPTLAX`),
		"lightning": regexp.MustCompile(`Lightning Detected`),
		"distance":  regexp.MustCompile(`at\s\d\d\skm|at\s\d\skm`),
	}
	places = map[string]int{
		"INAOE":  1,
		"UPTLAX": 1,
	}
)
