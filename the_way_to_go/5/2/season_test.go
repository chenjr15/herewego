package season_test

import "testing"

func TestSeason(t *testing.T) {

	for index := 0; index < 15; index++ {
		t.Logf("The season of %d is %s .", index, Season(index))
	}

}

// Season 返回指定月份的季节
func Season(season int) (SeasonString string) {
	switch season {
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		SeasonString = "Spring"
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		SeasonString = "Summer"
	case 9:
		fallthrough
	case 10:
		fallthrough
	case 11:
		SeasonString = "Fall"
	case 12:
		fallthrough
	case 1:
		fallthrough
	case 2:
		SeasonString = "Winter"
	default:
		SeasonString = "Unknown"
	}

	return
}
