package tree_sitter_awk_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_awk "github.com/givo29/tree-sitter-awk/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_awk.Language())
	if language == nil {
		t.Errorf("Error loading Awk grammar")
	}
}
