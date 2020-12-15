package api

import (
	"encoding/json"
	"github.com/cbhakar/mphasis/s3"
	"github.com/cbhakar/mphasis/store"
	"github.com/cbhakar/mphasis/utils"
	"log"
	"net/http"
)

var (
	imageFormatTable = []string{
		"image/jpg",
		"image/jpeg",
		"image/gif",
		"image/png",
		"image/bmp",
	}
	sortOrderList = []string{
		"asc",
		"desc",
	}
	sortByList = []string{
		"created_at",
		"image_name",
	}
)

func AddImage(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("file")
	if err != nil {
		err := map[string]interface{}{"message": "enter add file in form data"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer file.Close()
	if !utils.StrInListStatus(handler.Header.Get("Content-Type"), imageFormatTable) {
		err := map[string]interface{}{"message": "please upload a image"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = s3.UploadFileToS3(handler.Filename, file)
	if err != nil {
		err := map[string]interface{}{"message": err.Error()}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	success := map[string]interface{}{"message": "image uploaded successfully"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
	return
}
func GetImages(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var qParms store.QueryDetails
	var ok bool
	imageId := queryParams.Get("id")
	if qParms.ImageId, ok = utils.CheckIntValue(imageId); !ok && imageId != "" {
		err := map[string]interface{}{"message": "id should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	page := queryParams.Get("page")
	if qParms.Page, ok = utils.CheckIntValue(page); !ok && page != "" {
		err := map[string]interface{}{"message": "page should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	size := queryParams.Get("size")
	if qParms.Size, ok = utils.CheckIntValue(size); !ok && size != "" {
		err := map[string]interface{}{"message": "size should be a integer value"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	qParms.Order = queryParams.Get("order")
	if ok = utils.StrInListStatus(qParms.Order, sortOrderList); !ok && qParms.Order != "" {
		err := map[string]interface{}{"message": "enter a valid sort order parameter"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	qParms.Sort = queryParams.Get("sort")
	if ok = utils.StrInListStatus(qParms.Sort, sortByList); !ok && qParms.Sort != "" {
		err := map[string]interface{}{"message": "enter a valid Sort by parameter"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	resp, err := store.GetImages(qParms)
	if err != nil {
		err := map[string]interface{}{"message": "error fetching data from db"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	success := map[string]interface{}{"result": resp}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
	return
}

func Stop() (err error) {
	err = store.CloseDbConn()
	if err != nil {
		log.Println("error : ", err.Error())
	}
	return
}
