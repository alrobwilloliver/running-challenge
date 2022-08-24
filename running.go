package main

import (
	"fmt"
	"regexp"
)

type RunResult struct {
	activity  string
	distance  string
	elevation string
}

func main() {
	sliceRunDetails := []string{
		"Last minute entry .. 3.8km 16m elevation",
		"Last one for me with some extra climbing over the railway bridge 4.72km with 21m of elevation",
		"Reluctant companion 4.41km walk 70m ascent",
		"Last little push!! Elevation wrong on strava photo. 5.07km and 55m elevation",
		"Last one from me - 4.3km 11m elevation",
		"Walking today - 7.1k with an elevation of 43m",
	}
	for _, result := range sliceRunDetails {
		fmt.Println(ExtractInfo(result))
	}
}

func ExtractInfo(runDetails string) RunResult {

	// activity
	activity := "Running"
	regActivity := regexp.MustCompile(`walk|Walk`)
	if regActivity.MatchString(runDetails) {
		activity = "Walking"
	}

	// elevation
	regEl := regexp.MustCompile(`\d+m`)
	elevation := regEl.FindString(runDetails)

	// distance
	regDist := regexp.MustCompile(`(\d+\.\d+)km|(\d+\.\d+)k`)
	distance := regDist.FindString(runDetails)

	return RunResult{
		activity:  activity,
		distance:  distance,
		elevation: elevation,
	}
}
