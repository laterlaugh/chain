package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/pkg/tss/testutil"
	"github.com/bandprotocol/chain/v2/x/bandtss/types"
	tsstypes "github.com/bandprotocol/chain/v2/x/tss/types"
)

func (s *KeeperTestSuite) TestSetInActive() {
	ctx, k, tssKeeper := s.ctx, s.app.BandtssKeeper, s.app.TSSKeeper
	s.SetupGroup(tsstypes.GROUP_STATUS_ACTIVE)
	address := sdk.AccAddress(testutil.TestCases[0].Group.Members[0].PubKey())

	k.SetInactiveStatuses(ctx, []sdk.AccAddress{address})

	status := k.GetStatus(ctx, address)

	member, err := tssKeeper.GetMemberByAddress(ctx, testutil.TestCases[0].Group.ID, address.String())
	s.Require().NoError(err)
	s.Require().False(member.IsActive)
	s.Require().Equal(types.MEMBER_STATUS_INACTIVE, status.Status)

}

func (s *KeeperTestSuite) TestSetActive() {
	ctx, k, tssKeeper := s.ctx, s.app.BandtssKeeper, s.app.TSSKeeper
	s.SetupGroup(tsstypes.GROUP_STATUS_ACTIVE)
	address := sdk.AccAddress(testutil.TestCases[0].Group.Members[0].PubKey())

	// Success case
	err := k.SetActiveStatuses(ctx, []sdk.AccAddress{address})
	s.Require().NoError(err)

	status := k.GetStatus(ctx, address)
	member, err := tssKeeper.GetMemberByAddress(ctx, testutil.TestCases[0].Group.ID, address.String())
	s.Require().NoError(err)
	s.Require().Equal(types.MEMBER_STATUS_ACTIVE, status.Status)
	s.Require().True(isActive)

	// Failed case - penalty
	k.SetInactiveStatuses(ctx, []sdk.AccAddress{address})

	err = k.SetActiveStatuses(ctx, []sdk.AccAddress{address})
	s.Require().ErrorIs(err, types.ErrTooSoonToActivate)

	// Failed case - no member
	err = k.SetActiveStatuses(ctx, []sdk.AccAddress{address})
	s.Require().Error(err)
}

func (s *KeeperTestSuite) TestSetLastActive() {
	ctx, k := s.ctx, s.app.BandtssKeeper
	s.SetupGroup(tsstypes.GROUP_STATUS_ACTIVE)
	address := sdk.AccAddress(testutil.TestCases[0].Group.Members[0].PubKey())

	// Success case
	err := k.SetLastActive(ctx, address)
	s.Require().NoError(err)

	status := k.GetStatus(ctx, address)
	s.Require().Equal(ctx.BlockTime(), status.LastActive)

	// Failed case
	k.SetInactiveStatuses(ctx, []sdk.AccAddress{address})

	err = k.SetLastActive(ctx, address)
	s.Require().Error(err)
}
