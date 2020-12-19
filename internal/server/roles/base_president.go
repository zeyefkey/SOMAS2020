package roles

//base President Object
type basePresident struct {
	id                 int
	budget             int
	speakerSalary      int
	resourceRequests   map[int]int
	resourceAllocation map[int]int
	ruleToVote         int
	taxAmount          int
}
func (p* basePresident) withdrawSpeakerSalary() error{
    return nil
}
func (p* basePresident) signalAllocationRequests(int) error{
    return nil
}
func (p* basePresident) replyAllocationRequests(int) error{
    return nil
}
func (p* basePresident) sendRuleToSpeaker(int) error{
    return nil
}
