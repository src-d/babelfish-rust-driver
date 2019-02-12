package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/rust-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/driver"
	"gopkg.in/bblfsh/sdk.v2/driver/fixtures"
	"gopkg.in/bblfsh/sdk.v2/driver/native"
	"gopkg.in/bblfsh/sdk.v2/uast/transformer/positioner"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "rust",
	Ext:  ".ext", // TODO: specify correct file extension for source files in ./fixtures
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return native.NewDriverAt(filepath.Join(projectRoot, "build/bin/native"), native.UTF8)
	},
	Transforms: normalizer.Transforms,
	//BenchName: "fixture-name", // TODO: specify a largest file
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			// TODO: list native types that should be converted to semantic UAST
		},
	},
	VerifyTokens: []positioner.VerifyToken{
	    // TODO: list nodes that needs to be checked for token correctness
	},
}

func TestRustDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkRustDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
