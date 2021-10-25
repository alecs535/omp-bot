package coupon

import (
	"fmt"
)

type Coupon struct {
	Code    string
	Percent uint // 0..100
}

var couponList []Coupon

func (c *Coupon) String() string {
	return fmt.Sprintf("Coupon %s - %d%%", c.Code, c.Percent)
}
