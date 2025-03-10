package types

import (
	"fmt"
	"time"

	proto "github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bandprotocol/chain/v3/pkg/tss"
	tsstypes "github.com/bandprotocol/chain/v3/x/tss/types"
)

var (
	_ sdk.Msg = &MsgTransitionGroup{}
	_ sdk.Msg = &MsgForceTransitionGroup{}
	_ sdk.Msg = &MsgRequestSignature{}
	_ sdk.Msg = &MsgActivate{}
	_ sdk.Msg = &MsgUpdateParams{}

	_ sdk.HasValidateBasic = (*MsgTransitionGroup)(nil)
	_ sdk.HasValidateBasic = (*MsgForceTransitionGroup)(nil)
	_ sdk.HasValidateBasic = (*MsgRequestSignature)(nil)
	_ sdk.HasValidateBasic = (*MsgActivate)(nil)
	_ sdk.HasValidateBasic = (*MsgUpdateParams)(nil)

	_ types.UnpackInterfacesMessage = &MsgRequestSignature{}
)

// ====================================
// MsgTransitionGroup
// ====================================

// NewMsgTransitionGroup creates a new MsgTransitionGroup instance.
func NewMsgTransitionGroup(
	members []string,
	threshold uint64,
	execTime time.Time,
	authority string,
) *MsgTransitionGroup {
	return &MsgTransitionGroup{
		Members:   members,
		Threshold: threshold,
		Authority: authority,
		ExecTime:  execTime,
	}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgTransitionGroup) ValidateBasic() error {
	// Validate members address
	for _, member := range m.Members {
		if _, err := sdk.AccAddressFromBech32(member); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("invalid member address: %s", err)
		}
	}

	// Check duplicate member
	existed := make(map[string]bool)
	for _, member := range m.Members {
		if existed[member] {
			return ErrMemberDuplicated.Wrapf("member %s is duplicated", member)
		}
		existed[member] = true
	}

	// Validate sender address
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	// Validate threshold must be less than or equal to members but not zero
	if m.Threshold == 0 {
		return ErrInvalidThreshold.Wrap("threshold must be greater than zero")
	}
	if m.Threshold > uint64(len(m.Members)) {
		return ErrInvalidThreshold.Wrapf(
			"threshold (%d) must be less than or equal to the number of members (%d)",
			m.Threshold,
			len(m.Members),
		)
	}

	return nil
}

// ====================================
// MsgForceTransitionGroup
// ====================================

// NewMsgForceTransitionGroup creates a new MsgForceTransitionGroup instance.
func NewMsgForceTransitionGroup(
	incomingGroupID tss.GroupID,
	execTime time.Time,
	authority string,
) *MsgForceTransitionGroup {
	return &MsgForceTransitionGroup{
		IncomingGroupID: incomingGroupID,
		ExecTime:        execTime,
		Authority:       authority,
	}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgForceTransitionGroup) ValidateBasic() error {
	// Validate sender address
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	if m.IncomingGroupID == 0 {
		return ErrInvalidGroupID.Wrap("incoming group ID must not be zero")
	}

	return nil
}

// ====================================
// MsgRequestSignature
// ====================================

// NewMsgRequestSignature creates a new MsgRequestSignature.
func NewMsgRequestSignature(
	content tsstypes.Content,
	feeLimit sdk.Coins,
	sender string,
) (*MsgRequestSignature, error) {
	m := &MsgRequestSignature{
		FeeLimit: feeLimit,
		Sender:   sender,
	}
	err := m.SetContent(content)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// GetContent returns the content of MsgRequestSignature.
func (m *MsgRequestSignature) GetContent() tsstypes.Content {
	content, ok := m.Content.GetCachedValue().(tsstypes.Content)
	if !ok {
		return nil
	}
	return content
}

// ValidateBasic does a sanity check on the provided data
func (m MsgRequestSignature) ValidateBasic() error {
	// Validate sender address
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	// validate fee limit
	if !m.FeeLimit.IsValid() {
		return sdkerrors.ErrInvalidCoins.Wrap(m.FeeLimit.String())
	}

	// validate if fee limits are all positive
	if !m.FeeLimit.IsAllPositive() {
		return sdkerrors.ErrInvalidCoins.Wrap(m.FeeLimit.String())
	}

	content, ok := m.Content.GetCachedValue().(tsstypes.Content)
	if !ok {
		return sdkerrors.ErrInvalidType.Wrapf("expected type Content, got %T", content)
	}

	if err := content.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// SetContent sets the content for MsgRequestSignature.
func (m *MsgRequestSignature) SetContent(content tsstypes.Content) error {
	msg, ok := content.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.Content = any
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgRequestSignature) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content tsstypes.Content
	return unpacker.UnpackAny(m.Content, &content)
}

// ====================================
// MsgActivate
// ====================================

// NewMsgActivate creates a new MsgActivate instance.
func NewMsgActivate(sender string, groupID tss.GroupID) *MsgActivate {
	return &MsgActivate{
		Sender:  sender,
		GroupID: groupID,
	}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgActivate) ValidateBasic() error {
	// Validate member address
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	// Validate group ID
	if m.GroupID == 0 {
		return ErrInvalidGroupID.Wrap("group ID must not be zero")
	}

	return nil
}

// ====================================
// MsgUpdateParams
// ====================================

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(authority string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority,
		Params:    params,
	}
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}
