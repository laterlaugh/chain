package round1

import (
	"errors"

	"github.com/bandprotocol/chain/v2/pkg/event"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event represents the parsed information from a create_group event.
type Event struct {
	GroupID    tss.GroupID
	MemberID   tss.MemberID
	Threshold  uint64
	DKGContext []byte
	Members    []string
}

// ParseEvent parses the ABCIMessageLog and extracts the relevant information to create a create_group event.
// It returns the parsed Event or an error if parsing fails.
func ParseEvent(log sdk.ABCIMessageLog, address string) (*Event, error) {
	gid, err := event.GetEventValueUint64(log, types.EventTypeCreateGroup, types.AttributeKeyGroupID)
	if err != nil {
		return nil, err
	}

	threshold, err := event.GetEventValueUint64(log, types.EventTypeCreateGroup, types.AttributeKeyThreshold)
	if err != nil {
		return nil, err
	}

	dkgContext, err := event.GetEventValueBytes(log, types.EventTypeCreateGroup, types.AttributeKeyDKGContext)
	if err != nil {
		return nil, err
	}

	members := event.GetEventValues(log, types.EventTypeCreateGroup, types.AttributeKeyMember)

	mid, err := GetMemberID(members, address)
	if err != nil {
		return nil, err
	}

	return &Event{
		GroupID:    tss.GroupID(gid),
		MemberID:   mid,
		Threshold:  threshold,
		DKGContext: dkgContext,
		Members:    members,
	}, nil
}

// GetMemberID returns the member ID corresponding to the provided address.
// It searches through the members list and returns the member ID if found.
// If no member with the given address is found, it returns an error.
func GetMemberID(members []string, address string) (tss.MemberID, error) {
	for idx, member := range members {
		if member == address {
			return tss.MemberID(idx + 1), nil
		}
	}

	return 0, errors.New("failed to find member in the event")
}
