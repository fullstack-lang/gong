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
        "../vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/public-api.ts",
        "../../vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/public-api.ts"
      ],
      "gongdoc": [
        "../vendor/github.com/fullstack-lang/gongdoc/ng/projects/gongdoc/src/public-api.ts"
      ],
      "gongdocspecific": [
        "../vendor/github.com/fullstack-lang/gongdoc/ng/projects/gongdocspecific/src/public-api.ts"
      ],
      "gongsvg": [
        "../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvg/src/public-api.ts"
      ],
      "gongsvgdatamodel": [
        "../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvgdatamodel/src/public-api.ts"
      ],
      "gongsvgspecific": [
        "../vendor/github.com/fullstack-lang/gongsvg/ng/projects/gongsvgspecific/src/public-api.ts"
      ],
      "gongtree": [
        "../vendor/github.com/fullstack-lang/gongtree/ng/projects/gongtree/src/public-api.ts"
      ],
      "gongtreedatamodel": [
        "../vendor/github.com/fullstack-lang/gongtree/ng/projects/gongtreedatamodel/src/public-api.ts"
      ],
      "gongtreespecific": [
        "../vendor/github.com/fullstack-lang/gongtree/ng/projects/gongtreespecific/src/public-api.ts"
      ],
      "gongrouter": [
        "../vendor/github.com/fullstack-lang/gongrouter/ng/projects/gongrouter/src/public-api.ts"
      ],
      "gongrouterspecific": [
        "../vendor/github.com/fullstack-lang/gongrouter/ng/projects/gongrouterspecific/src/public-api.ts"
      ],
      "gongrouterdatamodel": [
        "../vendor/github.com/fullstack-lang/gongrouter/ng/projects/gongrouterdatamodel/src/public-api.ts"
      ],
`
