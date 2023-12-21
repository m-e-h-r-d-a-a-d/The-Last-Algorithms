package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoCrystalBalls(t *testing.T) {
	testCases := []struct {
		name   string
		breaks []bool
		except int
	}{
		{
			name:   "test 1",
			breaks: []bool{true, true, true, true, true, true, true, true, true},
			except: 0,
		},
		{
			name:   "test 2",
			breaks: []bool{false, false, false, false, true, true, true, true},
			except: 4,
		},
		{
			name:   "test 3",
			breaks: []bool{true},
			except: 0,
		},
		{
			name:   "test 4",
			breaks: []bool{false},
			except: -1,
		},
		{
			name:   "test 5",
			breaks: []bool{},
			except: -1,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			result := TwoCrystalBalls(tc.breaks)
			require.Equal(t, result, tc.except)
		})
	}
}
