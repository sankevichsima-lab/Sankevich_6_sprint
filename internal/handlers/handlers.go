package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlRoot(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func HandleUploat(w http.ResponseWriter, r *http.Request){
	file, header, err := r.FormFile("myFile")
	if err!=nil{
		http.Error(w, "error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err!=nil{
		http.Error(w, "file read error", http.StatusInternalServerError)
		return
	}
	text := string(data)
	result, err := service.CheckMorse(text)
	if err!=nil{
		http.Error(w, "conversion error", http.StatusInternalServerError)
		return
	}

	originalExt := filepath.Ext(header.Filename)
	newFileName := time.Now().UTC().String() + originalExt

	newFile, err := os.Create(newFileName)
	if err!=nil{
		http.Error(w, "error creating file on disk", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_,err =newFile.WriteString(result)
	if err!=nil{
		http.Error(w, "error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}