package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evansb/snapper/cmd"
)

func TestSnapper(t *testing.T) {
	t.Run("RunWithoutArguments", func(t *testing.T) {
		app := cmd.Snapper()
		assert.NotNil(t, app)
		assert.NotPanics(t, func() {
			app.Run([]string{"snapper"})
		})
	})
}
