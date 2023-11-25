package sp

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
