package web

import (
	"github.com/labstack/echo/v4"
	"github.com/madlemon/estim8/internal/model"
	"github.com/madlemon/estim8/web/template"
	"log"
	"net/http"
)

func CreateRoom(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Redirect(c, "/", http.StatusUnauthorized)
	}

	user := model.User(authCookie.Value)

	modeName := c.QueryParam("mode")
	mode, err := model.MakeEstimationMode(modeName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	roomId, newRoom := model.NewRoom(mode)
	newRoom.AddUser(user)

	component := template.Room(Config.PublicUrl, roomId, *newRoom)
	c.Response().Header().Add("HX-Push-Url", "/rooms/"+roomId)
	log.Printf("New room created: %q\n", roomId)
	return Render(component, c)
}

func GetRoom(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Render(template.LoginView(), c)
	}

	id := c.Param("roomId")

	room, ok := model.Global.Rooms[id]
	if !ok {
		return Render(template.NotFound(), c)
	}

	user := model.User(authCookie.Value)
	room.AddUser(user)

	return Render(template.RoomView(Config.PublicUrl, id, *room), c)
}

func GetEstimates(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Render(template.LoginView(), c)
	}

	id := c.Param("roomId")

	room, ok := model.Global.Rooms[id]

	if !ok {
		return Render(template.RoomGone(), c)
	}
	return Render(template.EstimatesPollResponse(id, *room), c)
}

func SetEstimate(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Redirect(c, "/", http.StatusUnauthorized)
	}
	user := model.User(authCookie.Value)

	roomId := c.Param("roomId")

	room, ok := model.Global.Rooms[roomId]
	if !ok {
		log.Printf("Room not found: %q\n", roomId)
		return Render(template.RoomGone(), c)
	}

	if room.CurrentMode == model.TimeMode {
		formData := new(TimeEstimateDTO)
		if err := c.Bind(formData); err != nil {
			log.Fatal(err)
			return err
		}
		err := room.AddTimeEstimate(user, formData.Value)
		if err != nil {
			msg := "Cannot parse your estimate :("
			textInput := template.EstimatesAndTextInput(roomId, *room, formData.Value, &msg)
			return Render(textInput, c)
		}
	} else {
		formData := new(PointEstimateDTO)
		if err := c.Bind(formData); err != nil {
			return err
		}
		room.AddPointEstimate(user, formData.Value)
	}

	component := template.EstimatesAndTextInput(roomId, *room, "", nil)
	return Render(component, c)
}

type PointEstimateDTO struct {
	Value int `form:"value"`
}

type TimeEstimateDTO struct {
	Value string `form:"value"`
}

func SetResultVisibility(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Redirect(c, "/", http.StatusUnauthorized)
	}

	roomId := c.Param("roomId")

	room, ok := model.Global.Rooms[roomId]
	if !ok {
		log.Printf("Room not found: %q\n", roomId)
		return Render(template.RoomGone(), c)
	}

	formData := new(VisibilityFormData)
	if err := c.Bind(formData); err != nil {
		log.Fatal(err)
		return err
	}

	room.ResultsVisible = formData.ResultsVisible

	log.Printf("Visibility of estimates in room %q set to %t\n", roomId, room.ResultsVisible)

	component := template.Room(Config.PublicUrl, roomId, *room)
	return Render(component, c)
}

type VisibilityFormData struct {
	ResultsVisible bool `form:"resultsVisible"`
}

func DeleteEstimates(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Redirect(c, "/", http.StatusUnauthorized)
	}

	roomId := c.Param("roomId")

	room, ok := model.Global.Rooms[roomId]
	if !ok {
		log.Printf("Room not found: %q\n", roomId)
		return Render(template.NotFound(), c)
	}

	room.ClearEstimates()
	log.Printf("Estimates cleared in room: %q\n", roomId)
	log.Printf("ResultsVisible: %b", room.ResultsVisible)
	component := template.Room(Config.PublicUrl, roomId, *room)
	return Render(component, c)
}
