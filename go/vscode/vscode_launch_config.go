package vscode

const VsCodeLaunchConfig = `{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Edge",
            "request": "launch",
            "type": "msedge",
            "url": "http://localhost:4200",
            "webRoot": "${workspaceFolder}/ng"
        },
        {
            "name": "Launch Chrome",
            "request": "launch",
            "type": "chrome",
            "url": "http://localhost:4200",
            "webRoot": "${workspaceFolder}/ng"
        },
        {
            "name": "go server -unmarshallFromCode=stage.go -marshallOnCommit=stage",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/go/cmd/{{pkgname}}",
            "args": [
                "-unmarshallFromCode=stage.go",
                "-marshallOnCommit=stage"
            ]
        },
    ]
}
`
