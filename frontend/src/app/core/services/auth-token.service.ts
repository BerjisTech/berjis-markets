import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AuthTokenService {
  private readonly storageKey = 'markets-token';
  private readonly tokenSubject = new BehaviorSubject<string | null>(this.readToken());
  readonly token$ = this.tokenSubject.asObservable();

  get snapshot(): string | null {
    return this.tokenSubject.value;
  }

  set(token: string | null): void {
    if (typeof window !== 'undefined') {
      if (token) {
        window.localStorage.setItem(this.storageKey, token);
      } else {
        window.localStorage.removeItem(this.storageKey);
      }
    }
    this.tokenSubject.next(token);
  }

  private readToken(): string | null {
    return typeof window !== 'undefined' ? window.localStorage.getItem(this.storageKey) : null;
  }
}
