package main

import (
	"reflect"
	"testing"
)

func Test_Parseini(t *testing.T) {
	playbooks, err := parseini("./test/playbook.ini")
	if err != nil {
		t.Error("parseini error: ", err)
	}
	my_playbook_1 := playbooks[0]
	my_playbook_2 := playbooks[1]
	if my_playbook_1.playbook_type == "scp" && my_playbook_2.playbook_type == "ssh" {
		t.Log("parseini get type pass")
	} else {
		t.Error("parseini get type failed")
	}
	if my_playbook_1.src == "playbook.ini" && my_playbook_1.dst == "/home/playbook.ini" && my_playbook_2.command == "cat /home/playbook.ini" {
		t.Log("get other paraments pass")
	} else {
		t.Error("parseini get other paraments failed")
	}
}

func Test_StringToSshconfig(t *testing.T) {
	ssh_string_type_1, err := stringToSshconfig("192.168.10.1")
	if err != nil {
		t.Error("stringToSshconfig error: ", err)
	}
	if ssh_string_type_1.user == "root" {
		t.Log("192.168.10.1 format pass")
	} else {
		t.Error("192.168.10.1 format failed")
	}
	ssh_string_type_2, err2 := stringToSshconfig("weakiwi@192.168.10.1:2233")
	if err2 != nil {
		t.Error("stringToSshconfig error: ", err2)
	}
	if ssh_string_type_2.user == "weakiwi" {
		t.Log("weakiwi@192.168.10.1:2233 user part  pass")
	} else {
		t.Error("weakiwi@192.168.10.1:2233 user part pass")
	}
	if ssh_string_type_2.address == "192.168.10.1" {
		t.Log("weakiwi@192.168.10.1:2233 address part  pass")
	} else {
		t.Error("weakiwi@192.168.10.1:2233 address part pass")
	}
	if ssh_string_type_2.port == "2233" {
		t.Log("weakiwi@192.168.10.1:2233 port part  pass")
	} else {
		t.Error("weakiwi@192.168.10.1:2233 port part pass")
	}
}

func Test_ParseHostfile(t *testing.T) {
	sshconfigs, err := parseHostfile("./test/host1")
	if err != nil {
		t.Error("parseHostfile error: ", err)
	}
	var my_sshconfig sshconfig
	my_sshconfig.user = "root"
	my_sshconfig.port = "22"
	my_sshconfig.address = "118.190.75.175"
	if reflect.DeepEqual(sshconfigs[0], my_sshconfig) {
		t.Log("parseHostfile pass")
	} else {
		t.Error("parseHostfile failed", sshconfigs[0])
		t.Error("parseHostfile failed", my_sshconfig)
	}
}
