package ckit

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segmentio/encoding/json"
)

func TestBatchedStrings(t *testing.T) {
	var cases = []struct {
		desc     string
		s        []string
		n        int
		expected [][]string
	}{
		{
			"empty slice", []string{}, 100, [][]string{
				[]string{},
			},
		},
		{
			"single batch", []string{"a", "b", "c"}, 100, [][]string{
				[]string{"a", "b", "c"},
			},
		},
		{
			"two batches", []string{"a", "b", "c"}, 2, [][]string{
				[]string{"a", "b"}, []string{"c"},
			},
		},
		{
			"full batches", []string{"a", "b", "c", "d"}, 2, [][]string{
				[]string{"a", "b"}, []string{"c", "d"},
			},
		},
	}
	for _, c := range cases {
		result := batchedStrings(c.s, c.n)
		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("{%s] got %v, want %v", c.desc, result, c.expected)
		}
	}
}

func TestApplyInstitutionFilter(t *testing.T) {
	var cases = []struct {
		desc        string
		institution string
		resp        []byte // use serialized form for simplicity
		expected    []byte
	}{
		{
			desc:        "empty",
			institution: "",
			resp:        []byte("{}"),
			expected:    []byte("{}"),
		},
		{
			desc:        "empty",
			institution: "any",
			resp:        []byte("{}"),
			expected:    []byte(`{"extra": {"institution": "any"}}`),
		},
		// TODO: add more cases
	}
	for _, c := range cases {
		var (
			resp     Response
			expected Response
		)
		if err := json.Unmarshal(c.resp, &resp); err != nil {
			t.Fatalf("could not unmarshal test response")
		}
		if err := json.Unmarshal(c.expected, &expected); err != nil {
			t.Fatalf("could not unmarshal test response")
		}
		resp.applyInstitutionFilter(c.institution)
		if !cmp.Equal(resp, expected) {
			t.Fatalf("[%s] got %v, want %v", c.desc, resp, expected)
		}
	}
}
