package service

import "sync"

var Glock sync.Mutex
var BusiSettle = NewSettleStruct()
