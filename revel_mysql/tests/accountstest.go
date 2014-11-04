package tests

import "github.com/revel/revel"

type AccountsTest struct {
	revel.TestSuite
}

func (t *AccountsTest) Before() {
	println("Set up")
}

func (t *AccountsTest) TestThatIndexPageWorks() {
	t.Get("/accounts")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AccountsTest) TestReturnUserRecords() {
	t.Get("/accounts")
	t.AssertContainsRegex("3wZbIOQHLHCt58Kt")
}

func (t *AccountsTest) After() {
	println("Tear down")
}
