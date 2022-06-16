package models

const NgSplitterTemplateHTML = `<div style="width: 100%; height: 100%; background: grey(16);">
    <as-split direction="horizontal">
        <as-split-area [size]="20">
            <as-split direction="vertical">
                <as-split-area>
                    <app-{{pkgname}}-sidebar></app-{{pkgname}}-sidebar>
                </as-split-area>
            </as-split>
        </as-split-area>
        <as-split-area [size]="80">
            <as-split direction="vertical">
                <as-split-area  style="overflow: hidden;">
                    <router-outlet name="{{PkgPathRootWithoutSlashes}}_table"></router-outlet>
                </as-split-area>
                <as-split-area [size]="20">
                    <as-split direction="horizontal">
                        <as-split-area>
                            <router-outlet name="{{PkgPathRootWithoutSlashes}}_presentation"></router-outlet>
                        </as-split-area>
                        <as-split-area>
                            <router-outlet name="{{PkgPathRootWithoutSlashes}}_editor"></router-outlet>
                        </as-split-area>
                    </as-split>
                </as-split-area>
            </as-split>
        </as-split-area>
    </as-split>
</div>`
