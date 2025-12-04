import { createReducer, on } from '@ngrx/store';

import { AuthActions } from './auth.actions';

export interface AuthState {
  token: string | null;
  status: 'idle' | 'loading' | 'authenticated' | 'error';
  error?: string | null;
}

const initialState: AuthState = {
  token: null,
  status: 'idle',
  error: null,
};

export const authReducer = createReducer(
  initialState,
  on(AuthActions.loginRequested, (state) => ({ ...state, status: 'loading', error: null })),
  on(AuthActions.loginSucceeded, (state, { token }) => ({ ...state, token, status: 'authenticated', error: null })),
  on(AuthActions.loginFailed, (state, { error }) => ({ ...state, status: 'error', error })),
  on(AuthActions.logout, () => initialState),
  on(AuthActions.setToken, (state, { token }) => ({
    ...state,
    token,
    status: token ? 'authenticated' : 'idle',
  }))
);
