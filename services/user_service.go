package services

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"simple-upload-file/models"
	"simple-upload-file/repositories"
	"strings"
	"time"
)

type IUserService interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User, id string) (*models.User, error)
	DeleteUser(id string) (int64, error)
}

type userService struct {
	db       *sql.DB
	userRepo repositories.IUserRepository
}

func NewUserService(db *sql.DB) IUserService {
	return &userService{
		db:       db,
		userRepo: repositories.NewUserRepo(db),
	}

}

func (u *userService) GetUser(id string) (*models.User, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileLocation := filepath.Join(dir, "files", user.Picture)
	targetFile, err := os.OpenFile(fileLocation, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer targetFile.Close()

	byteSlice, _ := ioutil.ReadFile(targetFile.Name())
	var base64Type string
	mimeType := http.DetectContentType(byteSlice)
	switch mimeType {
	case "image/jpeg":
		base64Type += "data:image/jpeg;base64,"
	case "image/png":
		base64Type += "data:image/png;base64,"
	}

	base64Type += base64.StdEncoding.EncodeToString(byteSlice)

	user.Picture = base64Type

	return user, nil
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {

	imgSplit := strings.Split(user.Picture, ",")

	// detect ext img
	var extImg string
	imgType := imgSplit[0]
	switch imgType {
	case "data:image/jpeg;base64":
		extImg += "jpeg"
	case "data:image/png;base64":
		extImg += ".png"
	}

	bytesImage, err := base64.StdEncoding.DecodeString(imgSplit[1])
	if err != nil {
		fmt.Printf("Error Decode image %v", err.Error())
		return nil, err
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	time := time.Now().Format("20060102150405")
	filename := "image" + time + extImg

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer targetFile.Close()

	_, err = targetFile.Write(bytesImage)
	if err != nil {
		fmt.Printf("Error Write file %s", err.Error())
	}

	targetFile.Sync()

	if _, err := io.Copy(targetFile, strings.NewReader(filename)); err != nil {
		return nil, err
	}

	user.Picture = filename
	user, err = u.userRepo.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUser(user *models.User, id string) (*models.User, error) {

	// getUser from repo
	newUser, err := u.userRepo.GetUser(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	imgSplit := strings.Split(user.Picture, ",")
	// detect ext img
	var extImg string
	imgType := imgSplit[0]
	switch imgType {
	case "data:image/jpeg;base64":
		extImg += "jpeg"
	case "data:image/png;base64":
		extImg += ".png"
	}
	// decode image base64 to []byte
	bytesImage, err := base64.StdEncoding.DecodeString(imgSplit[1])
	if err != nil {
		fmt.Printf("Error Decode image %v", err.Error())
		return nil, err
	}
	// get root directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	time := time.Now().Format("20060102150405")
	filename := "image" + time + extImg

	fileLocation := filepath.Join(dir, "files", newUser.Picture)
	targetFile, err := os.OpenFile(fileLocation, os.O_RDWR, 0666)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// close file and rename file
	defer func() {
		targetFile.Close()
		newFileLocation := filepath.Join(dir, "files", filename)
		err = os.Rename(fileLocation, newFileLocation)
		if err != nil {
			log.Println(err.Error())
		}
	}()

	// edit / write image file from []byte
	_, err = targetFile.Write(bytesImage)
	if err != nil {
		fmt.Printf("Error Write file %s", err.Error())
	}

	targetFile.Sync()

	// parsing to repo
	user.Picture = filename
	user, err = u.userRepo.UpdateUser(user, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteUser(id string) (int64, error) {
	idDelete, err := u.userRepo.DeleteUser(id)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return idDelete, nil
}
