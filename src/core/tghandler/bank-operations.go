package tghandler

import "github.com/sirupsen/logrus"

func (h *Handler) sendMoney(username string, amount float64, recipient string) {
	err := h.service.Transaction.SendMoney(username, amount, recipient)
	if err != nil {
		logrus.Errorf("Ошибка в отправке денег: %s", err.Error())
	}
}

func (h *Handler) topUpMoney(username string, amount float64) {

	err := h.service.TopUpMoney.TopUpMoney(username, amount)
	if err != nil {
		logrus.Errorf("Ошибка в пополнении баланса: %s", err.Error())
	}
}

func (h *Handler) takeBalance(username string) float64 {
	balance, err := h.service.TakeBalance(username)
	if err != nil {
		logrus.Errorf("Ошибка в получении баланса: %s", err.Error())
	}
	return balance
}
