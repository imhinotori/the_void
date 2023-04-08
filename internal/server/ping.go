package server

import (
	"encoding/json"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/net"
	"github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func handlePingPacket(conn net.Conn) {
	var p packet.Packet

	for i := 0; i < 2; i++ {
		err := conn.ReadPacket(&p)
		if err != nil {
			log.Debug().Msgf("failed to read packet: %v", err)
			return
		}

		switch p.ID {
		case 0x00: // Player List Packet
			log.Debug().Msgf("Sending Player List to conn: %v", conn)
			err = conn.WritePacket(packet.Marshal(0x00, packet.String(listResp())))
		case 0x01: // Server Ping Packet
			log.Debug().Msgf("Sending Server Ping to conn: %v", conn)
			err = conn.WritePacket(p)
		}

		if err != nil {
			log.Debug().Msgf("failed to write packet: %v", err)
			return
		}
	}
}

type player struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func listResp() string {
	var list struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int      `json:"max"`
			Online int      `json:"online"`
			Sample []player `json:"sample"`
		} `json:"players"`
		Description chat.Message `json:"description"`
		FavIcon     string       `json:"favicon,omitempty"`
	}

	list.Version.Name = "Chat Server"
	list.Version.Protocol = 756
	list.Players.Max = 200
	list.Players.Online = 123
	list.Players.Sample = []player{} // must init. can't be nil
	list.Description = chat.Message{Text: "TO-DO", Color: "blue"}

	data, err := json.Marshal(list)
	if err != nil {
		log.Panic().Msg("Marshal JSON for status checking fail")
	}
	return string(data)
}
