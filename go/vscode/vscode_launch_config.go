package vscode

const VsCodeLaunchConfig = `{
    "version": "0.2.0",
    "configurations": [
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
            "name": "go server w/o diagrams",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/go/cmd/{{pkgname}}",
            "args": [
                "-diagrams=false",
                "--unmarshall=stage",
                "-marshallOnCommit=stage"
            ]
        },
        {
            "name": "go server with diagrams",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/go/cmd/{{pkgname}}",
            "args": [
                "--unmarshall=stage",
                "-marshallOnCommit=stage"
            ]
        },
    ]
}
`
