package interfaces

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"watchWebsite/application"
	"watchWebsite/domain/entity"
	"watchWebsite/infrastructure/auth"
	"watchWebsite/notifications"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Watchs struct {
	w application.WatchAppInterface
}

func NewWatchs(w application.WatchAppInterface) *Watchs {
	return &Watchs{
		w: w,
	}
}

func (w *Watchs) GetAllWatch(c echo.Context) error {
	cookie, err := c.Cookie("token")
	notices := make(map[string]string)
	watchesCopy := []struct {
		WatchName string
		Price     uint32
		Image     string
		Quantity  uint32
	}{}
	verifications := struct {
		IsCustomer bool `json:"isCustomer"`
		Watches    []struct {
			WatchName string
			Price     uint32
			Image     string
			Quantity  uint32
		} `json:"watches"`
	}{}
	if err != nil {
		if err.Error() == "http: named cookie not present" {
			watches, err := w.w.GetAllWatch()
			if err != nil {
				notices := make(map[string]string)
				notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
				return c.JSON(http.StatusInternalServerError, notices)
			}

			for _, value := range watches {
				watch := struct {
					WatchName string
					Price     uint32
					Image     string
					Quantity  uint32
				}{}
				watch.WatchName = strings.TrimSpace(value.WatchName)
				watch.Price = *value.Price
				watch.Image = strings.TrimSpace(value.Image)
				watch.Quantity = *value.Quantity
				watchesCopy = append(watchesCopy, watch)
			}
			verifications.IsCustomer = false
			verifications.Watches = watchesCopy
			return c.JSON(http.StatusContinue, verifications)
		}
	}
	token, err := auth.IsTokenValid(cookie.Value)

	if err != nil {
		if err == echo.ErrUnauthorized {
			notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.EXPIRED_TOKEN_ERROR
			return c.JSON(http.StatusUnauthorized, "notices")
		} else if err == echo.ErrBadRequest {
			notices[notifications.BAD_REQUEST_ERROR] = notifications.EXPIRED_TOKEN_ERROR
			return c.JSON(http.StatusUnauthorized, "notices")
		}
	}

	claims := token.Claims.(jwt.MapClaims)

	isCustomer := claims["isCustomer"].(bool)
	if !isCustomer {
		notices[notifications.UNAUTHORIZED_ACCESS_ERROR] = notifications.EXPIRED_TOKEN_ERROR
		return c.JSON(http.StatusUnauthorized, "notices")
	}

	watches, err := w.w.GetAllWatch()
	if err != nil {
		notices := make(map[string]string)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	for _, value := range watches {
		watch := struct {
			WatchName string
			Price     uint32
			Image     string
			Quantity  uint32
		}{}
		watch.WatchName = strings.TrimSpace(value.WatchName)
		watch.Price = *value.Price
		watch.Image = strings.TrimSpace(value.Image)
		watch.Quantity = *value.Quantity
		watchesCopy = append(watchesCopy, watch)
	}
	verifications.IsCustomer = true
	verifications.Watches = watchesCopy
	return c.JSON(http.StatusOK, verifications)
}

func (w *Watchs) GetAllWatchs(c echo.Context) error {
	watches, err := w.w.GetAllWatch()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, watches)
}

func (w *Watchs) GetWatchByNameRelative(c echo.Context) error {
	notices := make(map[string]string)
	name := c.QueryParam("name")

	if err := c.Bind(&name); err != nil {
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	var ws []*entity.Watch
	if name == "" {
		fmt.Println(name)
		notices[notifications.BAD_REQUEST_ERROR] = notifications.EMPTY_NAME_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}
	ws, err := w.w.GetWatchByNameRelative(name)
	watches := []struct {
		WatchName string
		Price     uint32
		Image     string
		Quantity  uint32
	}{}

	for _, v := range ws {
		watch := struct {
			WatchName string
			Price     uint32
			Image     string
			Quantity  uint32
		}{}
		watch.WatchName = v.WatchName
		watch.Price = *v.Price
		watch.Image = v.Image
		watch.Quantity = *v.Quantity
		watches = append(watches, watch)
	}

	if err != nil {
		fmt.Println(err)
		notices[notifications.INTERNAL_SERVER_ERROR] = notifications.SYSTEM_ERROR
		return c.JSON(http.StatusInternalServerError, notices)
	}

	return c.JSON(http.StatusOK, watches)
}

func (w *Watchs) UpdateWatchCus(c echo.Context) error {

	watch := &entity.Watch{}
	err := c.Bind(&watch)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var decodedImage []byte
	var imagePath, filename string
	if watch.Image == "" {
		watch.Image = "-1"
	} else {
		decodedImage, err = base64.StdEncoding.DecodeString(watch.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		imagePath = "D:/DoAnThucTap/Front-end/assets/img"

		filename = "image" + strconv.FormatUint(uint64(*watch.WatchId), 10) + ".png" // Change the extension based on the actual image type
		watch.Image = "assets/img/" + filename
		// Write the decoded image data to the file

	}

	watchTmp, err := w.w.GetWatchByName(watch.WatchName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = w.w.UpdateWatchCus(watch)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if imagePath != "" {
		oldImagePath := filepath.Join(imagePath, strings.Split(watchTmp.Image, "/")[2])

		err = os.Remove(oldImagePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = os.MkdirAll(imagePath, os.ModePerm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = os.WriteFile(filepath.Join(imagePath, filename), decodedImage, 0644)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, "")
}

func (w *Watchs) CreateWatch(c echo.Context) error {
	watch := &entity.Watch{}

	var tmp uint32 = 0
	var status int32 = -1

	watch.Price = &tmp
	watch.TypeOfWatchId = &tmp
	watch.BrandId = &tmp

	err := c.Bind(&watch)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	watchID, err := w.w.CreateWatch(watch)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	watch.WatchId = watchID
	watch.Status = &status

	decodedImage, err := base64.StdEncoding.DecodeString(watch.Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	imagePath := "D:/DoAnThucTap/Front-end/assets/img"
	err = os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	filename := "image " + strconv.FormatUint(uint64(*watchID), 10) + ".png" // Change the extension based on the actual image type
	watch.Image = "assets/img/" + filename
	// Write the decoded image data to the file
	err = os.WriteFile(filepath.Join(imagePath, filename), decodedImage, 0644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = w.w.UpdateWatchCus(watch)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func (w *Watchs) DeleteWatch(c echo.Context) error {
	watchIDString := c.Param("watchID")

	watchIDInt, err := strconv.Atoi(watchIDString)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	watchID := uint32(watchIDInt)

	err = w.w.DeleteWatch(&watchID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
