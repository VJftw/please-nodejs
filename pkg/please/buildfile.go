package please

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/edit"
	"github.com/rs/zerolog/log"
)

// BuildFile represents a Please BUILD file which wraps the underlying bazel
// *build.File implementation.
type BuildFile struct {
	*build.File
}

func NewBuildFile(path string) (*BuildFile, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("could not determine absolute path of '%s': %w", path, err)
	}

	return &BuildFile{
		File: &build.File{
			Path: absPath,
			Pkg:  "//" + filepath.Dir(path),
		},
	}, nil
}

func LoadBuildFileFromFile(path string) (*BuildFile, error) {
	buildFileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file contents '%s': %w", path, err)
	}

	parsedBuildFile, err := build.Parse(path, buildFileContents)
	if err != nil {
		return nil, fmt.Errorf("could not parse BUILD file '%s': %w", path, err)
	}

	return &BuildFile{
		File: parsedBuildFile,
	}, nil
}

func (bf *BuildFile) HasSubinclude(target string) bool {
	subincludes := bf.Rules("subinclude")

	for _, subinclude := range subincludes {
		for _, attrKey := range subinclude.Call.List {
			switch v := attrKey.(type) {
			case *build.StringExpr:
				if v.Value == target {
					return true
				}
			case *build.ListExpr:
				for _, item := range v.List {
					if strVal, ok := item.(*build.StringExpr); ok {
						if strVal.Value == target {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

func (bf *BuildFile) EnsureSubinclude(target string) error {
	if bf.HasSubinclude(target) {
		return nil
	}

	subincludes := bf.Rules("subinclude")
	subincludeRule, _ := edit.ExprToRule(&build.CallExpr{
		X:    &build.Ident{Name: "subinclude"},
		List: []build.Expr{},
	}, "subinclude")

	if len(subincludes) > 0 {
		subincludeRule = subincludes[0]
	}

	subincludeRule.Call.List = append(subincludeRule.Call.List, &build.StringExpr{Value: target})

	if len(subincludes) < 1 {
		bf.Stmt = append([]build.Expr{subincludeRule.Call}, bf.Stmt...)
	}

	// TODO: figure out how to update rule, maybe it's fine as it's a pointer?

	return nil
}

func (bf *BuildFile) UpsertRule(rule *Rule) error {
	bf.DelRules("", rule.Name())
	bf.Stmt = append(bf.Stmt, rule.Call)

	return nil
}

func (bf *BuildFile) Save() error {
	formattedBytes := build.Format(bf.File)

	if err := os.MkdirAll(filepath.Dir(bf.Path), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(bf.Path, formattedBytes, 0660); err != nil {
		return err
	}

	log.Debug().Str("path", bf.Path).Msg("saved BUILD file")

	return nil
}
