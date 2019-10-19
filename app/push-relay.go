package app

import (
	"../util"
)
func Run() {
	api := NewApi(8080)
	api.Start()

	//po := push.NewPushOver("a5vtbmmvyuxui6pp9of7ayn8abnovg", "F0Jh5uUexIXDc0OF9KOyVYww8TE9GQ")
	//resp,err := po.Send(push.Message{
	//	Title:       "his",
	//	Message:     "another test",
	//})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(resp)
	<-util.Signal()
	api.Stop()
}