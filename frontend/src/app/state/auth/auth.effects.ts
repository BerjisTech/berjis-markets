import { HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Actions, ROOT_EFFECTS_INIT, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, map, of, tap } from 'rxjs';

import { ApiHttpService } from '../../core/services/api-http.service';
import { AuthTokenService } from '../../core/services/auth-token.service';
import { AuthActions } from './auth.actions';

interface LoginResponse {
  token: string;
}

@Injectable()
export class AuthEffects {
  login$ = createEffect(() =>
    this.actions$.pipe(
      ofType(AuthActions.loginRequested),
      exhaustMap(({ email, password }) =>
        this.api.post<LoginResponse>('/auth/login', { email, password }).pipe(
          map((res) => AuthActions.loginSucceeded({ token: res.token })),
          catchError((error: HttpErrorResponse) =>
            of(AuthActions.loginFailed({ error: error.error?.message ?? 'Unable to login right now.' }))
          )
        )
      )
    )
  );

  hydrateToken$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ROOT_EFFECTS_INIT),
      map(() => AuthActions.setToken({ token: this.tokens.snapshot }))
    )
  );

  persistToken$ = createEffect(
    () =>
      this.actions$.pipe(
        ofType(AuthActions.loginSucceeded, AuthActions.setToken, AuthActions.logout),
        tap((action) => {
          if ('token' in action) {
            this.tokens.set(action.token ?? null);
          } else {
            this.tokens.set(null);
          }
        })
      ),
    { dispatch: false }
  );

  constructor(
    private readonly actions$: Actions,
    private readonly api: ApiHttpService,
    private readonly tokens: AuthTokenService
  ) {}
}
