package packet

type craftRequest struct {
	h Header
}

func New() *craftRequest {
	return &craftRequest{}
}


func (r *craftRequest) clone() craftRequest {
	return *r
}
func (r *craftRequest) WithDUP() *craftRequest {
	c := r.clone()
	c.h = *c.h.WithDUP(true)
	return &c
}
func (r *craftRequest) WithoutDUP() *craftRequest {
	c := r.clone()
	c.h = *c.h.WithDUP(false)
	return &c
}

func (r *craftRequest) WithRETAIN() *craftRequest {
	c := r.clone()
	c.h = *c.h.WithRETAIN(true)
	return &c
}

func (r *craftRequest) WithoutRETAIN() *craftRequest {
	c := r.clone()
	c.h = *c.h.WithRETAIN(false)
	return &c
}

func (r *craftRequest) WithQoS(q byte) *craftRequest {
	c := r.clone()
	c.h = *c.h.WithQoS(q)
	return &c
}

func (r *craftRequest) Connect() *Connect {
	r.h.packetType = CONNECT
	return &Connect{
		header: r.h,
	}
}

func (r *craftRequest) Connack() *Connack {
	r.h.packetType = CONNACK
	return &Connack{
		header: r.h,
	}
}

func (r *craftRequest) Subscribe() *Subscribe {
	r.h.packetType = SUBSCRIBE
	return &Subscribe{
		header: r.h,
	}
}


func (r *craftRequest) Suback() *Suback {
	r.h.packetType = SUBACK
	return &Suback{
		header: r.h,
	}
}

func (r *craftRequest) Publish() *Publish {
	r.h.packetType = PUBLISH
	return &Publish{
		header: r.h,
	}
}

func (r *craftRequest) PingResp() *Pingresp {
	r.h.packetType = PINGRESP
	return &Pingresp{
		header: r.h,
	}
}

func (r *craftRequest) PubAck() *Puback {
	r.h.packetType = PUBACK
	return &Puback{
		header: r.h,
	}
}

func (r *craftRequest) Unsuback() *Unsuback {
	r.h.packetType = UNSUBACK
	return &Unsuback{
		header: r.h,
	}
}

