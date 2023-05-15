package server

import (
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/net/packet"
	"github.com/imhinotori/void/internal/config"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Server struct {
	Listener      *net.Listener
	Configuration *config.Config
}

func CreateServer(cfgOptions ...config.Option) (*Server, error) {
	log.Log().Msgf("Creating Server...")
	cfg := config.Build(cfgOptions...)

	mc, err := net.ListenMC(":" + strconv.Itoa(cfg.Server.Port))
	if err != nil {
		return nil, err
	}

	return &Server{
		Listener:      mc,
		Configuration: cfg,
	}, nil
}

func (s *Server) AcceptPackets() error {
	log.Log().Msgf("Starting to accept Packets...")
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	protocol, intention, err := s.handshakePacket(conn)
	if err != nil {
		log.Error().Msgf("failed to handshake: %v", err)
	}

	handleIntention(conn, protocol, intention)
}

func handleIntention(conn net.Conn, protocol, intention int32) {
	switch intention {
	default: // Invalid
		log.Log().Msgf("unknown handshake intention: %v", intention)
	case 1:
		log.Log().Msgf("ping request")
		handlePingPacket(conn)
	case 2:
		log.Log().Msgf("login request")
	}
}

func (s *Server) handshakePacket(conn net.Conn) (protocol, intention int32, err error) {
	var (
		p                   packet.Packet
		Protocol, Intention packet.VarInt
		ServerAddress       packet.String
		ServerPort          packet.UnsignedShort
	)

	if err = conn.ReadPacket(&p); err != nil {
		return
	}

	err = p.Scan(&Protocol, &ServerAddress, &ServerPort, &Intention)
	return int32(Protocol), int32(Intention), err
}
