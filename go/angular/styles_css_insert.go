package angular

const StylesCssInsert = `
@import 'material-symbols/outlined.css';

@import '@angular/material/prebuilt-themes/indigo-pink.css';

/* Force the material-icons class (used by default in mat-icon) to use the Material Symbols font */
.material-icons {
    font-family: 'Material Symbols Outlined';
    font-weight: normal;
    font-style: normal;
    font-size: 24px;
    line-height: 1;
    letter-spacing: normal;
    text-transform: none;
    display: inline-block;
    white-space: nowrap;
    word-wrap: normal;
    direction: ltr;
    -webkit-font-feature-settings: 'liga';
    -webkit-font-smoothing: antialiased;
}
`
