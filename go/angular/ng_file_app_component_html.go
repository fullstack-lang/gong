package angular

const NgFileAppComponentHtml = `<as-split unit="pixel" direction="vertical">
    <as-split-area [size]="400">
        <lib-{{pkgname}}-specific GONG__StackPath="{{pkgname}}">
        </lib-{{pkgname}}-specific>
    </as-split-area>
    <as-split-area [size]="$any('*')">
        <lib-split-specific [GONG__StackPath]="StackName"></lib-split-specific>
    </as-split-area>
</as-split>
`
