package handler

import (
	"app/api/application/interactor"
	"app/api/presentation/request"
	"app/api/presentation/response"
	"net/http"

	"github.com/pkg/errors"
)

type TagHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

type tagHandler struct {
	tagInteractor      interactor.TagInteractor
	categoryInteractor interactor.CategoryInteractor
}

func NewTagHandler(ti interactor.TagInteractor, ci interactor.CategoryInteractor) TagHandler {
	return &tagHandler{
		tagInteractor:      ti,
		categoryInteractor: ci,
	}
}

func (th *tagHandler) Create(w http.ResponseWriter, r *http.Request) {
	src, err := ReadRequestBody(r, &request.CreateTagRequest{})
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to read request"), "failed to read request")
		return
	}
	req, _ := src.(*request.CreateTagRequest)
	err = req.Validation(th.categoryInteractor)
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to validation"), err.Error())
		return
	}

	tag, err := th.tagInteractor.Create(req.Tag, req.CategoryID)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to create tag"), "failed to create tag")
		return
	}

	response.Success(w, response.ConvertToTagResponse(tag))
}

func (th *tagHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tags, err := th.tagInteractor.GetAll()
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to get"), "failed to get tags")
		return
	}
	response.Success(w, response.ConvertToTagsResponse(tags))
}

func (th *tagHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := ReadPathParam(r, "id")
	tag, err := th.tagInteractor.GetByID(id)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to get tag"), "failed to get tag")
		return
	}
	response.Success(w, response.ConvertToTagResponse(tag))
}
