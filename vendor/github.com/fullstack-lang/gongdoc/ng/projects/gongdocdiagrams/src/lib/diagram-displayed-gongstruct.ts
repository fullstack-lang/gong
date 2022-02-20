import { Subject } from 'rxjs';

export class ClassdiagramContext {
    ClassdiagramID: number = 0
}

// observable of the content of the current diagram
export var ClassdiagramContextSubject = new Subject<ClassdiagramContext>()