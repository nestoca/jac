package printing

type PrintFormat int

const (
	FormatTable PrintFormat = iota // Default
	FormatTree
	FormatYAML
)

func NewPrintOpts(formatTree, formatYaml, showAll, hideGroupColumns, showIdentifierNames, highlightMatches bool) *PrintOpts {
	opts := PrintOpts{
		Format:              FormatTable,
		ShowAll:             showAll,
		HideGroupColumns:    hideGroupColumns,
		ShowIdentifierNames: showIdentifierNames,
		HighlightMatches:    highlightMatches,
	}
	if formatTree {
		if formatYaml {
			panic("cannot specify both --tree and --yaml")
		}
		opts.Format = FormatTree
	} else if formatYaml {
		opts.Format = FormatYAML
	}
	return &opts
}

type PrintOpts struct {
	// The format to use when printing the output.
	Format PrintFormat

	// Whether to show all entries, including those that are not matched by the
	// filters, but highlighting the matched ones.
	ShowAll bool

	// Whether to hide the group columns in the table format.
	HideGroupColumns bool

	// Whether to show only the identifier names instead of the full names.
	ShowIdentifierNames bool

	// Whether to highlight the matches in the output (when ShowAll is true and
	// some filters have been applied).
	HighlightMatches bool
}
