package author

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"papvan/cvmaker/internal/apperror"
	service2 "papvan/cvmaker/internal/author/service"
	"papvan/cvmaker/internal/handlers"
	"papvan/cvmaker/pkg/api/sorting"
	"papvan/cvmaker/pkg/logging"
)

const (
	authorsURL = "/authors"
	authorURL  = "/author/:uuid"
)

type handler struct {
	logger  *logging.Logger
	service *service2.Service
}

func NewHandler(service *service2.Service, logger *logging.Logger) handlers.Handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, sorting.Middleware(apperror.Middleware(h.GetList), "created_at ", sorting.SortASC))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	var sortOptions sorting.Options
	if options, ok := r.Context().Value(sorting.OptionsContextKey).(sorting.Options); ok {
		sortOptions = options
	}

	all, err := h.service.GetAll(r.Context(), sortOptions)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}
