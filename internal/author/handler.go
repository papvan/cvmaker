package author

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"papvan/cvmaker/internal/apperror"
	service2 "papvan/cvmaker/internal/author/service"
	"papvan/cvmaker/internal/handlers"
	"papvan/cvmaker/pkg/api/filter"
	"papvan/cvmaker/pkg/api/sorting"
	"papvan/cvmaker/pkg/logging"
	"strconv"
	"strings"
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
	// TODO change getting default parameters from config or something else
	router.HandlerFunc(http.MethodGet, authorsURL, filter.Middleware(sorting.Middleware(apperror.Middleware(h.GetList), "created_at ", sorting.SortASC), 10))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	filterOptions := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	name := r.URL.Query().Get("name")
	if name != "" {
		filterOptions.AddField("name", filter.OperatorLike, name, filter.DataTypeString)
	}
	age := r.URL.Query().Get("age")
	if age != "" {
		operator := filter.OperatorEq
		value := age
		if strings.Index(age, ":") != -1 {
			split := strings.Split(age, ":")
			operator = split[0]
			value = split[1]
		}
		err := filterOptions.AddField("age", operator, value, filter.DataTypeInt)
		if err != nil {
			return err
		}
	}

	isAlive := r.URL.Query().Get("is_alive")
	if isAlive != "" {
		_, err := strconv.ParseBool(isAlive)
		if err != nil {
			validationErr := apperror.BadRequestError("filter params validation failed", "bool value param wrong parameter")
			validationErr.WithParams(map[string]string{
				"is_alive": "this field should be boolean: true or false",
			})
			return validationErr
		}
		err = filterOptions.AddField("is_alive", filter.OperatorEq, isAlive, filter.DataTypeBool)
		if err != nil {
			return err
		}
	}
	createdAt := r.URL.Query().Get("created_at")
	if createdAt != "" {
		var operator string
		if strings.Index(createdAt, ":") != -1 {
			operator = filter.OperatorBetween
		} else {
			operator = filter.OperatorEq
		}
		err := filterOptions.AddField("created_at", operator, createdAt, filter.DataTypeDate)
		if err != nil {
			return err
		}
	}

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
