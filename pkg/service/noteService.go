package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/horvatic/zracni-udar-service/pkg/model"
	"github.com/horvatic/zracni-udar-service/pkg/store"
)

type NoteService interface {
	GetNotesByProjectId(id string) ([]model.Note, ErrorType, error)
	CreateNote(projectId string, body *io.ReadCloser) (ErrorType, error)
	UpdateNote(noteId string, projectId string, body *io.ReadCloser) (ErrorType, error)
	DeleteNote(noteId string, projectId string) (ErrorType, error)
	GetNote(projectId string, noteId string) (*model.Note, ErrorType, error)
}

type noteService struct {
	store store.Store
}

func BuildNoteService(store store.Store) NoteService {
	return &noteService{
		store: store,
	}
}

func (n *noteService) GetNotesByProjectId(id string) ([]model.Note, ErrorType, error) {
	var notes []model.Note
	project := n.store.GetProjectById(id)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}

	for _, p := range project.ProjectNotes {
		notes = append(notes, model.Note{
			Id:        p.Id,
			ProjectId: project.ProjectId,
			Name:      p.Name,
			Note:      p.Note,
		})
	}
	return notes, NoError, nil
}

func (n *noteService) GetNote(projectId string, noteId string) (*model.Note, ErrorType, error) {
	project := n.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return nil, BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, p := range project.ProjectNotes {
		if p.Id == noteId {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, BadRequest, errors.New("can not find note")
	}

	return &model.Note{
		Id:        project.ProjectNotes[index].Id,
		ProjectId: project.ProjectId,
		Name:      project.ProjectNotes[index].Name,
		Note:      project.ProjectNotes[index].Note,
	}, NoError, nil
}

func (n *noteService) CreateNote(projectId string, body *io.ReadCloser) (ErrorType, error) {
	var note model.Note
	err := json.NewDecoder(*body).Decode(&note)
	if err != nil {
		return JsonError, err
	}
	project := n.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	project.ProjectNotes = append(project.ProjectNotes, model.ProjectNotes{
		Id:   uuid.NewString(),
		Name: note.Name,
		Note: note.Note,
	})
	updateErr := n.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (n *noteService) UpdateNote(noteId string, projectId string, body *io.ReadCloser) (ErrorType, error) {
	var note model.Note
	err := json.NewDecoder(*body).Decode(&note)
	if err != nil {
		return JsonError, err
	}
	project := n.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectNotes {
		if pn.Id == noteId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find note")
	}
	project.ProjectNotes[index].Name = note.Name
	project.ProjectNotes[index].Note = note.Note
	updateErr := n.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}

func (n *noteService) DeleteNote(noteId string, projectId string) (ErrorType, error) {
	project := n.store.GetProjectById(projectId)
	if project == nil || project.ProjectId == "" {
		return BadRequest, errors.New("can not find project")
	}
	index := -1
	for i, pn := range project.ProjectNotes {
		if pn.Id == noteId {
			index = i
			break
		}
	}
	if index == -1 {
		return BadRequest, errors.New("can not find note")
	}
	project.ProjectNotes = append(project.ProjectNotes[:index], project.ProjectNotes[index+1:]...)

	updateErr := n.store.UpdateProject(project)
	if updateErr != nil {
		return DatabaseError, updateErr
	}

	return NoError, nil
}
