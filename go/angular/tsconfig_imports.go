package angular

const TsConfigInsertForPaths = `
      // 
      // https://angular.io/guide/creating-libraries#use-typescript-path-mapping-for-peer-dependencies
      //
      // Angular libraries should list any @angular/* 
      // dependencies the library depends on as peer dependencies. This ensures that when modules ask for Angular, 
      // they all get the exact same module. If a library lists @angular/core in dependencies instead of peerDependencies, 
      // it might get a different Angular module instead, which would cause your application to break.
      "@angular/*": [
        "./node_modules/@angular/*"
      ],
      "rxjs/operators": [
        "./node_modules/rxjs/operators"
      ],
      "rxjs": [
        "./node_modules/rxjs"
      ],
      "@angular-material-components/datetime-picker": [
        "./node_modules/@angular-material-components/datetime-picker"
      ],
      "angular-split": [
        "./node_modules/angular-split"
      ],
      "tslib": [
        "./node_modules/tslib"
      ],
      "gong": [
      "../vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/public-api.ts"
      ],
      "gongdoc": [
      "../vendor/github.com/fullstack-lang/gongdoc/ng/projects/gongdoc/src/public-api.ts"
      ],
      "gongdocdiagrams": [
      "../vendor/github.com/fullstack-lang/gongdoc/ng/projects/gongdocdiagrams/src/public-api.ts"
      ],
      "gongsvg": [
        "../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvg/src/public-api.ts"
      ],
      "gongsvgspecific": [
        "../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvgspecific/src/public-api.ts"
      ],
`
