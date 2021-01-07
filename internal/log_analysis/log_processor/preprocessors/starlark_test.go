package preprocessors

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStarlark(t *testing.T) {
	config := StarlarkConfig{
		Code: `
def parse(entry):
	event = json.loads(entry)
	event['keys'] = [key for key in event]
	return event
`,
	}
	pp, err := config.BuildPreprocessor()
	assert := require.New(t)
	assert.NoError(err)
	out, err := pp.PreProcessLog(`{"foo":"bar"}`)
	assert.NoError(err)
	assert.Equal(`{"foo":"bar","keys":["foo"]}`, out)
}
