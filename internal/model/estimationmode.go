package model

type EstimationMode int

const (
	StoryPointMode EstimationMode = iota
	TimeMode
)

func MakeEstimationMode(modeName string) (EstimationMode, error) {
	if modeName == "StoryPointMode" {
		return StoryPointMode, nil
	} else if modeName == "TimeMode" {
		return TimeMode, nil
	}
	return -1, &ModeError{"unsupported mode"}
}

type ModeError struct {
	s string
}

func (e *ModeError) Error() string {
	return e.s
}
