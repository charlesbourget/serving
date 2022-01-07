package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func fileServer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	path := fmt.Sprintf("%s%s", *dir, strings.TrimPrefix(r.URL.Path, "/api"))

	isFile, err := checkIfFile(path)
	if err != nil {
		handleHttpError(http.StatusNotFound, "Directory not found", w)
		return
	}

	if isFile {
		http.ServeFile(w, r, path)
	} else {
		serveDir(w, r, path)
	}
}

func checkIfFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return !fileInfo.IsDir(), nil
}

func serveDir(w http.ResponseWriter, r *http.Request, path string) {
	srvAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)

	files, err := os.ReadDir(path)
	if err != nil {
		handleHttpError(http.StatusNotFound, "Directory not found", w)
		return
	}

	dirContent := &content{}

	for _, f := range files {
		info, err := f.Info()
		if err != nil {
			log.Fatal(err)
		}
		if f.IsDir() {
			dirContent.Directories = append(dirContent.Directories, directory{
				f.Name(),
				info.Mode().String(),
				fmt.Sprintf("http://%s%s/%s", srvAddr, strings.TrimSuffix(r.URL.Path, "/"), f.Name()),
			})
		} else {
			dirContent.Files = append(dirContent.Files, file{
				f.Name(),
				int(info.Size()),
				info.Mode().String(),
				fmt.Sprintf("http://%s%s/%s", srvAddr, strings.TrimSuffix(r.URL.Path, "/"), f.Name()),
			})
		}
	}

	switch *format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(dirContent)
	case "xml":
		w.Header().Set("Content-Type", "application/xml")
		err = xml.NewEncoder(w).Encode(dirContent)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func handleHttpError(status int, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	errorMessage := &httpError{
		status,
		message,
	}
	var err error
	switch *format {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(errorMessage)
	case "xml":
		w.Header().Set("Content-Type", "application/xml")
		err = xml.NewEncoder(w).Encode(errorMessage)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
