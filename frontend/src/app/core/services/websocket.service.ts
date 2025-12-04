import { Injectable } from '@angular/core';
import { Observable, retry, shareReplay } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class WebSocketService {
  connect<T>(path: string): Observable<T> {
    const url = `${environment.wsBaseUrl}${path}`;
    const socket$: WebSocketSubject<T> = webSocket({ url });
    return socket$.pipe(retry({ count: 5, delay: 2000 }), shareReplay({ bufferSize: 1, refCount: true }));
  }
}
