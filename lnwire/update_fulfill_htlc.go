package lnwire

import "io"

// UpdateFulfillHTLC is sent by Alice to Bob when she wishes to settle a
// particular HTLC referenced by its HTLCKey within a specific active channel
// referenced by ChannelPoint.  A subsequent CommitSig message will be sent by
// Alice to "lock-in" the removal of the specified HTLC, possible containing a
// batch signature covering several settled HTLC's.
type UpdateFulfillHTLC struct {
	// ChanID references an active channel which holds the HTLC to be
	// settled.
	ChanID ChannelID

	// ID denotes the exact HTLC stage within the receiving node's
	// commitment transaction to be removed.
	ID uint64

	// PaymentPreimage is the R-value preimage required to fully settle an
	// HTLC.
	PaymentPreimage [32]byte

	// is this packet marked because of a large queue delay somewhere
	Marked uint32
}

// NewUpdateFulfillHTLC returns a new empty UpdateFulfillHTLC.
func NewUpdateFulfillHTLC(chanID ChannelID, id uint64,
	preimage [32]byte, marked uint32) *UpdateFulfillHTLC {

	return &UpdateFulfillHTLC{
		ChanID:          chanID,
		ID:              id,
		PaymentPreimage: preimage,
		Marked:          marked,
	}
}

// A compile time check to ensure UpdateFulfillHTLC implements the lnwire.Message
// interface.
var _ Message = (*UpdateFulfillHTLC)(nil)

// Decode deserializes a serialized UpdateFulfillHTLC message stored in the passed
// io.Reader observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *UpdateFulfillHTLC) Decode(r io.Reader, pver uint32) error {
	return readElements(r,
		&c.ChanID,
		&c.ID,
		c.PaymentPreimage[:],
		&c.Marked,
	)
}

// Encode serializes the target UpdateFulfillHTLC into the passed io.Writer
// observing the protocol version specified.
//
// This is part of the lnwire.Message interface.
func (c *UpdateFulfillHTLC) Encode(w io.Writer, pver uint32) error {
	return writeElements(w,
		c.ChanID,
		c.ID,
		c.PaymentPreimage[:],
		c.Marked,
	)
}

// MsgType returns the integer uniquely identifying this message type on the
// wire.
//
// This is part of the lnwire.Message interface.
func (c *UpdateFulfillHTLC) MsgType() MessageType {
	return MsgUpdateFulfillHTLC
}

// MaxPayloadLength returns the maximum allowed payload size for an UpdateFulfillHTLC
// complete message observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *UpdateFulfillHTLC) MaxPayloadLength(uint32) uint32 {
	// 32 + 8 + 32 + 4
	return 76
}
