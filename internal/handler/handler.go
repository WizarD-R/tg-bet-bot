package handler

type Handler struct {
	Service Service
}

type Service interface {
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Start() {

}

func (h *Handler) CreateEvent() {

}

func (h *Handler) AddResult() {

}

func (h *Handler) FinishTournament() {

}

func (h *Handler) MyPredictions() {

}

func (h *Handler) MakePrediction() {

}

// case update.Message.Text == "/start":
// 			// msg := handlers.HandleStart(update.Message)
// 			// bot.Send(msg)
// 			// // другие команды...
// 		case update.Message.Text == "/create-event":
// 			// ds
// 		case update.Message.Text == "/add-result":
// 			//ds
// 		case update.Message.Text == "/finish-tournament":
// 			// ыв
// 		case update.Message.Text == "/my-predictions":
// 			//ds
// 		case strings.Contains(update.Message.Text, "/match"):
