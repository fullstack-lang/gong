package angular

const TsConfigInsertForPaths = `      // with gong, some angular standalone component are imported from other stacks
      "@vendored_components/*": [
      	"../vendor/*",
      ],`
