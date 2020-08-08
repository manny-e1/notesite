package models

import "time"

type Upload struct{
    ID uint
    Title string `gorm:"type:varchar(30);not null"`
    File_Description string `gorm:"type:text;not null"`
    File_Type string `gorm:"type:varchar(255);not null"`
    File_Uploader string `gorm:"type:varchar(255);not null"`
    File_Uploader_ID uint `gorm:"type:integer;not null"`
    File_Uploaded_On time.Time
    File_Uploaded_To string `gorm:"type:varchar(30);not null"`
    File_Path string `gorm:"type:varchar(255);not null"`
    File_Status string `gorm:"type:varchar(30);not null"`
}