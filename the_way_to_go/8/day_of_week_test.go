package maps

import "testing"

func TestDayOfWeekMap(t *testing.T) {
	// use map as set
	daysMap := map[string]bool{
		"Sunsay":    true,
		"Monday":    true,
		"Tuesday":   true,
		"Wednesday": true,
		"Thursday":  true,
		"Friday":    true,
	}
	testcase := []struct {
		day    string
		exists bool
	}{
		{"Tuesday", true},
		{"Hollyday", false},
	}
	for _, s := range testcase {
		ok := daysMap[s.day]
		if ok == s.exists {

			t.Log("Passed", s.day)
		} else {
			t.Error("Failed", s.day)
		}
	}

}
