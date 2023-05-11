package main

import "fmt"

/*
命令模式
*/

type Docker struct{}

func (d Docker) TreatEyes() {
	fmt.Println("治疗眼睛")
}

func (d Docker) TreatNose() {
	fmt.Println("治疗鼻子")
}

// Command 病历单
type Command interface {
	Treat()
}

// CommandTreatEye 治疗眼睛的病历单
type CommandTreatEye struct {
	docker Docker
}

func NewCommandTreatEye(docker Docker) *CommandTreatEye {
	return &CommandTreatEye{docker: docker}
}

func (c CommandTreatEye) Treat() {
	c.docker.TreatEyes()
}

// CommandTreatNose 治疗鼻子的病历单
type CommandTreatNose struct {
	docker Docker
}

func NewCommandTreatNose(docker Docker) *CommandTreatNose {
	return &CommandTreatNose{docker: docker}
}

func (c CommandTreatNose) Treat() {
	c.docker.TreatNose()
}

// 护士
type Nurse struct {
	command []Command
}

func (n Nurse) Noifty() {
	for _, v := range n.command {
		if v == nil {
			return
		} else {
			//执行病例
			v.Treat()
		}
	}
}

func main() {
	docker := Docker{}

	CommandTreatEye := NewCommandTreatEye(docker)
	CommandTreatNose := NewCommandTreatNose(docker)

	Nurse := Nurse{}
	Nurse.command = append(Nurse.command, CommandTreatEye)
	Nurse.command = append(Nurse.command, CommandTreatNose)
	Nurse.Noifty()
}
