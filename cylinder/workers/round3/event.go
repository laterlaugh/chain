package round3

import (
	"github.com/bandprotocol/chain/v2/pkg/event"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event represents the data structure for round3 events.
type Event struct {
	GroupID tss.GroupID
}

// ParseEvent parses the round3 event from the given message log.
// It extracts the group ID from the log and returns the parsed Event or an error if parsing fails.
func ParseEvent(log sdk.ABCIMessageLog) (*Event, error) {
	gid, err := event.GetEventValueUint64(log, types.EventTypeRound2Success, types.AttributeKeyGroupID)
	if err != nil {
		return nil, err
	}

	return &Event{
		GroupID: tss.GroupID(gid),
	}, nil
}
