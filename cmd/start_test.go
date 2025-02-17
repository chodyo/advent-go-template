package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type YearDayValidationTestCase struct {
	name         string
	opts         Start
	now          time.Time
	expectedYear int
	expectedDay  int
	expectedErr  bool
}

func TestValidateYearAndDay(t *testing.T) {
	testCases := map[string][]YearDayValidationTestCase{
		"No Params Specified": {{
			name:         "During December",
			opts:         Start{},
			now:          time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC),
			expectedYear: 2024,
			expectedDay:  1,
			expectedErr:  false,
		}, {
			name:         "Outside December",
			opts:         Start{},
			now:          time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedYear: 2023,
			expectedDay:  1,
			expectedErr:  false,
		}},
		"Year Specified": {{
			name:         "During December",
			opts:         Start{Year: 2024},
			now:          time.Date(2024, time.December, 20, 0, 0, 0, 0, time.UTC),
			expectedYear: 2024,
			expectedDay:  1,
			expectedErr:  false,
		}, {
			name:         "Outside December",
			opts:         Start{Year: 2024},
			now:          time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedYear: 2024,
			expectedDay:  1,
			expectedErr:  false,
		}},
		"Day Specified": {{
			name:         "During December",
			opts:         Start{Day: 5},
			now:          time.Date(2023, time.December, 15, 0, 0, 0, 0, time.UTC),
			expectedYear: 2023,
			expectedDay:  5,
			expectedErr:  false,
		}, {
			name:         "Outside December",
			opts:         Start{Day: 5},
			now:          time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedYear: 2022,
			expectedDay:  5,
			expectedErr:  false,
		}},
		"Both Year and Day Specified": {{
			name:         "Specific year and day",
			opts:         Start{Year: 2022, Day: 10},
			now:          time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedYear: 2022,
			expectedDay:  10,
			expectedErr:  false,
		}},

		// No good way to test this case without overhauling the isSolutionExists function
		// "All Solutions Exist": {{
		// 	name:         "No available days",
		// 	opts:         Start{},
		// 	now:          time.Date(2023, time.December, 15, 0, 0, 0, 0, time.UTC),
		// 	expectedYear: 0,
		// 	expectedDay:  0,
		// 	expectedErr:  true,
		// }},
	}

	// Iterate through test case groups
	for groupName, group := range testCases {
		t.Run(groupName, func(t *testing.T) {
			for _, tc := range group {
				t.Run(tc.name, func(t *testing.T) {
					year, day, err := validateOrGetDefaultYearAndDay(tc.opts, tc.now)
					assert.Equal(t, tc.expectedYear, year, "Year should match expected value")
					assert.Equal(t, tc.expectedDay, day, "Day should match expected value")
					assert.Equal(t, tc.expectedErr, err != nil, "Error expectation should match")
				})
			}
		})
	}
}
