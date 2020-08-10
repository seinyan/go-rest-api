package models

const (
	FileImage     = 1
	FileFile      = 2
	FileAudio     = 2
	FileFilePath  = "files/"
	FileImagePath = "images/"
	FileAudioPath = "audio/"
)

type File struct {
	BaseModel
	Type uint8  `gorm:"type:int;" json:"type"`
	Name string `gorm:"type:string;" json:"name"`
	Path string `gorm:"type:string;" json:"path"`
	Url  string `json:"url"`
	File string `gorm:"-" binding:"required" json:"file"`
}

