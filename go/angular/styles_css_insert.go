package angular

const StylesCssInsert = `
@import 'material-symbols/outlined.css';

@import '@angular/material/prebuilt-themes/azure-blue.css';

html, body { height: 100%; }
body { margin: 0; font-family: Roboto, "Helvetica Neue", sans-serif; }

/* Fallback for M3 typography when Roboto is not installed offline */
:root {
  --mat-sys-body-large-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-body-medium-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-body-small-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-display-large-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-display-medium-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-display-small-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-headline-large-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-headline-medium-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-headline-small-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-label-large-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-label-medium-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-label-small-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-title-large-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-title-medium-font: Roboto, "Helvetica Neue", sans-serif;
  --mat-sys-title-small-font: Roboto, "Helvetica Neue", sans-serif;
}

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
