package handlers

import (
	"fmt"
	"github.com/rancherio/go-rancher/client"
	"testing"
)

// Test the filterDockerMessage to make sure are filtering the right messages
func TestFilterDockerMessages(t *testing.T) {
	errChan := make(chan string, 2)
	defer close(errChan)
	machine := &client.Machine{
		ExternalId: "uuid-1",
		Name:       "machine-1",
	}
	prefix := "time=\"2015-04-01T13:47:25-07:00\" level=\"info\" "

	testString := prefix + "msg=\"Message\" "
	checkField("Test1", "Message", filterDockerMessage(testString, machine, errChan), t)

	testString = prefix + "msg=\"Message with externalId=uuid-1\" "
	checkField("Test2", "", filterDockerMessage(testString, machine, errChan), t)

	testString = prefix + "msg=\"Message with name=machine-1\" "
	checkField("Test3", "", filterDockerMessage(testString, machine, errChan), t)

	testString = prefix + "msg=\"Message with random characters: =\"=\"\" "
	checkField("Test4", "Message with random characters: =\"=\"", filterDockerMessage(testString, machine, errChan), t)

	testString = prefix + "msg="
	checkField("Test5", "", filterDockerMessage(testString, machine, errChan), t)

	testString = prefix + "Really weird message"
	checkField("Test6", "", filterDockerMessage(testString, machine, errChan), t)

	// warning level check
	testString = "time=\"2015-04-01T13:47:25-07:00\" level=\"warning\" msg=\"error message\" "
	checkField("Test7", "", filterDockerMessage(testString, machine, errChan), t)

	// fatal level check
	testString = "time=\"2015-04-01T13:47:25-07:00\" level=\"fatal\" msg=\"error message\" "
	checkField("Test8", "", filterDockerMessage(testString, machine, errChan), t)

	//error level check
	testString = "time=\"2015-04-01T13:47:25-07:00\" level=\"error\" msg=\"Error creating machine: error message\" "
	checkField("Test9", "", filterDockerMessage(testString, machine, errChan), t)
	checkField("Test9:errString", "error message", <-errChan, t)
}

// Tests the simplest case of successfully receiving, routing, and handling
// three events.
func TestParseConnectionArgs(t *testing.T) {
	ca := "/foo/bar/ca.pem"
	cert := "/foo/bar\\ baz/cert.pem"
	key := "/foo/bar/key.pem"
	endpoint := "tcp://127.0.0.1:2376"

	// "Normal"
	testArgs := fmt.Sprintf("--tls --tlscacert=%v   --tlscert=%v --tlskey=%v -H=\"%v\" ",
		ca, cert, key, endpoint)
	testParse(testArgs, ca, cert, key, endpoint, t)

	// -H at beginning
	testArgs = fmt.Sprintf("-H=\"%v\" --tlscacert=%v --tlscert=%v --tlskey=%v --tls",
		endpoint, ca, cert, key)
	testParse(testArgs, ca, cert, key, endpoint, t)

	// -H at end with no quotes
	testArgs = fmt.Sprintf("--tls --tlscacert=%v   --tlscert=%v --tlskey=%v -H=%v",
		ca, cert, key, endpoint)
	testParse(testArgs, ca, cert, key, endpoint, t)

	// -H at beginning with no quotes
	testArgs = fmt.Sprintf("-H=%v --tlscacert=%v --tlscert=%v --tlskey=%v --tls",
		endpoint, ca, cert, key)
	testParse(testArgs, ca, cert, key, endpoint, t)

	testArgs = fmt.Sprintf("--tlscacert=%v --tlscert=%v --tlskey=%v --tls", ca, cert, key)
	_, err := parseConnectionArgs(testArgs)
	if err == nil {
		t.Fatalf("Parse should have failed because of missing host.")
	}

}

// Test the generic build create command function
func TestBuildMachineCreateCmd(t *testing.T) {

	// Tests: false boolean param, true boolean param, missing boolean param, and string params
	testCmd := []string{
		"create",
		"-d",
		"digitalocean",
		"--digitalocean-access-token",
		"abc",
		"--digitalocean-image",
		"ubuntu-14-04-x64",
		"--digitalocean-ipv6",
		"--digitalocean-region",
		"sfo1",
		"--digitalocean-size",
		"1gb",
		"testDO"}
	machine := &client.Machine{
		DigitaloceanConfig: &client.DigitaloceanConfig{
			AccessToken: "abc",
			Region:      "sfo1",
			Size:        "1gb",
			Image:       "ubuntu-14-04-x64",
			Ipv6:        true,
			Backups:     false,
		},
		Kind:   "machine",
		Driver: "DigitalOcean",
		Name:   "testDO",
	}
	checkCommands(testCmd, machine, t)

	// Test for no params
	testCmd = []string{
		"create",
		"-d",
		"virtualbox",
		"testVB"}
	machine = &client.Machine{
		VirtualboxConfig: &client.VirtualboxConfig{},
		Kind:             "machine",
		Driver:           "VirtualBox",
		Name:             "testVB",
	}
	checkCommands(testCmd, machine, t)
}

func strsEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkCommands(testCmd []string, machine *client.Machine, t *testing.T) {
	cmd, err := buildMachineCreateCmd(machine)
	if err != nil {
		t.Fatalf("Error building command", err)
	}

	if !strsEquals(cmd, testCmd) {
		t.Logf("Mismatch commands.  Expected: [%v], Actual: [%v]", testCmd, cmd)
		t.FailNow()
	}
}

func testParse(testArgs, ca, cert, key, endpoint string, t *testing.T) {
	config, err := parseConnectionArgs(testArgs)
	if err != nil {
		t.Fatalf("Error parsing.", err)
	}
	checkField("endpoint", endpoint, config.endpoint, t)
	checkField("ca", ca, config.caCert, t)
	checkField("cert", cert, config.cert, t)
	checkField("key", key, config.key, t)
}

func checkField(field, expected, actual string, t *testing.T) {
	if expected != actual {
		t.Logf("Mismatch on field [%v]. Expected: [%v],  Actual: [%v]", field, expected, actual)
		t.FailNow()
	}
}
