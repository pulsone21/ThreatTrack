package tests

import (
	"ContentManagement/api"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	server       *api.ApiServer
	serverAdress string
}

func TestSuite(t *testing.T) {
	os.Setenv("BACKEND_PORT", "5666")
	fmt.Println("Creating new Test Suite")
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	suite.Run(t, new(SuiteTest))
	fmt.Println("Creating new Test Suite")
}

func (t *SuiteTest) SetupSuite() {
	t.server = api.NewServer()
	t.serverAdress = fmt.Sprintf("http://localhost:%s", os.Getenv("BACKEND_PORT"))
	go func() { t.server.Run() }()
}

// Run After All Test Done
func (t *SuiteTest) TearDownSuite() {
	// stop data base docker start ContentDB
	cmd := exec.Command("docker", "stop", "ContentDB")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error in TearDownSuite")
		panic(err.Error())
	}
	fmt.Println("TearDownSuite Output: ", string(out))
}

// Run Before a Test
func (t *SuiteTest) SetupTest() {

}

// Run After a Test
func (t *SuiteTest) TearDownTest() {
}
