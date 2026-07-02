import * as table from '../../../table/src/public-api'

export interface TableDialogData {
    Name: string
    TableName: string
    AssociationStorage: string
    initialSelection?: table.Row[] // Add this optional property
}
