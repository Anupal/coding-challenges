package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseString(t *testing.T) {
	testCases := []struct {
		testString string
		pass       bool
	}{
		{testString: `"hello"`, pass: true},
		{testString: `"hello-yellow"`, pass: true},
		{testString: `hello-yellow`, pass: false},
		{testString: `"hello-yellow`, pass: false},
		{testString: `"hello-yellow:`, pass: false},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.testString, func(t *testing.T) {
			endIndex, parsedString, err := ParseString(0, tc.testString)

			if tc.pass {
				require.NoError(t, err)
				require.Equal(t, len(parsedString), len(tc.testString)-2)
				require.Equal(t, endIndex, len(tc.testString))
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestParseObject(t *testing.T) {
	testCases := []struct {
		testString     string
		expectedObject Object
		pass           bool
	}{
		{
			testString: `{"hello1":"world1","hello2":"world2"}`,
			expectedObject: Object{
				"hello1": "world1",
				"hello2": "world2",
			},
			pass: true,
		},
		{testString: `"hello1":"world1","hello2":"world2"}`, pass: false},
		{testString: `{"hello1":"world1","hello2":"world2"`, pass: false},
		{testString: `"hello:"world1","hello2":"world2"}`, pass: false},
		{testString: `"hello","hello2":"world2"}`, pass: false},
		{testString: `{}`, expectedObject: Object{}, pass: true},
		{testString: `{"hello1":"world1""hello2":"world2"}`, pass: false},
		{testString: `{hello1:"world1","hello2":"world2"}`, pass: false},
		{testString: `"hello1""world1","hello2":"world2"}`, pass: false},
		{testString: `"hello1":"world1","hello2":`, pass: false},
		{testString: `"hello1":"world1}`, pass: false},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.testString, func(t *testing.T) {
			parsedObject, err := ParseObject(0, tc.testString)

			if tc.pass {
				require.NoError(t, err)
				require.Equal(t, parsedObject, tc.expectedObject)
			} else {
				require.Error(t, err)
			}
		})
	}
}
