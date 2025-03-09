package angular

const TsConfigInsertForPaths = `      // with gong, some angular standalone component are imported from other stacks
      "@vendored_components/*": [
        "../vendor/*", // path for stack at the root of the module
        "../../vendor/*", // path for stack within 1 level of the module
        "../../../vendor/*", // path for stack within 2 levels of the module
        "../../../../../../*", // path with gong internal lib and test libraries
      ],`
