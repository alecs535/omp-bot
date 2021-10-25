package coupon

import (
	"github.com/ozonmp/omp-bot/internal/service/loyalty/coupon"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyCouponCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.SplitN(inputMessage.CommandArguments(), " ", 3)

	if len(args) < 3 {
		log.Println("wrong args: need used /edit__loyalty__coupon <id> <new_coupon_code> <new_percent>")
		return
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil || idx < 0 {
		log.Printf("wrong args, need exists ID of coupon (%s)", args[0])
		return
	}

	percent, err := strconv.Atoi(args[2])
	if err != nil || percent < 0 || percent > 100 {
		log.Printf("wrong args: need percent value in the range from 0 to 100 (%s)", args[2])
		return
	}

	if _, err := c.service.Describe(uint64(idx - 1)); err != nil {
		log.Printf("fail to locate coupon with ID %d: %v", idx, err)
		return
	}

	coupon := coupon.Coupon{Code: args[1], Percent: uint(percent)}

	var msg tgbotapi.MessageConfig
	if err := c.service.Update(uint64(idx-1), coupon); err != nil {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			err.Error(),
		)
	} else {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			coupon.String(),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCouponCommander.Edit: error sending reply message to chat - %v", err)
	}
}
