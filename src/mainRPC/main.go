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
	mgo "subdmongo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//export
//protoc --go_out=plugins=grpc,import_path=protobuf:. *.proto
//protoc --csharp_out ArenaService --grpc_out ArenaService ArenaService.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin
//%UserProfile%\.nuget\packages\Grpc.Tools\1.16.0\tools\windows_x64\protoc.exe --csharp_out=ArenaService --grpc_out=ArenaService ArenaService.proto --plugin=protoc-gen-grpc=%UserProfile%\.nuget\packages\Grpc.Tools\1.16.0\tools\windows_x64\grpc_csharp_plugin.exe
//%UserProfile%\.nuget\packages\Grpc.Tools\1.16.0\tools\windows_x64\protoc.exe --csharp_out=ArenaService --grpc_out=ArenaService ArenaService.proto --plugin=protoc-gen-grpc=%UserProfile%\.nuget\packages\Grpc.Tools\1.16.0\tools\windows_x64\grpc_csharp_plugin.exe
type server struct {
	pb.ArenaServer
}

func (s *server) News(c context.Context, p *pb.BalanceRequest) (*pb.NewsReply, error) {
	per, b := PersonStruct.FindPersonByToken(p.TockenID)
	if !b {
		return new(pb.NewsReply), errors.New("INVALID TOCKEN")
	}
	str := mgo.GetNews()

	pr := pb.NewsReply{NewsAlias: str[:]}
	return &pr, nil
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
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5050"
	} else {
		port = ":" + port
	}
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
		}
		if "exit" == guessColor {
			os.Exit(0)
			return
		}
		if "registr" == guessColor {
			if _, err := fmt.Scanf("%s", &guessColor); err != nil {
				fmt.Printf("%s\n", err)
			}
			n := guessColor
			if _, err := fmt.Scanf("%s", &guessColor); err != nil {
				fmt.Printf("%s\n", err)
			}
			p := guessColor
			_, err := PersonStruct.InsertPerson(n, p)
			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Printf("added")
			}
		}
		if "drop" == guessColor {
			mgo.DropBase()
			return
		}
	}
}
