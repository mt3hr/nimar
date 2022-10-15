package nimar

import (
	context "context"
	"errors"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/peer"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func NewNimeRServer() (NimaRServer, error) {
	serverImpl := &ServerImpl{}
	go func() {
		for _ = 0; true; time.Sleep(time.Second * 10) {
			for roomID, table := range serverImpl.tables {
				table.GetPlayer1().GetName()
				delete(serverImpl.tables, roomID)
			}
		}
	}()
	return serverImpl, nil
}

type ServerImpl struct {
	tables  map[string]*MTable
	players map[string]string
}

func (s *ServerImpl) MessageStream(req *JoinRoomRequest, ss NimaR_MessageStreamServer) error {
	s.tables[s.players[req.PlayerID]].GetPlayerByID(req.PlayerID).SetNimaRMessageStreamServer(&ss)
	return nil
}

func (s *ServerImpl) ListRooms(_ context.Context, _ *emptypb.Empty) (*Rooms, error) {
	rooms := &Rooms{}
	for _, room := range s.tables {
		r := &Room{
			RoomID:   room.GetID(),
			RoomName: room.GetName(),
		}
		rooms.Rooms = append(rooms.Rooms, r)
	}
	return rooms, nil
}

func (s *ServerImpl) tableToRoom(roomID string) *Room {
	playerNames := []string{}
	if s.tables[roomID].GetPlayer1() != nil {
		playerNames = append(playerNames, s.tables[roomID].GetPlayer1().GetName())
	}
	if s.tables[roomID].GetPlayer2() != nil {
		playerNames = append(playerNames, s.tables[roomID].GetPlayer2().GetName())
	}
	room := &Room{
		RoomID:      s.tables[roomID].GetID(),
		RoomName:    s.tables[roomID].GetName(),
		PlayerNames: playerNames,
	}
	return room
}

func (s *ServerImpl) CreateRoom(_ context.Context, r *CreateRoomRequest) (*Room, error) {
	roomID := uuid.New().String()
	s.tables[roomID] = NewTable(roomID, r.RoomName)
	return s.tableToRoom(roomID), nil
}

func (s *ServerImpl) GameTableStream(joinRequest *JoinRoomRequest, gameTableStreamServer NimaR_GameTableStreamServer) error {
	player := NewPlayer(joinRequest.PlayerName, joinRequest.PlayerID, &gameTableStreamServer)

	if s.tables[joinRequest.RoomID].GetPlayer1() == nil {
		s.tables[joinRequest.RoomID].SetPlayer1(player)
	} else if s.tables[joinRequest.RoomID].GetPlayer2() == nil {
		s.tables[joinRequest.RoomID].SetPlayer2(player)
	} else {
		return errors.New("すでに満室です")
	}

	return nil
}

func (s *ServerImpl) GetPlayerID(ctx context.Context, _ *emptypb.Empty) (*PlayerID, error) {
	if pr, ok := peer.FromContext(ctx); ok {
		addr := pr.Addr.String()
		if id, exist := s.players[addr]; exist {
			return &PlayerID{
				Playerid: id,
			}, nil
		} else {
			id = uuid.New().String()
			s.players[addr] = id
			return &PlayerID{
				Playerid: id,
			}, nil
		}
	}
	err := errors.New("なんか変です") // TODO
	return nil, err
}

func (s *ServerImpl) OperatorsStream(joinRequest *JoinRoomRequest, operatorsStreamServer NimaR_OperatorsStreamServer) error {
	if s.tables[joinRequest.RoomID].GetPlayer1().GetID() == joinRequest.PlayerID {
		s.tables[joinRequest.RoomID].GetPlayer1().SetNimaROperatorsStreamServer(&operatorsStreamServer)
	} else if s.tables[joinRequest.RoomID].GetPlayer2().GetID() == joinRequest.PlayerID {
		s.tables[joinRequest.RoomID].GetPlayer2().SetNimaROperatorsStreamServer(&operatorsStreamServer)
	} else {
		err := errors.New("なんか変です") //TODO
		return err
	}
	return nil
}

func (s *ServerImpl) Operate(_ context.Context, operator *Operator) (*emptypb.Empty, error) {
	err := s.tables[operator.RoomID].GetGameManager().ExecuteOperator(operator)
	return &emptypb.Empty{}, err
}

func (s *ServerImpl) mustEmbedUnimplementedNimaRServer() {
}
