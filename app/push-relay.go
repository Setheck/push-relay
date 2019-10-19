package app

import "github.com/setheck/push-relay/util"

func Run() {
	api := NewApi(8080)
	api.Start()

	//po := push.NewPushOver("", "")
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
