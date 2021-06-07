package handlers

import (
	"net/http"
	"github.com/labstack/echo"

	"shortener/logic"
)

type ShortenerHandler struct {
	logic logic.ShortenerLogic
}


func NewShortenerHandler(e *echo.Echo, l logic.ShortenerLogic){
	handler := &ShortenerHandler{
		logic: l,
	}
	e.GET("/a/", handler.CreateKey)
	e.GET("/s/*", handler.UseKey)
}



func (h *ShortenerHandler) CreateKey(c echo.Context) (err error){
	ctx := c.Request().Context()
	addr := c.Echo().Server.Addr
	long := c.QueryParam("url")
	short, err  := h.logic.CreatePair(ctx, long, addr)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, short)
}


func (h *ShortenerHandler) UseKey(c echo.Context) (err error){
	ctx := c.Request().Context()
	url :=c.Request().URL.String()
	short := url[len(url)-8:]
	parent, err := h.logic.GetLongUrl(ctx, short)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}
	if parent == ""{
		return c.JSON(http.StatusOK, "Не удалось найти первоначальный адрес по короткому варианту")
	}
	c.Redirect(302, parent)
	return
}
