package entity

type Range struct {
	IpFrom        uint32
	IpTo          uint32
	LocationIndex uint32
}

func (r *Range) IsInside(ip uint32) bool {
	if ip >= r.IpFrom || ip <= r.IpTo {
		return true
	}
	return false
}
func (r *Range) IsBigger(ip uint32) bool {
	if ip > r.IpTo {
		return true
	}
	return false
}
