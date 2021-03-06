package lnwire

import "io"
import "time"

type UpdatePriceProbe struct {
	X_Remote uint64
	N_Remote uint64
	Q_Remote uint64
	I_Remote uint64		// Inflight_Remote
	Adiff_Remote time.Duration
	Sdiff_Remote time.Duration
	ChanID ChannelID
}

func NewUpdatePriceProbe() *UpdatePriceProbe {
	return &UpdatePriceProbe{}
}

//func NewUpdatePriceProbe(x_remote int32) *UpdatePriceProbe {
	//return &UpdatePriceProbe {
		//X_Remote:               x_remote,
	//}
//}

// A compile time check to ensure UpdatePriceProbe implements the lnwire.Message
// interface.
var _ Message = (*UpdatePriceProbe)(nil)

// Decode deserializes a serialized UpdatePriceProbe message stored in the passed
// io.Reader observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *UpdatePriceProbe) Decode(r io.Reader, pver uint32) error {
	return readElements(r,
		&c.ChanID,
		&c.X_Remote,
		&c.N_Remote,
		&c.Q_Remote,
		&c.I_Remote,
		&c.Adiff_Remote,
		&c.Sdiff_Remote,
	)
}

// Encode serializes the target UpdatePriceProbe into the passed io.Writer observing
// the protocol version specified.
//
// This is part of the lnwire.Message interface.
func (c *UpdatePriceProbe) Encode(w io.Writer, pver uint32) error {
	return writeElements(w,
		c.ChanID,
		c.X_Remote,
		c.N_Remote,
		c.Q_Remote,
		c.I_Remote,
		c.Adiff_Remote,
		c.Sdiff_Remote,
	)
}

// MsgType returns the integer uniquely identifying this message type on the
// wire.
//
// This is part of the lnwire.Message interface.
func (c *UpdatePriceProbe) MsgType() MessageType {
	return MsgUpdatePriceProbe
}

// MaxPayloadLength returns the maximum allowed payload size for an UpdatePriceProbe
// complete message observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (c *UpdatePriceProbe) MaxPayloadLength(uint32) uint32 {
	// FIXME:
	return 512
}
