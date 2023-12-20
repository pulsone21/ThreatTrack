package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	ApiServer string
}

func TestSuite(t *testing.T) {
	os.Setenv("BACKEND_PORT", "5666")
	fmt.Println("Creating new Test Suite")
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	suite.Run(t, new(SuiteTest))
}

func (t *SuiteTest) SetupSuite() {
	t.ApiServer = fmt.Sprintf("http://%s:%s", os.Getenv("BACKEND_ADRESS"), os.Getenv("BACKEND_PORT"))
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
	cmd = exec.Command("docker", "stop", "TestBackend")
	out, err = cmd.Output()
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
