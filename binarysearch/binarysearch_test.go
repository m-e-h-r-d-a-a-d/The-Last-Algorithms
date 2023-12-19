package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		haystack []int
		needle   int
		except   bool
	}{
		{
			name:     "exist 1",
			haystack: []int{1, 12, 15, 55, 60, 133, 1112, 2304},
			needle:   15,
			except:   true,
		},
		{
			name:     "exist 2",
			haystack: []int{-100, -45, -34, -23, -1, 0, 22, 250, 2345},
			needle:   -1,
			except:   true,
		},
		{
			name:     "not exist 1",
			haystack: []int{1, 12, 15, 55, 60, 133, 1112, 2304},
			needle:   -10,
			except:   false,
		},
		{
			name:     "not exist 2",
			haystack: []int{-100, -45, -34, -23, -1, 14, 22, 250, 2345},
			needle:   0,
			except:   false,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.haystack, tc.needle)
			require.Equal(t, result, tc.except)
		})
	}
}
