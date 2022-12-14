package please_test

import (
	"testing"

	"github.com/VJftw/please-nodejs/pkg/please"
	"github.com/stretchr/testify/assert"
)

func TestNewTargetFromReference(t *testing.T) {
	var tests = []struct {
		inRef     string
		outTarget *please.Target
		outErr    error
	}{
		{
			"//foo:bar",
			&please.Target{Pkg: "//foo", Name: "bar"},
			nil,
		},
		{
			"//foo/bar:baz",
			&please.Target{Pkg: "//foo/bar", Name: "baz"},
			nil,
		},
		{
			"//foo",
			&please.Target{Pkg: "//foo", Name: "foo"},
			nil,
		},
		{
			"//foo/bar",
			&please.Target{Pkg: "//foo/bar", Name: "bar"},
			nil,
		},
		{
			"//foo:bar:invalid",
			nil,
			please.ErrUnexpectedAmountOfColons,
		},
		{
			"foo/bar:missing_pkg_refix",
			nil,
			please.ErrMissingDoubleSlashPrefixFromPkg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.inRef, func(t *testing.T) {
			actualTarget, err := please.NewTargetFromReference(tt.inRef)
			assert.ErrorIs(t, err, tt.outErr)
			assert.Equal(t, tt.outTarget, actualTarget)
		})
	}
}

func TestTargetString(t *testing.T) {
	var tests = []struct {
		inTarget  *please.Target
		outString string
	}{
		{
			&please.Target{Pkg: "//foo", Name: "bar"},
			"//foo:bar",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "baz"},
			"//foo/bar:baz",
		},
		{
			&please.Target{Pkg: "//foo", Name: "foo"},
			"//foo:foo",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "bar"},
			"//foo/bar:bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.inTarget.String(), func(t *testing.T) {
			actualString := tt.inTarget.String()
			assert.Equal(t, tt.outString, actualString)
		})
	}
}

func TestTargetPkgDir(t *testing.T) {
	var tests = []struct {
		inTarget  *please.Target
		outPkgDir string
	}{
		{
			&please.Target{Pkg: "//foo", Name: "bar"},
			"foo",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "baz"},
			"foo/bar",
		},
		{
			&please.Target{Pkg: "//foo", Name: "foo"},
			"foo",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "bar"},
			"foo/bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.inTarget.String(), func(t *testing.T) {
			actualPkgDir := tt.inTarget.PkgDir()
			assert.Equal(t, tt.outPkgDir, actualPkgDir)
		})
	}
}

func TestTargetBuildFilePath(t *testing.T) {
	var tests = []struct {
		inTarget         *please.Target
		outBuildFilePath string
	}{
		{
			&please.Target{Pkg: "//foo", Name: "bar"},
			"foo/BUILD",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "baz"},
			"foo/bar/BUILD",
		},
		{
			&please.Target{Pkg: "//foo", Name: "foo"},
			"foo/BUILD",
		},
		{
			&please.Target{Pkg: "//foo/bar", Name: "bar"},
			"foo/bar/BUILD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.inTarget.String(), func(t *testing.T) {
			actualBuildFilePath := tt.inTarget.BuildFilePath()
			assert.Equal(t, tt.outBuildFilePath, actualBuildFilePath)
		})
	}
}

func TestTargetPrefixPkg(t *testing.T) {
	var tests = []struct {
		inTarget  *please.Target
		inPrefix  string
		outPrefix string
	}{
		{
			&please.Target{Pkg: "//foo", Name: "bar"},
			"//third_party/nodejs",
			"//third_party/nodejs/foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.inTarget.String(), func(t *testing.T) {
			tt.inTarget.PrefixPkg(tt.inPrefix)
			assert.Equal(t, tt.outPrefix, tt.inTarget.Pkg)
		})
	}
}
