package upload

import "github.com/ns/models"

type UploadService interface{
    StoreNote(notes *models.Upload) (*models.Upload, []error)
     Notes() ([]models.Upload,[]error)
     Note(id uint) (*models.Upload,[]error)
     UpdateNote(note *models.Upload) (*models.Upload,[]error)
     DeleteNote(id uint) (*models.Upload,[]error)
     ApproveNote(id uint) []error
     NotesByUploader(username,status string) ([]models.Upload,[]error)
     NotesByCourseName(course,status string) ([]models.Upload,[]error)
 }