package sp

import (
	"fmt"
)

var (
	SplitflapStateType  SplitFlapType = "SplitflapStateType"
	LogType             SplitFlapType = "LogType"
	AckType             SplitFlapType = "AckType"
	SupervisorStateType SplitFlapType = "SupervisorStateType"
	Unknown             SplitFlapType = "Unknown"
)

type SplitFlapType string

func (m *FromSplitflap) GetPayloadType() SplitFlapType {
	switch m.GetPayload().(type) {
	case *FromSplitflap_SplitflapState:
		return SplitflapStateType
	case *FromSplitflap_Log:
		return LogType
	case *FromSplitflap_Ack:
		return AckType
	case *FromSplitflap_SupervisorState:
		return SupervisorStateType
	}

	return Unknown
}

func (m *FromSplitflap) PrintSplitflapState() {
	msg := ""
	payloadType := m.GetPayloadType()
	if payloadType == SplitflapStateType {
		state := m.GetPayload().(*FromSplitflap_SplitflapState)
		for i, module := range state.SplitflapState.GetModules() {
			if module.GetState().String() != "NORMAL" {
				msg += fmt.Sprintf("Index %d: %s \n", i+1, module.GetState())
			}
		}
	}
	if msg != "" {
		fmt.Println(msg)
	}
}
