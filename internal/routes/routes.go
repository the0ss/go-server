package routes

import (
	"backend/controller"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseBody struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

type RequestBody struct {
	ToSort [][]int `json:"to_sort"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandle(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIserver struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/process-single", makeHTTPHandle(s.processSingleHandler)).Methods("POST")
	router.HandleFunc("/process-concurrent", makeHTTPHandle(s.processConcurrentHandler)).Methods("POST")
	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIserver) processSingleHandler(w http.ResponseWriter, r *http.Request) error {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return WriteJSON(w, http.StatusBadRequest, "Invalid Request body")
	}
	sortedArrays, elapsedTime := controller.ProcessSequential(requestBody.ToSort)

	response := ResponseBody{
		SortedArrays: sortedArrays,
		TimeNs:       elapsedTime,
	}

	return WriteJSON(w, http.StatusOK, response)
}

func (s *APIserver) processConcurrentHandler(w http.ResponseWriter, r *http.Request) error {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return WriteJSON(w, http.StatusBadRequest, "Invalid Request body")
	}
	sortedArrays, elapsedTime := controller.ProcessConcurrent(requestBody.ToSort)

	response := ResponseBody{
		SortedArrays: sortedArrays,
		TimeNs:       elapsedTime,
	}

	return WriteJSON(w, http.StatusOK, response)
}
