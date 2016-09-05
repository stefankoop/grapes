package main

import (
	"testing"
)

var demoInputOk = input{
	asyncFlag:   asyncFlag(false),
	configPath:  configPath("/demo"),
	keyPath:     keyPath("/demo"),
	serverGroup: serverGroup("demoServerGroup"),
	commandName: commandName("demo"),
	verifyFlag:  verifyFlag(false),
}

var demoInputNotOk = input{
	asyncFlag:   asyncFlag(false),
	configPath:  configPath(""),
	keyPath:     keyPath(""),
	serverGroup: serverGroup(""),
	verifyFlag:  verifyFlag(true),
}

func TestParse(t *testing.T) {
	in := input{}
	in.parse()
}

func TestValidateInput(t *testing.T) {
	dOk := demoInputOk
	dNotOk := demoInputNotOk
	if dOk.validate() != nil {
		t.FailNow()
	}
	if dNotOk.validate() == nil {
		t.FailNow()
	}
	dNotOk.configPath = demoInputOk.configPath
	if dNotOk.validate() == nil {
		t.FailNow()
	}
	dNotOk.keyPath = demoInputOk.keyPath
	if dNotOk.validate() == nil {
		t.FailNow()
	}
	dNotOk.serverGroup = demoInputOk.serverGroup
	if dNotOk.validate() == nil {
		t.FailNow()
	}
}

func TestValidateConfigPath(t *testing.T) {
	if err := demoInputOk.configPath.validate(&demoInputOk); err != nil {
		t.Fatalf("1")
	}
	if err := demoInputNotOk.configPath.validate(&demoInputNotOk); err == nil {
		t.Fatalf("2")
	}
}

func TestValidateKeyPath(t *testing.T) {
	if err := demoInputOk.keyPath.validate(&demoInputOk); err != nil {
		t.FailNow()
	}
	if err := demoInputNotOk.keyPath.validate(&demoInputNotOk); err == nil {
		t.FailNow()
	}
}

func TestValidateServerGroup(t *testing.T) {
	if err := demoInputOk.serverGroup.validate(&demoInputOk); err != nil {
		t.FailNow()
	}
	if err := demoInputNotOk.serverGroup.validate(&demoInputNotOk); err == nil {
		t.FailNow()
	}
}

func TestValidateCommandName(t *testing.T) {
	if err := demoInputOk.commandName.validate(&demoInputOk); err != nil {
		t.FailNow()
	}
	if err := demoInputNotOk.commandName.validate(&demoInputNotOk); err == nil {
		t.FailNow()
	}
}

func TestVerifyAction(t *testing.T) {

	s := servers{
		server{
			Name: "testN",
			Host: "testH",
			User: "testU",
		},
	}

	//should panic
	func() {
		defer func() {
			if err := recover(); err != nil {
				recover()
			}
		}()
		demoInputOk.verifyAction(s)
		t.FailNow()
	}()

	//should not panic
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.FailNow()
			}
		}()
		demoInputOk.verifyFlag = true
		demoInputOk.verifyAction(s)
	}()

}
