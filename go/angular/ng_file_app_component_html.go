package angular

const NgFileAppComponentHtml = `<as-split unit="pixel" direction="vertical">
    <as-split-area [size]="400">
        <lib-{{pkgname}}-specific></lib-{{pkgname}}-specific>
    </as-split-area>
    <as-split-area [size]="$any('*')">
        <lib-split-specific Name="{{pkgname}}-probe"></lib-split-specific>
    </as-split-area>
</as-split>
`
