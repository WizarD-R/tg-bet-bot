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
	// b.SendText("без полезное хуйня")

	// 1 привет как дела
	// 2 определить админ ли юзер
	// 3 меняем ui кнопок на доступный фун-ции взависимости от админа или не админа



	// // b.SendText("регистарция пользователя")
	// // 1 загрузка о пользователе в бд через сервис
	// // 2 поменять
}

// admin
func (h *Handler) CreateEvent() {
	// 1 получение списка в формате [{name: navi-faze, date: 21.01.2000T16:30}]
	// 2 парсинг списка
	// 3 добавление в бд списка событий

	// формат события id name date result 
}

//adimn
func (h *Handler) AddResult() {

}

//admin
func (h *Handler) FinishTournament() {

}

// user
func (h *Handler) MyPredictions() {

}

// user
func (h *Handler) MakePrediction() {

}
