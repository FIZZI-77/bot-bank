package messages

const MsgHelp = `Я бот для транзакций, а могу переводить деньги на счет разным пользователям

Вот мои команды:

/send - перевести пользователю
/top-up - пополнить баланс`

const MsgHello = "Hi there! 👾\n\n" + MsgHelp

const (
	MsgUnknownCommand = "Unknown command 🤔"
	MsgBalance        = "Ваш баланс 💰: "
)
