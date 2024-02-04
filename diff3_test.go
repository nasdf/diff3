package diff3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	textO  string
	textA  string
	textB  string
	expect string
}

var testCases = []TestCase{
	{
		textO:  "Aa\nBb\nCc\n",
		textA:  "Aa\nBb\nCc\nDd\nEe\nFf\n",
		textB:  "Aa\nBb\n",
		expect: "Aa\nBb\n<<<<<<<\nCc\nDd\nEe\nFf\n=======\n>>>>>>>\n",
	},
	{
		textO:  "celery\ngarlic\nonions\nsalmon\ntomatoes\nwine\n",
		textA:  "celery\nsalmon\ntomatoes\ngarlic\nonions\nwine\n",
		textB:  "celery\ngarlic\nsalmon\ntomatoes\nonions\nwine\n",
		expect: "celery\nsalmon\ntomatoes\ngarlic\n<<<<<<<\nonions\n=======\nsalmon\ntomatoes\nonions\n>>>>>>>\nwine\n",
	},
}

func TestMerge(t *testing.T) {
	for _, tc := range testCases {
		result := Merge(tc.textO, tc.textA, tc.textB)
		assert.Equal(t, tc.expect, result)
	}
}
