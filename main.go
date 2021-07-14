package main

import (
	"log"

	pb "github.com/or0986113303/protocolbufferlearn/student"
)

func main() {
	s := &pb.Student{
		Name:   "Peng Jie",
		Age:    "24",
		Gender: "Male",
		Number: 99,
	}
	log.Println(
		s.GetName(),
		s.GetAge(),
		s.GetGender(),
		s.GetNumber(),
	)
}
