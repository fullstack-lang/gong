// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number = 0
    Valid: boolean = true // exception to the golang convention that boolean default value is false
}
