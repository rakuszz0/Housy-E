package handlers

import (
	housesdto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	// "context"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
)

var path_file = "http://localhost:5000/uploads/"

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}

func (h *handlerHouse) FindHouses(c echo.Context) error {
	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for i, p := range houses {
		houses[i].Image = path_file + p.Image
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerHouse) FindHousesFilter(c echo.Context) error {
	querParams := c.QueryParams()

	price, _ := strconv.Atoi(querParams.Get("price"))

	houses, err := h.HouseRepository.FindHousesFilter(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// for i, p := range houses {
	// 	houses[i].Image = path_file + p.Image
	// }

	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerHouse) GetHouse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	house.Image = path_file + house.Image

	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseHouse(house)}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerHouse) CreateHouse(c echo.Context) error {
	dataContex := c.Get("dataFile") // add this code
	filename := dataContex.(string) // add this code

	price, _ := strconv.Atoi(c.FormValue("price"))
	bedroom, _ := strconv.Atoi(c.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(c.FormValue("bathroom"))
	request := housesdto.HouseRequest{
		Name:      c.FormValue("name"),
		CityName:  c.FormValue("city_name"),
		Address:   c.FormValue("address"),
		TypeRent:  c.FormValue("type_rent"),
		Amenities: datatypes.JSON(c.FormValue("amenities")),
		Price:     price,
		Bedroom:   bedroom,
		Bathroom:  bathroom,

		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	house := models.House{
		Name:        request.Name,
		CityName:    request.CityName,
		Address:     request.Address,
		Price:       request.Price,
		TypeRent:    request.TypeRent,
		Amenities:   request.Amenities,
		Bedroom:     request.Bedroom,
		Bathroom:    request.Bathroom,
		Image:       filename,
		Description: request.Description,
	}

	house, err = h.HouseRepository.CreateHouse(house)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	house, _ = h.HouseRepository.GetHouse(house.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: house})
}

func (h *handlerHouse) UpdateHouse(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	dataContext := c.Get("dataFile").(string)

	price, _ := strconv.Atoi(c.FormValue("price"))
	bedroom, _ := strconv.Atoi(c.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(c.FormValue("bathroom"))
	request := housesdto.HouseRequest{
		Name:        c.FormValue("name"),
		CityName:    c.FormValue("city_name"),
		Address:     c.FormValue("address"),
		TypeRent:    c.FormValue("type_rent"),
		Amenities:   datatypes.JSON(c.FormValue("amenities")),
		Price:       price,
		Bedroom:     bedroom,
		Bathroom:    bathroom,
		Image:       dataContext,
		Description: c.FormValue("description"),
	}

	id, _ := strconv.Atoi(c.Param("id"))
	house, err := h.HouseRepository.GetHouse(int(id))
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.Name != "" {
		house.Name = request.Name
	}

	if request.CityName != "" {
		house.CityName = request.CityName
	}

	if request.Address != "" {
		house.Address = request.Address
	}

	if request.Price != 0 {
		house.Price = request.Price
	}

	if request.TypeRent != "" {
		house.TypeRent = request.TypeRent
	}

	if request.Amenities != nil {
		house.Amenities = request.Amenities
	} else {
		// set empty JSON object if request.Amenities is nil
		house.Amenities = datatypes.JSON("{}")
	}

	if request.Bedroom != 0 {
		house.Bedroom = request.Bedroom
	}

	if request.Bathroom != 0 {
		house.Bathroom = request.Bathroom
	}

	if request.Image != "" {
		house.Image = request.Image
	}

	if request.Description != "" {
		house.Description = request.Description
	}

	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerHouse) DeleteHouse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponseHouse(u models.House) models.HouseResponse {
	return models.HouseResponse{
		ID:          u.ID,
		Name:        u.Name,
		CityName:    u.CityName,
		Address:     u.Address,
		Price:       u.Price,
		TypeRent:    u.TypeRent,
		Amenities:   u.Amenities,
		Description: u.Description,
		Area:        u.Area,
		Bedroom:     u.Bedroom,
		Bathroom:    u.Bathroom,
		Image:       u.Image,
	}
}
