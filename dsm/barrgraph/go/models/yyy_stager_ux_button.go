// generated code (do not edit)
package models

import (
	"encoding/base64"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) GetPkgName() string {
	pkgPath := stager.stage.MetaPackageImportPath
	pkgName := ""
	parts := strings.Split(pkgPath, "/")
	if len(parts) >= 3 {
		pkgName = parts[len(parts)-3]
	}
	return pkgName
}

func (stager *Stager) button() {
	buttonStage := stager.buttonStage
	buttonStage.Reset()

	layout := new(button.Layout)

	group1 := new(button.Group)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Stop",
		Icon:  string(buttons.BUTTON_stop_circle),
		Label: "Stop",
		OnClick: func() {
			log.Println("Stop")
			os.Exit(0)
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:    "Web",
		Icon:    string(buttons.BUTTON_web),
		Label:   "Web",
		OnClick: stager.exportWebsite,
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Go",
		Icon:  string(buttons.BUTTON_file_download),
		Label: "Go",
		OnClick: func() {
			log.Println("Exporting stage as Go file")

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload)

			if stager.fileName == "" {
				pkgPath := stager.stage.MetaPackageImportPath
				pkgName := ""
				parts := strings.Split(pkgPath, "/")
				if len(parts) >= 3 {
					pkgName = parts[len(parts)-3]
				}
				stager.fileName = pkgName + "-" + stager.stage.GetName() + ".go"
			}

			prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

			stageString, err := stager.stage.MarshallToString(stager.stage.MetaPackageImportPath, "main")
			if err != nil {
				log.Println("Error serializing stage: " + err.Error())
				return
			}

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))

			load.StageBranch(stager.loadStage, fileToDownload)
			stager.loadStage.Commit()

			time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
			stager.load()
		},
	})

	group1.Buttons = append(group1.Buttons, &button.Button{
		Name:  "Export launcher HTML",
		Icon:  string(buttons.BUTTON_launch),
		Label: "Export launcher HTML",
		OnClick: func() {
			log.Println("Exporting launcher HTML")

			stager.loadStage.Reset()

			fileToDownload := new(load.FileToDownload)

			pkgName := stager.GetPkgName()

			if stager.fileName == "" {
				stager.fileName = pkgName + "-" + stager.stage.GetName() + ".go"
			}

			prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
			cleanFileName := prefixRegex.ReplaceAllString(stager.fileName, "")

			fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName + ".html"

			stageString, err := stager.stage.MarshallToString(stager.stage.MetaPackageImportPath, "main")
			if err != nil {
				log.Println("Error serializing stage: " + err.Error())
				return
			}

			// Escape the Go string for safe inclusion in a JavaScript template literal
			jsGoCode := strings.ReplaceAll(stageString, "\\", "\\\\")
			jsGoCode = strings.ReplaceAll(jsGoCode, "`", "\\`")
			jsGoCode = strings.ReplaceAll(jsGoCode, "${", "\\${")
			jsGoCode = strings.ReplaceAll(jsGoCode, "</script>", "<\\/script>")

			htmlString := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Launch ` + pkgName + ` App</title>
    <style>
        body { font-family: sans-serif; display: flex; justify-content: center; margin-top: 50px; }
        button { padding: 10px 20px; font-size: 16px; cursor: pointer; }
    </style>
</head>
<body>

    <button id="launchBtn">Open App and Process ` + cleanFileName + `</button>

    <script>
        const rawGoCode = ` + "`" + jsGoCode + "`" + `;
        
        // Base64 encode the UTF-8 string to pass it as fileData
        const utf8Bytes = new TextEncoder().encode(rawGoCode);
        let binary = '';
        for (let i = 0; i < utf8Bytes.length; i++) {
            binary += String.fromCharCode(utf8Bytes[i]);
        }
        const fileXContent = btoa(binary);
        
        const fileXName = "` + cleanFileName + `";
        const targetUrl = "https://fullstack-lang.github.io/gong/` + pkgName + `-app-portable.html";
        const targetOrigin = "https://fullstack-lang.github.io"; 

        document.getElementById('launchBtn').addEventListener('click', () => {
            const appWindow = window.open(targetUrl, '_blank');
            const payload = {
                action: 'PROCESS_INJECTED_FILE',
                fileName: fileXName,
                fileData: fileXContent
            };

            let ackReceived = false;
            window.addEventListener('message', (event) => {
                if (event.data && event.data.action === 'FILE_PROCESSED') {
                    ackReceived = true;
                    console.log("Application successfully received and processed the file!");
                }
            });

            const sendInterval = setInterval(() => {
                if (ackReceived || appWindow.closed) {
                    clearInterval(sendInterval);
                    return;
                }
                appWindow.postMessage(payload, targetOrigin);
                console.log("File payload sent to application (waiting for ACK)...");
            }, 1000);
        });
    </script>
</body>
</html>`

			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(htmlString))

			load.StageBranch(stager.loadStage, fileToDownload)
			stager.loadStage.Commit()

			time.Sleep(1 * time.Second)
			stager.load()
		},
	})

	button.StageBranch(buttonStage, layout)

	buttonStage.Commit()
}
