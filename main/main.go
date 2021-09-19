package main

import (
	"dci/dci/context"
	"dci/dci/object"
	"dci/ddd/aggregate"
	"dci/ddd/entity"
)

func DDD() {
	paul := entity.NewPeople("Paul")
	mit := aggregate.NewSchool("mit")
	google := aggregate.NewCompany("Google")
	home := aggregate.NewHome()
	summerPalace := aggregate.NewPark("Summer Palace")

	// 上学
	mit.Receive(paul)
	mit.Run()
	// 回家
	home.ComeBack(paul)
	home.Run()
	// 工作
	google.Employ(paul)
	google.Run()
	// 公园游玩
	summerPalace.Welcome(paul)
	summerPalace.Run()
}

func DCI() {
	paul := object.NewPeople("Paul")
	mit := context.NewSchool("mit")
	google := context.NewCompany("Google")
	home := context.NewHome()
	summerPalace := context.NewPark("Summer Palace")

	// 上学
	mit.Receive(paul.CastStudent())
	mit.Start()
	// 回家
	home.ComeBack(paul.CastHuman())
	home.Start()
	// 工作
	google.Employ(paul.CastWorker())
	google.Start()
	// 公园游玩
	summerPalace.Welcome(paul.CastEnjoyer())
	summerPalace.Start()

}

func main() {
	DCI()
}
