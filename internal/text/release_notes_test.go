package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseNotes(t *testing.T) {
	expected := "\n\n## Features :rocket:\n\n- 0000000 ci test\n\n## Bug fixes :bug:\n\n- 0000000 huge bug\n\n## Chores and Improvements :wrench:\n\n- 0000000 testing\n- 0000000 this should end up in chores\n\n## Other :package:\n\n- 0000000 merge master in something\n- 0000000 random\n\n"

	sections := Sections{
		Features: []Commit{Commit{Category: "feat", Scope: "ci", Heading: "ci test"}},
		Chores:   []Commit{Commit{Category: "chore", Scope: "", Heading: "testing"}, Commit{Category: "improvement", Scope: "", Heading: "this should end up in chores"}},
		Bugs:     []Commit{Commit{Category: "bug", Scope: "", Heading: "huge bug"}},
		Others:   []Commit{Commit{Category: "other", Scope: "", Heading: "merge master in something"}, Commit{Category: "bs", Scope: "", Heading: "random"}},
	}

	releaseNotes := ReleaseNotes(sections)

	assert.Equal(t, expected, releaseNotes)
}

func TestReleaseNotesWithMissingSections(t *testing.T) {
	expected := "\n\n## Features :rocket:\n\n- 0000000 ci test\n\n"

	sections := Sections{
		Features: []Commit{Commit{Heading: "ci test"}},
	}

	releaseNotes := ReleaseNotes(sections)

	assert.Equal(t, expected, releaseNotes)
}
