package main

import (
	"PersonStruct"
	"errors"
	"log"
	mem "memcash"
	"net"
	pb "protobuf"
	mongo "subdmongo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//export
const (
	port = ":50051"
)

type server struct {
	pb.ArenaServer
}

func (s *server) SayBalance(c context.Context, p *pb.BalanceRequest) (*pb.BalanceReply, error) {

	//	pb.BalanceReply reply{}
	//LoginInformation
	b, r := mongo.GetBalance(p.Name)
	if !b {
		return new(pb.BalanceReply), errors.New("Error")
	}
	pr := pb.BalanceReply{Name: p.Name, Balance: float32(r)}
	return &pr, nil
}
func (s *server) LogIn(c context.Context, p *pb.LoginRequest) (*pb.LoginReply, error) {
	pers, b := PersonStruct.FindPersonByLogin(p.Name, p.Password)
	if !b {
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
	result := pb.ArenaResultReply{MessageOut: pb.ArenaResultReply_GOOD}
	return &result, nil
}

//var s mongo.SessionGame
func main() {
	//mongo.InitiateSession()
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
