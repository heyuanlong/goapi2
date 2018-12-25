package service

type SettleStruct struct {
	tm int64
}

func NewSettleStruct() *SettleStruct {
	return &SettleStruct{}
}

func (ts *SettleStruct) Init() {

}
func (ts *SettleStruct) Run() {
	Glock.Lock()
	defer func() {
		Glock.Unlock()
	}()

}
