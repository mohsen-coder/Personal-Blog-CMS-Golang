package domain

type MimType string

const (
	FileNone MimType = "none"
	FileJPG  MimType = "image/jpeg"
	FilePNG  MimType = "image/png"
	FileGIF  MimType = "image/gif"
	FileZIP  MimType = "application/zip"
	FileRAR  MimType = "application/vnd.rar"
	FilePDF  MimType = "application/pdf"
	FileText MimType = "text/plain"
)
