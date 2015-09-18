//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package examples

import (
	"testing"

	"github.com/smartystreets/gunit"
)

///////////////////////////////////////////////////////////////////////////////

func Test_BowlingGameScoringTests__after_all_gutter_balls_the_score_should_be_zero(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["bowling_game_test.go"])
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllGutterBallsTheScoreShouldBeZero()
}

func Test_BowlingGameScoringTests__after_all_ones_the_score_should_be_twenty(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["bowling_game_test.go"])
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestAfterAllOnesTheScoreShouldBeTwenty()
}

func Test_BowlingGameScoringTests__spare_receives_single_roll_bonus(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["bowling_game_test.go"])
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestSpareReceivesSingleRollBonus()
}

func Test_BowlingGameScoringTests__strike_recieves_double_roll_bonus(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["bowling_game_test.go"])
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestStrikeRecievesDoubleRollBonus()
}

func Test_BowlingGameScoringTests__perfect_game(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["bowling_game_test.go"])
	defer fixture.Finalize()
	test := &BowlingGameScoringTests{Fixture: fixture}
	test.Setup()
	test.TestPerfectGame()
}

///////////////////////////////////////////////////////////////////////////////

func Test_EnvironmentControllerFixture__everything_turned_off_at_startup(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingTurnedOffAtStartup()
}

func Test_EnvironmentControllerFixture__everything_off_when_comfortable(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestEverythingOffWhenComfortable()
}

func Test_EnvironmentControllerFixture__cooler_and_blower_when_hot(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestCoolerAndBlowerWhenHot()
}

func Test_EnvironmentControllerFixture__heater_and_blower_when_cold(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHeaterAndBlowerWhenCold()
}

func Test_EnvironmentControllerFixture__high_alarm_on_if_at_threshold(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestHighAlarmOnIfAtThreshold()
}

func Test_EnvironmentControllerFixture__low_alarm_on_if_at_threshold(t *testing.T) {
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose(), __code__["environment_controller_test.go"])
	defer fixture.Finalize()
	test := &EnvironmentControllerFixture{Fixture: fixture}
	test.Setup()
	test.TestLowAlarmOnIfAtThreshold()
}

///////////////////////////////////////////////////////////////////////////////

