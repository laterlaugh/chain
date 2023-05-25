package types

const (
	EventTypeCreateGroup     = "create_group"
	EventTypeSubmitDKGRound1 = "submit_dkg_round1"
	EventTypeRound1Success   = "round1_success"
	EventTypeSubmitDKGRound2 = "submit_dkg_round2"
	EventTypeRound2Success   = "round2_success"
	EventTypeComplainSuccess = "complain_success"
	EventTypeComplainFailed  = "complain_failed"
	EventTypeConfirmSuccess  = "confirm_success"
	EventTypeConfirmFailed   = "confirm_failed"
	EventTypeRound3Success   = "round3_success"
	EventTypeRound3Failed    = "round3_failed"

	AttributeKeyGroupID    = "group_id"
	AttributeKeyMemberID   = "member_id"
	AttributeKeyMember     = "member"
	AttributeKeySize       = "size"
	AttributeKeyThreshold  = "threshold"
	AttributeKeyPubKey     = "pub_key"
	AttributeKeyStatus     = "status"
	AttributeKeyDKGContext = "dkg_context"
	AttributeKeyRound1Data = "round1_data"
	AttributeKeyRound2Data = "round2_data"
	AttributeKeyMemberIDI  = "member_id_i"
	AttributeKeyMemberIDJ  = "member_id_j"
	AttributeKeyKeySym     = "key_sym"
	AttributeKeyNonceSym   = "nonce_sym"
	AttributeKeySignature  = "signature"
	AttributeKeyComplain   = "complain"
	AttributeOwnPubKeySig  = "own_pub_key_sig"
)
