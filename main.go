package main

import (
	"log"

	pb "protocolbufferlearn/student"
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
