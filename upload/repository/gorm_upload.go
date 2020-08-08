package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/ns/models"
)

type UploadRepositoryImpl struct{
    conn *gorm.DB
}

func NewUploadRepositoryImpl(db *gorm.DB) *UploadRepositoryImpl{
    return &UploadRepositoryImpl{conn:db}
}
func (upld *UploadRepositoryImpl) StoreNote(notes *models.Upload) (*models.Upload, []error) {
	note := notes
	errs := upld.conn.Create(note).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return note, errs
}
func (upld *UploadRepositoryImpl) Notes() ([]models.Upload,[]error){
        notes := []models.Upload{}
    	errs := upld.conn.Find(&notes).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return notes, errs
}

func (upld *UploadRepositoryImpl) Note(id uint) (*models.Upload,[]error){
    note := models.Upload{}
    	errs := upld.conn.First(&note, id).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return &note, errs
}

func (upld *UploadRepositoryImpl) UpdateNote(update *models.Upload) (*models.Upload,[]error){
    note := update
    	errs := upld.conn.Save(note).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}
func (upld *UploadRepositoryImpl) DeleteNote(id uint) (*models.Upload,[]error){
    note, errs := upld.Note(id)
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	errs = upld.conn.Delete(note, id).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}


func (upld *UploadRepositoryImpl) NotesByUploader(username,status string) ([]models.Upload,[]error){
        note := []models.Upload{}
    	errs := upld.conn.Find(&note, "file_uploader=? and file_status=?", username,status).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}

func (upld *UploadRepositoryImpl) NotesByCourseName(course,status string) ([]models.Upload,[]error){
        note := []models.Upload{}
    	errs := upld.conn.Find(&note, "file_uploaded_to=? and file_status=?", course, status).GetErrors()
    	if len(errs) > 0 {
    		return nil, errs
    	}
    	return note, errs
}
func (upld *UploadRepositoryImpl) ApproveNote(id uint) []error{

    	errs := upld.conn.Exec("UPDATE uploads SET file_status='approved' WHERE id=?",id).GetErrors()
    	if len(errs) > 0 {
    		return errs
    	}
    	return nil
}