var __code__ = map[string]string{"bowling_game_test.go": "7061636b616765206578616d706c65730a0a2f2f676f3a67656e65726174652067756e69740a0a696d706f727420280a09226769746875622e636f6d2f736d61727479737472656574732f617373657274696f6e732f73686f756c64220a09226769746875622e636f6d2f736d61727479737472656574732f67756e6974220a290a0a7479706520426f776c696e6747616d6553636f72696e67546573747320737472756374207b0a092a67756e69742e466978747572650a0a0967616d65202a47616d650a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e675465737473292053657475702829207b0a0973656c662e67616d65203d204e657747616d6528290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920546573744166746572416c6c47757474657242616c6c7354686553636f726553686f756c6442655a65726f2829207b0a0973656c662e726f6c6c4d616e792832302c2030290a0973656c662e536f2873656c662e67616d652e53636f726528292c2073686f756c642e457175616c2c2030290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920546573744166746572416c6c4f6e657354686553636f726553686f756c6442655477656e74792829207b0a0973656c662e726f6c6c4d616e792832302c2031290a0973656c662e536f2873656c662e67616d652e53636f726528292c2073686f756c642e457175616c2c203230290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920546573745370617265526563656976657353696e676c65526f6c6c426f6e75732829207b0a0973656c662e726f6c6c537061726528290a0973656c662e67616d652e526f6c6c2834290a0973656c662e67616d652e526f6c6c2833290a0973656c662e726f6c6c4d616e792831362c2030290a0973656c662e536f2873656c662e67616d652e53636f726528292c2073686f756c642e457175616c2c203232290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e675465737473292054657374537472696b655265636965766573446f75626c65526f6c6c426f6e75732829207b0a0973656c662e726f6c6c537472696b6528290a0973656c662e67616d652e526f6c6c2834290a0973656c662e67616d652e526f6c6c2833290a0973656c662e726f6c6c4d616e792831362c2030290a0973656c662e536f2873656c662e67616d652e53636f726528292c2073686f756c642e457175616c2c203234290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920546573745065726665637447616d652829207b0a0973656c662e726f6c6c4d616e792831322c203130290a0973656c662e536f2873656c662e67616d652e53636f726528292c2073686f756c642e457175616c2c20333030290a7d0a0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920726f6c6c4d616e792874696d65732c2070696e7320696e7429207b0a09666f722078203a3d20303b2078203c2074696d65733b20782b2b207b0a090973656c662e67616d652e526f6c6c2870696e73290a097d0a7d0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920726f6c6c53706172652829207b0a0973656c662e67616d652e526f6c6c2835290a0973656c662e67616d652e526f6c6c2835290a7d0a66756e63202873656c66202a426f776c696e6747616d6553636f72696e6754657374732920726f6c6c537472696b652829207b0a0973656c662e67616d652e526f6c6c283130290a7d0a", "environment_controller_hvac_fake_test.go": "7061636b616765206578616d706c65730a0a696d706f72742022737472696e6773220a0a747970652046616b654856414320737472756374207b0a097374617465202020202020206d61705b737472696e675d626f6f6c0a0974656d706572617475726520696e740a7d0a0a66756e63204e657746616b6548617264776172652829202a46616b6548564143207b0a0972657475726e202646616b65485641437b0a090973746174653a206d61705b737472696e675d626f6f6c7b0a09090922686561746572223a2020202020202020202066616c73652c0a09090922626c6f776572223a2020202020202020202066616c73652c0a09090922636f6f6c6572223a2020202020202020202066616c73652c0a09090922686967682d74656d702d616c61726d223a2066616c73652c0a090909226c6f772d74656d702d616c61726d223a202066616c73652c0a09097d2c0a097d0a7d0a0a66756e63202873656c66202a46616b65485641432920416374697661746548656174657228292020202020202020202020202020207b2073656c662e73746174655b22686561746572225d203d2074727565207d0a66756e63202873656c66202a46616b654856414329204163746976617465426c6f77657228292020202020202020202020202020207b2073656c662e73746174655b22626c6f776572225d203d2074727565207d0a66756e63202873656c66202a46616b654856414329204163746976617465436f6f6c657228292020202020202020202020202020207b2073656c662e73746174655b22636f6f6c6572225d203d2074727565207d0a66756e63202873656c66202a46616b6548564143292041637469766174654869676854656d7065726174757265416c61726d2829207b2073656c662e73746174655b2268696768225d203d2074727565207d0a66756e63202873656c66202a46616b6548564143292041637469766174654c6f7754656d7065726174757265416c61726d282920207b2073656c662e73746174655b226c6f77225d203d2074727565207d0a0a66756e63202873656c66202a46616b654856414329204465616374697661746548656174657228292020202020202020202020202020207b2073656c662e73746174655b22686561746572225d203d2066616c7365207d0a66756e63202873656c66202a46616b6548564143292044656163746976617465426c6f77657228292020202020202020202020202020207b2073656c662e73746174655b22626c6f776572225d203d2066616c7365207d0a66756e63202873656c66202a46616b6548564143292044656163746976617465436f6f6c657228292020202020202020202020202020207b2073656c662e73746174655b22636f6f6c6572225d203d2066616c7365207d0a66756e63202873656c66202a46616b65485641432920446561637469766174654869676854656d7065726174757265416c61726d2829207b2073656c662e73746174655b2268696768225d203d2066616c7365207d0a66756e63202873656c66202a46616b65485641432920446561637469766174654c6f7754656d7065726174757265416c61726d282920207b2073656c662e73746174655b226c6f77225d203d2066616c7365207d0a0a66756e63202873656c66202a46616b65485641432920497348656174696e67282920626f6f6c2020202020202020202020207b2072657475726e2073656c662e73746174655b22686561746572225d207d0a66756e63202873656c66202a46616b654856414329204973426c6f77696e67282920626f6f6c2020202020202020202020207b2072657475726e2073656c662e73746174655b22626c6f776572225d207d0a66756e63202873656c66202a46616b654856414329204973436f6f6c696e67282920626f6f6c2020202020202020202020207b2072657475726e2073656c662e73746174655b22636f6f6c6572225d207d0a66756e63202873656c66202a46616b654856414329204869676854656d7065726174757265416c61726d282920626f6f6c207b2072657475726e2073656c662e73746174655b2268696768225d207d0a66756e63202873656c66202a46616b654856414329204c6f7754656d7065726174757265416c61726d282920626f6f6c20207b2072657475726e2073656c662e73746174655b226c6f77225d207d0a0a66756e63202873656c66202a46616b6548564143292053657443757272656e7454656d70657261747572652876616c756520696e7429207b2073656c662e74656d7065726174757265203d2076616c7565207d0a66756e63202873656c66202a46616b6548564143292043757272656e7454656d7065726174757265282920696e742020202020202020207b2072657475726e2073656c662e74656d7065726174757265207d0a0a2f2f20537472696e672072657475726e732074686520737461747573206f66206561636820686172647761726520636f6d706f6e656e7420656e636f64656420696e20612073696e676c652073706163652d64656c696d6974656420737472696e672e0a2f2f2055505045524341534520636f6d706f6e656e747320617265206163746976617465642e0a2f2f206c6f7765726361736520636f6d706f6e656e7473206172652064656163746976617465642e0a66756e63202873656c66202a46616b65485641432920537472696e67282920737472696e67207b0a0963757272656e74203a3d205b5d737472696e677b22686561746572222c2022626c6f776572222c2022636f6f6c6572222c20226c6f77222c202268696768227d0a09666f7220692c20636f6d706f6e656e74203a3d2072616e67652063757272656e74207b0a090969662073656c662e73746174655b636f6d706f6e656e745d207b0a09090963757272656e745b695d203d20737472696e67732e546f55707065722863757272656e745b695d290a09097d0a097d0a0972657475726e20737472696e67732e4a6f696e2863757272656e742c20222022290a7d0a", "environment_controller_test.go": "7061636b616765206578616d706c65730a0a696d706f727420280a09226769746875622e636f6d2f736d61727479737472656574732f617373657274696f6e732f73686f756c64220a09226769746875622e636f6d2f736d61727479737472656574732f67756e6974220a290a0a7479706520456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726520737472756374207b0a092a67756e69742e466978747572650a0968617264776172652020202a46616b65485641430a09636f6e74726f6c6c6572202a456e7669726f6e6d656e74436f6e74726f6c6c65720a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c657246697874757265292053657475702829207b0a0973656c662e6861726477617265203d204e657746616b65486172647761726528290a0973656c662e636f6e74726f6c6c6572203d204e6577436f6e74726f6c6c65722873656c662e6861726477617265290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529205465737445766572797468696e675475726e65644f66664174537461727475702829207b0a0973656c662e6163746976617465416c6c4861726477617265436f6d706f6e656e747328290a0973656c662e636f6e74726f6c6c6572203d204e6577436f6e74726f6c6c65722873656c662e6861726477617265290a0973656c662e617373657274416c6c4861726477617265436f6d706f6e656e7473417265446561637469766174656428290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529205465737445766572797468696e674f66665768656e436f6d666f727461626c652829207b0a0973656c662e7365747570436f6d666f727461626c65456e7669726f6e6d656e7428290a0973656c662e617373657274416c6c4861726477617265436f6d706f6e656e7473417265446561637469766174656428290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c657246697874757265292054657374436f6f6c6572416e64426c6f7765725768656e486f742829207b0a0973656c662e7365747570486f74456e7669726f6e6d656e7428290a0973656c662e536f2873656c662e68617264776172652e537472696e6728292c2073686f756c642e457175616c2c202268656174657220424c4f57455220434f4f4c4552206c6f77206869676822290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c657246697874757265292054657374486561746572416e64426c6f7765725768656e436f6c642829207b0a0973656c662e7365747570436f6c64456e7669726f6e6d656e7428290a0973656c662e536f2873656c662e68617264776172652e537472696e6728292c2073686f756c642e457175616c2c202248454154455220424c4f57455220636f6f6c6572206c6f77206869676822290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529205465737448696768416c61726d4f6e496641745468726573686f6c642829207b0a0973656c662e7365747570426c617a696e67456e7669726f6e6d656e7428290a0973656c662e536f2873656c662e68617264776172652e537472696e6728292c2073686f756c642e457175616c2c202268656174657220424c4f57455220434f4f4c4552206c6f77204849474822290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c6572466978747572652920546573744c6f77416c61726d4f6e496641745468726573686f6c642829207b0a0973656c662e7365747570467265657a696e67456e7669726f6e6d656e7428290a0973656c662e536f2873656c662e68617264776172652e537472696e6728292c2073686f756c642e457175616c2c202248454154455220424c4f57455220636f6f6c6572204c4f57206869676822290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529207365747570436f6d666f727461626c65456e7669726f6e6d656e742829207b0a0973656c662e68617264776172652e53657443757272656e7454656d706572617475726528434f4d464f525441424c45290a0973656c662e636f6e74726f6c6c65722e526567756c61746528290a7d0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529207365747570486f74456e7669726f6e6d656e742829207b0a0973656c662e68617264776172652e53657443757272656e7454656d706572617475726528544f4f5f484f54290a0973656c662e636f6e74726f6c6c65722e526567756c61746528290a7d0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529207365747570426c617a696e67456e7669726f6e6d656e742829207b0a0973656c662e68617264776172652e53657443757272656e7454656d7065726174757265285741595f544f4f5f484f54290a0973656c662e636f6e74726f6c6c65722e526567756c61746528290a7d0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529207365747570436f6c64456e7669726f6e6d656e742829207b0a0973656c662e68617264776172652e53657443757272656e7454656d706572617475726528544f4f5f434f4c44290a0973656c662e636f6e74726f6c6c65722e526567756c61746528290a7d0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529207365747570467265657a696e67456e7669726f6e6d656e742829207b0a0973656c662e68617264776172652e53657443757272656e7454656d7065726174757265285741595f544f4f5f434f4c44290a0973656c662e636f6e74726f6c6c65722e526567756c61746528290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c65724669787475726529206163746976617465416c6c4861726477617265436f6d706f6e656e74732829207b0a0973656c662e68617264776172652e4163746976617465426c6f77657228290a0973656c662e68617264776172652e416374697661746548656174657228290a0973656c662e68617264776172652e4163746976617465436f6f6c657228290a0973656c662e68617264776172652e41637469766174654869676854656d7065726174757265416c61726d28290a0973656c662e68617264776172652e41637469766174654c6f7754656d7065726174757265416c61726d28290a7d0a0a66756e63202873656c66202a456e7669726f6e6d656e74436f6e74726f6c6c6572466978747572652920617373657274416c6c4861726477617265436f6d706f6e656e747341726544656163746976617465642829207b0a0973656c662e536f2873656c662e68617264776172652e537472696e6728292c2073686f756c642e457175616c2c202268656174657220626c6f77657220636f6f6c6572206c6f77206869676822290a7d0a"}

func init() {
	gunit.Validate("2486151de57167557cd22637f721ec34")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
