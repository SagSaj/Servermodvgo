package main

import (
	"PersonStruct"
	"errors"
	"fmt"
	"log"
	mem "memcash"
	"net"
	"os"
	pb "protobuf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//export
//protoc --go_out=plugins=grpc,import_path=protobuf:. *.proto
//protoc -I../../protos --csharp_out Greeter --grpc_out Greeter ../../protos/helloworld.proto --plugin=protoc-gen-grpc=${HOME}/.nuget/packages/grpc.tools.1.16.0/tools/linux_x64/grpc_csharp_plugin
const (
	port = ":50051"
)

type server struct {
	pb.ArenaServer
}

func (s *server) SayBalance(c context.Context, p *pb.BalanceRequest) (*pb.BalanceReply, error) {

	//	pb.BalanceReply reply{}
	//LoginInformation
	per, b := PersonStruct.FindPersonByToken(p.TockenID)
	if !b {
		return new(pb.BalanceReply), errors.New("INVALID TOCKEN")
	}
	pr := pb.BalanceReply{Name: per.Login, Balance: per.Balance}
	return &pr, nil
}
func (s *server) LogIn(c context.Context, p *pb.LoginRequest) (*pb.LoginReply, error) {
	pers, b := PersonStruct.FindPersonByLogin(p.Name, p.Password)
	if b != nil {
		return new(pb.LoginReply), errors.New("Error")
	}
	pr := pb.LoginReply{Name: p.Name, TockenID: pers.Tocken}
	return &pr, nil
}
func (s *server) ArenaInitiate(c context.Context, p *pb.ArenaInitiateRequest) (*pb.ArenaFindReply, error) {
	a := mem.Arena.AddArena(p.TockenArena)
	a.AddNewTocken(p.TockenID, p.TeamTocken)
	str := a.GetEnemies(p.TeamTocken)
	result := pb.ArenaFindReply{NamesAlias: str[:]}
	return &result, nil
}
func (s *server) ArenaFind(c context.Context, p *pb.ArenaFindRequest) (*pb.ArenaFindReply, error) {
	a := mem.Arena.FindArena(p.TockenArena)
	str := a.GetEnemies(p.TeamTocken)
	result := pb.ArenaFindReply{NamesAlias: str[:]}
	return &result, nil
}
func (s *server) ArenaResult(c context.Context, p *pb.ArenaResultRequest) (*pb.ArenaResultReply, error) {
	a := mem.Arena.FindArena(p.TockenArena)
	a.TeamWin(p.TeamTocken)
	result := pb.ArenaResultReply{MessageOut: pb.ArenaMessageOut_GOOD}
	return &result, nil
}
func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArenaServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//var s mongo.SessionGame
func main() {
	//mongo.InitiateSession()
	go RunServer()
	var guessColor string
	for {
		if _, err := fmt.Scanf("%s", &guessColor); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		if "exit" == guessColor {
			os.Exit(0)
			return
		}
		if "drop" == guessColor {

			return
		}
	}
}
