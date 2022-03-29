package main

import (
	"log"
	"simple-upload-file/config"

	"github.com/spf13/viper"
)

// type author struct {
// 	Name  string `json:"name"`
// 	Image string `json:"image"`
// }

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	if viper.GetBool("debug") {
		log.Println("Service Run On DEBUG mode")
	}
}

func main() {
	config.Run()
	// http.HandleFunc("/", routeGetIndex)
	// http.HandleFunc("/process", routeSubmit)
	// http.HandleFunc("/process", routeSubmitRaw)

	// port := 8080
	// log.Println("Server Starting at localhost:", port)
	// http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// func routeGetIndex(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, "Server Internal Error", http.StatusBadRequest)
// 	}

// 	// tmpl := template.Must(template.ParseFiles("view.html"))
// 	// err := tmpl.Execute(w, nil)
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// }
// 	json.NewEncoder(w).Encode("Selamat Datang Di Beranda")
// }

// func routeSubmit(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != "POST" {
// 		http.Error(w, "", http.StatusBadRequest)
// 		return
// 	}

// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	alias := r.FormValue("alias")
// 	log.Println("INI ALIAS", alias)
// 	uploadedFile, handler, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer uploadedFile.Close()

// 	dir, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	filename := handler.Filename
// 	if alias != "" {
// 		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
// 	}

// 	fileLocation := filepath.Join(dir, "files", filename)
// 	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer targetFile.Close()

// 	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("done"))
// }

// func routeSubmitRaw(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("EndPoint %v", r.RequestURI)
// 	var author author

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	imgSplit := strings.Split(author.Image, ",")
// 	// fmt.Println("INI IMAGE ", imgSplit[1])

// 	// detect ext img
// 	var extImg string
// 	imgType := imgSplit[0]
// 	switch imgType {
// 	case "data:image/jpeg;base64":
// 		extImg += "jpeg"
// 	case "data:image/png;base64":
// 		extImg += ".png"
// 	}

// 	bytesImage, err := base64.StdEncoding.DecodeString(imgSplit[1])
// 	if err != nil {
// 		fmt.Printf("Error Decode image %v", err.Error())
// 		return
// 	}

// 	dir, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	time := time.Now().Format("20060102150405")
// 	filename := "image" + time + extImg

// 	fileLocation := filepath.Join(dir, "files", filename)
// 	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer targetFile.Close()

// 	_, err = targetFile.Write(bytesImage)
// 	if err != nil {
// 		fmt.Printf("Error Write file %s", err.Error())
// 	}

// 	targetFile.Sync()

// 	if _, err := io.Copy(targetFile, strings.NewReader(filename)); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte("done"))
// }
