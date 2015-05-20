//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func TestBowlingGameScoringTests(t *testing.T) {
	fixture := gunit.NewFixture(t)

	test0 := &BowlingGameScoringTests{Fixture: fixture}
	test0.RunTestCase__(test0.TestAfterAllGutterBallsTheScoreShouldBeZero, "After all gutter balls the score should be zero")

	test1 := &BowlingGameScoringTests{Fixture: fixture}
	test1.RunTestCase__(test1.TestAfterAllOnesTheScoreShouldBeTwenty, "After all ones the score should be twenty")

	test2 := &BowlingGameScoringTests{Fixture: fixture}
	test2.RunTestCase__(test2.TestSpareReceivesSingleRollBonus, "Spare receives single roll bonus")

	test3 := &BowlingGameScoringTests{Fixture: fixture}
	test3.RunTestCase__(test3.TestStrikeRecievesDoubleRollBonus, "Strike recieves double roll bonus")

	test4 := &BowlingGameScoringTests{Fixture: fixture}
	test4.RunTestCase__(test4.TestPerfectGame, "Perfect game")

	fixture.Finalize()
}

func (self *BowlingGameScoringTests) RunTestCase__(test func(), description string) {
	self.Describe(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func TestEnvironmentControllerFixture(t *testing.T) {
	fixture := gunit.NewFixture(t)

	test0 := &EnvironmentControllerFixture{Fixture: fixture}
	test0.RunTestCase__(test0.TestEverythingTurnedOffAtStartup, "Everything turned off at startup")

	test1 := &EnvironmentControllerFixture{Fixture: fixture}
	test1.RunTestCase__(test1.TestEverythingOffWhenComfortable, "Everything off when comfortable")

	test2 := &EnvironmentControllerFixture{Fixture: fixture}
	test2.RunTestCase__(test2.TestCoolerAndBlowerWhenHot, "Cooler and blower when hot")

	test3 := &EnvironmentControllerFixture{Fixture: fixture}
	test3.RunTestCase__(test3.TestHeaterAndBlowerWhenCold, "Heater and blower when cold")

	test4 := &EnvironmentControllerFixture{Fixture: fixture}
	test4.RunTestCase__(test4.TestHighAlarmOnIfAtThreshold, "High alarm on if at threshold")

	test5 := &EnvironmentControllerFixture{Fixture: fixture}
	test5.RunTestCase__(test5.TestLowAlarmOnIfAtThreshold, "Low alarm on if at threshold")

	fixture.Finalize()
}

func (self *EnvironmentControllerFixture) RunTestCase__(test func(), description string) {
	self.Describe(description)
	self.Setup()
	test()
}

///////////////////////////////////////////////////////////////////////////////

func init() {
	gunit.Validate("0a085d3a3201a50b682d18a273aa68de")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
