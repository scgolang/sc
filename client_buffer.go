package sc

import (
	"github.com/pkg/errors"
	"github.com/scgolang/osc"
)

// AllocBuffer allocates a buffer on the server
func (c *Client) AllocBuffer(frames, channels int) (*Buffer, error) {
	buf, err := c.sendBufAllocMsg(frames, channels)
	if err != nil {
		return nil, err
	}
	if err := c.awaitBufAllocReply(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// awaitBufAllocReply waits for a reply to /b_alloc
func (c *Client) awaitBufAllocReply(buf *Buffer) error {
	var done osc.Message
	select {
	case done = <-c.doneChan:
	case err := <-c.errChan:
		return err
	}
	// error if this message was not an ack of /b_alloc
	if numargs := len(done.Arguments); numargs != 2 {
		return errors.Errorf("expected two arguments to /done message, got %d", numargs)
	}
	addr, err := done.Arguments[0].ReadString()
	if err != nil {
		return err
	}
	if addr != bufferAllocAddress {
		c.doneChan <- done

	}
	bufnum, err := done.Arguments[1].ReadInt32()
	if err != nil {
		return err
	}
	if bufnum != buf.Num {
		c.doneChan <- done
	}
	return nil
}

// awaitBufReadReply waits for a reply to /b_allocRead
func (c *Client) awaitBufReadReply(buf *Buffer) error {
	var done osc.Message
	select {
	case done = <-c.doneChan:
	case err := <-c.errChan:
		return err
	}

	// error if this message was not an ack of the buffer read
	if len(done.Arguments) != 2 {
		return errors.New("expected two arguments to /done message")
	}
	addr, err := done.Arguments[0].ReadString()
	if err != nil {
		return err
	}
	if addr != bufferReadAddress {
		c.doneChan <- done
	}
	bufnum, err := done.Arguments[1].ReadInt32()
	if err != nil {
		return err
	}
	if bufnum != buf.Num {
		c.doneChan <- done
	}
	return nil
}

// QueryBuffer gets information about a buffer from scsynth.
func (c *Client) QueryBuffer(num int32) (*Buffer, error) {
	if err := c.oscConn.Send(osc.Message{
		Address: bufferQueryAddress,
		Arguments: osc.Arguments{
			osc.Int(num),
		},
	}); err != nil {
		return nil, errors.Wrap(err, "sending buffer query message")
	}
	return c.awaitBufInfoReply()
}

// awaitBufInfoReply waits for a reply to /b_allocRead
func (c *Client) awaitBufInfoReply() (*Buffer, error) {
	var bufinfo osc.Message
	select {
	case bufinfo = <-c.bufferInfoChan:
	case err := <-c.errChan:
		return nil, err
	}
	// error if this message was not an ack of the buffer read
	if numargs := len(bufinfo.Arguments); numargs != 4 {
		return nil, errors.Errorf("expected four arguments to /b_info message, got %d", numargs)
	}
	bufnum, err := bufinfo.Arguments[0].ReadInt32()
	if err != nil {
		return nil, err
	}
	frames, err := bufinfo.Arguments[1].ReadInt32()
	if err != nil {
		return nil, err
	}
	channels, err := bufinfo.Arguments[2].ReadInt32()
	if err != nil {
		return nil, err
	}
	sampleRate, err := bufinfo.Arguments[3].ReadFloat32()
	if err != nil {
		return nil, err
	}
	return &Buffer{
		Channels:   channels,
		Frames:     frames,
		Num:        bufnum,
		SampleRate: sampleRate,
	}, nil
}

// ReadBuffer tells the server to read an audio file and load it into a buffer.
func (c *Client) ReadBuffer(path string, num int32, channels ...int) (*Buffer, error) {
	buf, err := c.sendBufReadMsg(path, num, channels...)
	if err != nil {
		return nil, err
	}
	if err := c.awaitBufReadReply(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// sendBufAllocMsg sends a /b_alloc message
func (c *Client) sendBufAllocMsg(frames, channels int) (*Buffer, error) {
	buf := &Buffer{client: c}
	msg := osc.Message{
		Address: bufferAllocAddress,
		Arguments: osc.Arguments{
			osc.Int(buf.Num),
			osc.Int(int32(frames)),
			osc.Int(int32(channels)),
		},
	}
	if err := c.oscConn.Send(msg); err != nil {
		return nil, err
	}
	return buf, nil
}

// sendBufReadMsg sends a /b_allocRead command.
func (c *Client) sendBufReadMsg(path string, num int32, channels ...int) (*Buffer, error) {
	buf := newReadBuffer(path, num, c)

	var addr string
	if len(channels) == 0 {
		addr = bufferReadAddress
	} else {
		addr = bufferReadChannelAddress
	}
	msg := osc.Message{
		Address: addr,
		Arguments: osc.Arguments{
			osc.Int(buf.Num),
			osc.String(path),
		},
	}
	for _, channel := range channels {
		msg.Arguments = append(msg.Arguments, osc.Int(channel))
	}
	if err := c.oscConn.Send(msg); err != nil {
		return nil, err
	}
	return buf, nil
}
