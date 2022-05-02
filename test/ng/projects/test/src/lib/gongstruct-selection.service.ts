import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class GongstructSelectionService {

    // Observable string sources
    private gongstructSelection = new Subject<string>();

    // Observable string streams
    gongtructSelected$ = this.gongstructSelection.asObservable();

    // Service message commands
    gongstructSelected(gongstructName: string) {
        this.gongstructSelection.next(gongstructName);
    }

}
