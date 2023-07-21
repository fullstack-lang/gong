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
            "name": "Launch firefox localhost:4200",
            "type": "firefox",
            "request": "launch",
            "reAttach": true,
            "url": "http://localhost:4200/",
            "webRoot": "${workspaceFolder}/ng",
            "pathMappings": [
                {
                    "url": "webpack:///vendor",
                    "path": "${workspaceFolder}/vendor"
                },
                {
                    "url": "webpack:///projects",
                    "path": "${workspaceFolder}/ng/projects"
                }
            ]
        },
        {
            "name": "Launch firefox localhost:8080",
            "type": "firefox",
            "request": "launch",
            "reAttach": true,
            "url": "http://localhost:8080/",
            "webRoot": "${workspaceFolder}/ng",
            "pathMappings": [
                {
                    "url": "webpack:///vendor",
                    "path": "${workspaceFolder}/vendor"
                },
                {
                    "url": "webpack:///projects",
                    "path": "${workspaceFolder}/ng/projects"
                }
            ]
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
