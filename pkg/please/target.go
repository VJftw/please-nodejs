package please

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

var (
	// ErrMissingDoubleSlashPrefixFromPkg is returned when '//' is missing from
	// the prefix of a package
	ErrMissingDoubleSlashPrefixFromPkg = errors.New("missing '//' prefix from pkg")

	// ErrUnexpectedAmountOfColons is returned when there is an unexpected
	// amount of colons in the ref.
	ErrUnexpectedAmountOfColons = errors.New("unexpected amound of colons")
)

// Target represents a Please Build Target.
type Target struct {
	Pkg  string
	Name string
}

// NewTargetFromReference returns a new Target from the given Please Target
// Reference e.g. '//my_pkg:my_rule'.
func NewTargetFromReference(ref string) (*Target, error) {
	t := &Target{}
	parts := strings.Split(ref, ":")
	switch len(parts) {
	case 1:
		t.Pkg = ref
		t.Name = path.Base(ref)
	case 2:
		t.Pkg = parts[0]
		t.Name = parts[1]
	default:
		return nil, fmt.Errorf("could not parse '%s': %w", ref, ErrUnexpectedAmountOfColons)
	}

	if !strings.HasPrefix(t.Pkg, "//") {
		return nil, fmt.Errorf("could not parse '%s': %w", ref, ErrMissingDoubleSlashPrefixFromPkg)
	}

	return t, nil
}

// String translates a target to a Please Build Target reference in the form of
// `//<pkg>:<name>`.
func (t *Target) String() string {
	return fmt.Sprintf("%s:%s", t.Pkg, t.Name)
}

// PkgDir returns the relative directory of the Pkg.
func (t *Target) PkgDir() string {
	return filepath.FromSlash(RelPkg(t.Pkg))
}

// BuildFilePath returns the relative path to the BUILD file.
func (t *Target) BuildFilePath() string {
	return filepath.Join(t.PkgDir(), "BUILD")
}

// PrefixPkg inserts the given prefix into the current pkg, effectively moving
// the target into a base subfolder.
func (t *Target) PrefixPkg(prefix string) {
	t.Pkg = fmt.Sprintf("//%s", path.Join(RelPkg(prefix), RelPkg(t.Pkg)))
}

// RelPkg returns the given Pkg with all leading `/` removed.
func RelPkg(pkg string) string {
	return strings.TrimLeft(pkg, "/")
}
