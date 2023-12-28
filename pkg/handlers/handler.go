package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/karthikkalarikal/assignmentOnepane/pkg/usecase"
)

var mutex sync.Mutex

func HandleCombinedResults(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	comments, posts, users := usecase.FetchDataConcurrently()

	response := usecase.CombineResults(comments, posts, users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
