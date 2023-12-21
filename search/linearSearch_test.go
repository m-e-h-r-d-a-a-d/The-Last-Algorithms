package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []int
		needle   int
		except   bool
	}{
		{
			name:     "exist 1",
			haystack: []int{2, 12, 15, 1, -60, 133, 12, -200},
			needle:   -200,
			except:   true,
		},
		{
			name:     "exist 2",
			haystack: []int{15, 11550, 111, -23, 232, 0, 22, -200},
			needle:   11550,
			except:   true,
		},
		{
			name:     "not exist 1",
			haystack: []int{2, 12, 15, 1, -60, 133, 12, -200},
			needle:   0,
			except:   false,
		},
		{
			name:     "not exist 2",
			haystack: []int{2, 12, 15, 23, -60, 133, 12, 0},
			needle:   11550,
			except:   false,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			result := linearSearch(tc.haystack, tc.needle)
			require.Equal(t, result, tc.except)
		})
	}
}
