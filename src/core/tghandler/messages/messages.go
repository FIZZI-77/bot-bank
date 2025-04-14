package messages

const MsgHelp = `Я бот для транзакций, а могу переводить деньги на счет разным пользователям

Вот мои команды:

/send - перевести пользователю
/topup - пополнить баланс`

const MsgHello = "Hi there! 👾\n\n" + MsgHelp

const (
	MsgUnknownCommand = "Unknown command 🤔"
	MsgBalance        = "Ваш баланс 💰: "
	MsgTopUpMoney     = "Баланс успешно пополнен ✅"
	MsgSendMoney      = "Деньги успешно отправлены ✅"
	MsgErrorSend      = "Отправка денег не удалась, попробуйте снова 🚫"
	MsgErrorTopUp     = "Не удалось пополнить баланс, попробуйте снова 🚫"
	MsgUserNotExist   = "Операция не отправлена. Такого получателя нет в базе данных 🚫"
)
