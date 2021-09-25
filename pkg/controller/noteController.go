package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horvatic/zracni-udar-service/pkg/service"
)

type NoteController interface {
	GetNotesByProjectId(w http.ResponseWriter, req *http.Request)
	GetNote(w http.ResponseWriter, req *http.Request)
	CreateNote(w http.ResponseWriter, req *http.Request)
	UpdateNote(w http.ResponseWriter, req *http.Request)
	DeleteNote(w http.ResponseWriter, req *http.Request)
}

type noteController struct {
	noteService service.NoteService
}

func BuildNoteController(noteService service.NoteService) NoteController {
	return &noteController{
		noteService: noteService,
	}
}

func (nc *noteController) GetNotesByProjectId(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	notes, errorType, err := nc.noteService.GetNotesByProjectId(vars["id"])
	sendJson(w, notes, errorType, err)
}

func (nc *noteController) GetNote(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	notes, errorType, err := nc.noteService.GetNote(vars["projectId"], vars["noteId"])
	sendJson(w, notes, errorType, err)
}

func (nc *noteController) CreateNote(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := nc.noteService.CreateNote(vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (nc *noteController) UpdateNote(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := nc.noteService.UpdateNote(vars["noteId"], vars["projectId"], &req.Body)
	sendCreateResult(w, errorType, err)
}

func (nc *noteController) DeleteNote(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	errorType, err := nc.noteService.DeleteNote(vars["noteId"], vars["projectId"])
	sendCreateResult(w, errorType, err)
}
