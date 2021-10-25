package coupon

import (
	"fmt"
)

type CouponService interface {
	Describe(CouponID uint64) (*Coupon, error)
	List(cursor uint64, limit uint64) ([]Coupon, error)
	Create(Coupon) (uint64, error)
	Update(CouponID uint64, coupon Coupon) error
	Remove(CouponID uint64) (bool, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) ([]Coupon, error) {
	if len(couponList) == 0 {
		return nil, nil
	}
	if cursor >= uint64(len(couponList)) {
		return nil, fmt.Errorf("coupon index is wrong")
	}
	if cursor+limit >= uint64(len(couponList)) {
		return couponList[cursor:], nil
	} else {
		return couponList[cursor : cursor+limit], nil
	}
}

func (s *Service) Describe(couponID uint64) (*Coupon, error) {
	if couponID >= uint64(len(couponList)) {
		return nil, fmt.Errorf("coupon index is wrong")
	}
	return &couponList[couponID], nil
}

func (s *Service) Create(coupon Coupon) (uint64, error) {
	if coupon.Percent < 0 {
		return 0, fmt.Errorf(": need positive percent value (%d)", coupon.Percent)
	}
	couponList = append(couponList, coupon)
	return uint64(len(couponList)), nil
}

func (s *Service) Remove(CouponID uint64) (bool, error) {
	if CouponID >= uint64(len(couponList)) {
		return false, fmt.Errorf(": id range check error (%d)", CouponID)
	}
	couponList = append(couponList[:CouponID], couponList[CouponID+1:]...)
	return true, nil
}

func (s *Service) Update(CouponID uint64, coupon Coupon) error {
	if CouponID >= uint64(len(couponList)) {
		return fmt.Errorf(": id range check error (%d)", CouponID)
	}
	couponList[CouponID].Code = coupon.Code
	couponList[CouponID].Percent = coupon.Percent
	return nil
}
