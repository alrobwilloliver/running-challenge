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
		"Ran 2.3km today, feeling goood!",
		"I walked from the sofa to the kitchen.",
		"I ran .5km today, feeling good!",
		"I ran .5km today, 1m uphill - feeling good!",
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
	elevation := "0m"
	if regEl.MatchString(runDetails) {
		elevation = regEl.FindString(runDetails)
	}

	// distance
	regDist := regexp.MustCompile(`(\d?\.\d+)km|(\d?\.\d+)k`)
	distance := "0km"
	if regDist.MatchString(runDetails) {
		distance = regDist.FindString(runDetails)
	}

	return RunResult{
		activity:  activity,
		distance:  distance,
		elevation: elevation,
	}
}
