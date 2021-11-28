package ssh-config

import “testing”

func TestGetCommandArgs(t *testing.T) {
	arguments, err := getCommandArgs()
	var expected = "$HOME/.ssh/config"
	if err != nil {
		t.Log("error should be nil",err)
		t.Fail()
	}
	if  arguments["configFile"] !=  expected {
		t.Log("error should be ", expected,", but got ",arguments["configFile"])
        t.Fail()
    }
}
