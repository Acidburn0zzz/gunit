//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func Test_ExampleFixture__with_assertions(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &ExampleFixture{Fixture: fixture}
	defer test.TeardownStuff()
	test.SetupStuff()
	test.TestWithAssertions()
}

func Test_ExampleFixture__skip_with_nothing(t *testing.T) {
	t.Skip("Skipping test case: 'SkipTestWithNothing'")

	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &ExampleFixture{Fixture: fixture}
	defer test.TeardownStuff()
	test.SetupStuff()
	test.SkipTestWithNothing()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("3cacd4d2ce8c2490d88e63233d3049bf")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
