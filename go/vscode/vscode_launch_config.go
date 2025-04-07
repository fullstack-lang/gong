package vscode

const VsCodeLaunchConfig = `{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Edge",
            "request": "launch",
            "type": "msedge",
            "url": "http://localhost:4200",
            "webRoot": "${workspaceFolder}/{{NgWorkspaceName}}"
        },
        {
            "name": "Launch Chrome",
            "request": "launch",
            "type": "chrome",
            "url": "http://localhost:4200",
            "webRoot": "${workspaceFolder}/{{NgWorkspaceName}}"
        },
        {
            "name": "go server -unmarshallFromCode=data/stage.go -marshallOnCommit=data/stage",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/go/cmd/{{pkgname}}",
            "args": [
                "-unmarshallFromCode=data/stage.go",
                "-marshallOnCommit=data/stage"
            ]
        },
    ]
}
`
