package vscode

const VsCodeTasksConfig = `{
	"version": "2.0.0",
	"tasks": [
	  {
		"label": "02 - run {{pkgname}}",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "run",
		  ".",
		  "-diagrams",
		],
		"group": "build"
	  },
	  {
		"label": "02 - run {{pkgname}} w/o diagrams",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/cmd/{{pkgname}}"
		},
		"command": "go",
		"args": [
		  "run",
		  ".",
		  "-diagrams=false"
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
		"command": "cd ng; ng build",
		"group": "build",
		"problemMatcher": []
	  },
	  {
		"label": "04 - ng serve",
		"type": "shell",
		"command": "cd ng; ng serve",
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
		"label": "01 - gongc {{pkgname}} -backendOnly",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": [
		  "-backendOnly"
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
		"label": "01 - gongc {{pkgname}} -backendOnly -skipGoModCommands",
		"type": "shell",
		"options": {
		  "cwd": "${workspaceFolder}/go/models"
		},
		"command": "gongc",
		"group": "build",
		"args": [
		  "-backendOnly",
		  " -skipGoModCommands"
		]
	  },
	  {
		"label": "00 - go mod vendor",
		"type": "shell",
		"group": "build",
		"command": "go mod vendor"
	  },
	  {
		"label": "00 - update gong dependencies",
		"type": "shell",
		"command": "go",
		"args": [
		  "get",
		  "-d",
		  "github.com/fullstack-lang/gong@HEAD",
		  "github.com/fullstack-lang/gongdoc@HEAD",
		],
		"group": "build"
	  },
	  {
		"label": "00 - go get -d github.com/fullstack-lang/gong@HEAD",
		"type": "shell",
		"command": "go",
		"args": [
		  "get",
		  "-d",
		  "github.com/fullstack-lang/gong@HEAD",
		],
		"group": "build"
	  },
	  {
		"label": "00 - go get -d github.com/fullstack-lang/gongdoc@HEAD",
		"type": "shell",
		"command": "go",
		"args": [
		  "get",
		  "-d",
		  "github.com/fullstack-lang/gongdoc@HEAD",
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
	]
  }`
