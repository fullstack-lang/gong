package vscode

const VsCodeTasksConfig = `{
	"version": "2.0.0",
	"tasks": [
	  {
		"label": "02 - run {{pkgname}} -unmarshallFromCode=data/stage.go -marshallOnCommit=data/stage",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "run",
		  "main.go",
		  "-unmarshallFromCode",
		  "data/stage.go",
		  "-marshallOnCommit",
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
		"label": "01 - gongc {{pkgname}}",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": []
	  },
	  {
		"label": "01 - gongc {{pkgname}} -skipNg",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": [
		  "-skipNg"
		]
	  },
	  {
		"label": "01 - gongc {{pkgname}} -skipGoModCommands",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": [
			"-skipGoModCommands"
		]
	  },
	  {
		"label": "01 - gongc {{pkgname}} -skipNg -skipGoModCommands",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": [
		  "-skipNg",
		  "-skipGoModCommands"
		]
	  },
	  {
		"label": "00 - update gong dependency",
		"type": "shell",
		"command": "go",
		"args": [
		  "get",
		  "github.com/fullstack-lang/gong@main",
		],
		"group": "build"
	  },
	  {
		"label": "00 - go mod tidy",
		"type": "shell",
		"command": "go",
		"args": [
		  "mod",
		  "tidy"
		],
		"group": "build"
	  },
	  	  {
		"label": "00 - go mod vendor",
		"type": "shell",
		"group": "build",
		"command": "go",
		"args": [
		  "mod",
		  "vendor"
		],
	  },
	  {
		"label": "00 - go build",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "build"
		],
		"group": "build"
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
