package flutter

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const main_dart_template = `import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    const textWidget = Text(
      'Hello, How are you?',
      textAlign: TextAlign.center,
      overflow: TextOverflow.ellipsis,
      style: TextStyle(fontWeight: FontWeight.bold),
      textDirection: TextDirection.ltr,
    );

    return textWidget;
  }
}
`

func GenerateDefaultMainFile(
	modelPkg *models.ModelPkg,
	mainFilePath string) {
	{
		codeGo := main_dart_template

		codeGo = models.Replace4(codeGo,
			"{{PkgName}}", modelPkg.Name,
			"{{TitlePkgName}}", strings.Title(modelPkg.Name),
			"{{pkgname}}", strings.ToLower(modelPkg.Name),
			"{{PkgPathRoot}}", strings.ReplaceAll(modelPkg.PkgPath, "/models", ""))
		codeGo = strings.ReplaceAll(codeGo, "{{PkgPathAboveRoot}}",
			strings.ReplaceAll(modelPkg.PkgPath, "/go/models", ""))

		file, err := os.Create(mainFilePath)
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeGo)
	}
}
