import { Injectable } from '@angular/core'
import { Observable } from 'rxjs'

import { DOCUMENT } from '@angular/common';
import { HttpParams } from '@angular/common/http';
import { Component, Inject, Input, OnInit } from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class WebSocketService {
    private socket: WebSocket | undefined

    public connect(stackPath: string): Observable<any> {

        let params = new HttpParams().set("GONG__StackPath", stackPath)
        let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongtable/go/v1/ws/stage'
        let paramString = params.toString()
        let url = `${basePath}?${paramString}`
        this.socket = new WebSocket(url)


        return new Observable(observer => {
            this.socket!.onmessage = event => {
                observer.next(event)
            }
            this.socket!.onerror = event => {
                observer.error(event)
            }
            this.socket!.onclose = event => {
                observer.complete()
            }

            return () => {
                this.socket!.close()
            }
        })

    }
}
