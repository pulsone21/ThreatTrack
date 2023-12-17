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
	s *api.ApiServer
}

func TestSuite(t *testing.T) {
	fmt.Println("Creating new Test Suite")
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	suite.Run(t, new(SuiteTest))
	fmt.Println("Creating new Test Suite")
}

func (t *SuiteTest) SetupSuite() {
	t.s = api.NewServer()

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
