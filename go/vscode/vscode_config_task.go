package vscode

const VsCodeTasksConfig = `{
	"version": "2.0.0",
	"tasks": [
	  {
		"label": "02 - run {{pkgname}} --unmarshallFromCode=data/stage.go --marshallOnCommit=data/stage",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "run",
		  "main.go",
		  "--unmarshallFromCode",
		  "data/stage.go",
		  "--marshallOnCommit",
		  "data/stage.go"
		],
		"group": "build"
	  },
	  {
		"label": "03 - lsof -i tcp:8080 ",
		"type": "shell",
		"command": "lsof -i tcp:8080 ",
		"group": "build"
	  },
	  {
		"label": "03 - ng build",
		"type": "shell",
		"command": "cd {{NgWorkspaceName}}; ng build",
		"group": "build",
		"problemMatcher": []
	  },
	  {
		"label": "04 - ng serve --host=127.0.0.1",
		"type": "shell",
		"command": "cd {{NgWorkspaceName}}; ng serve --host=127.0.0.1",
		"group": "build",
		"problemMatcher": []
	  },
	  {
		"label": "01 - gong generate",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gong",
		"group": "build",
		"args": [
			"generate",
		]
	  },
   	  {
		"label": "00 - go install",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "install"
		],
		"group": "build"
	  },
	]
  }`
