package probe

import (
	"encoding/base64"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

type Notification struct {
	Date    time.Time
	Message string
}

const NbNotificationMax = 100

func AddNotification(notifications *[]*Notification, date time.Time, message string) {
	notification := &Notification{
		Date:    date,
		Message: message,
	}
	*notifications = append(*notifications, notification)

	if len(*notifications) > NbNotificationMax {
		*notifications = (*notifications)[1:]
	}
}

func DownloadNotificationsCSV(notifications []*Notification, loadStage *load.Stage, initLoadStage func()) {
	var csvContent string
	csvContent += "Date,Message\n"
	for _, notification := range notifications {
		// Escape quotes in message
		escapedMessage := strings.ReplaceAll(notification.Message, "\"", "\"\"")
		csvContent += fmt.Sprintf("\"%s\",\"%s\"\n", notification.Date.Format(time.StampMicro), escapedMessage)
	}

	loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	fileToDownload.Name = "notifications.csv"
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(csvContent))

	loadStage.StageBranch(fileToDownload)
	loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	initLoadStage()
}

func ExportStageExcel(
	excelBytes []byte,
	err error,
	fileName string,
	defaultPkgName string,
	stageName string,
	loadStage *load.Stage,
	initLoadStage func(),
	addNotification func(time.Time, string),
	commitNotificationTable func(),
) {
	if err != nil {
		addNotification(time.Now(), "Error serializing stage: "+err.Error())
		commitNotificationTable()
		return
	}

	loadStage.Reset()

	fileToDownload := new(load.FileToDownload)

	if fileName == "" {
		fileName = defaultPkgName + "-" + stageName + ".go"
	}

	prefixRegex := regexp.MustCompile(`^\d{8} \d{4} `)
	cleanFileName := prefixRegex.ReplaceAllString(fileName, "")
	cleanFileName = strings.TrimSuffix(cleanFileName, ".go") + ".xlsx"

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(excelBytes)

	loadStage.StageBranch(fileToDownload)
	loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	initLoadStage()
}

func ExportStage(
	stageString string,
	err error,
	fileName string,
	defaultPkgName string,
	stageName string,
	loadStage *load.Stage,
	initLoadStage func(),
	addNotification func(time.Time, string),
	commitNotificationTable func(),
) {
	if err != nil {
		addNotification(time.Now(), "Error serializing stage: "+err.Error())
		commitNotificationTable()
		return
	}

	loadStage.Reset()

	fileToDownload := new(load.FileToDownload)

	if fileName == "" {
		fileName = defaultPkgName + "-" + stageName + ".go"
	}

	prefixRegex := regexp.MustCompile(`^\d{8} \d{4} `)
	cleanFileName := prefixRegex.ReplaceAllString(fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))

	loadStage.StageBranch(fileToDownload)
	loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	initLoadStage()
}

func InitLoadStage(loadStage *load.Stage, fileUploadProxy load.FileToUploadProxy) {
	loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name:              "Name of file",
		FileToUploadProxy: fileUploadProxy,
	}

	loadStage.StageBranch(fileToUpload)

	message := &load.Message{
		Name: "Drop your stage.go file here or ",
	}

	message.Stage(loadStage)

	loadStage.Commit()
}

func ParseUploadedStage(
	uploadedFile *load.FileToUpload,
	parseAst func(inFile *ast.File, fset *token.FileSet) error,
) (string, error) {
	fileName := uploadedFile.GetName()

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fileName, fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", decodedBytes, parser.ParseComments)
	if errParser != nil {
		return fileName, fmt.Errorf("Unable to parse: %w", errParser)
	}

	errParse := parseAst(inFile, fset)
	if errParse != nil {
		return fileName, errParse
	}

	return fileName, nil
}
