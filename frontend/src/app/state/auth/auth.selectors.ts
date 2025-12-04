import { createSelector } from '@ngrx/store';

import { AppState } from '../app.state';

export const selectAuthState = (state: AppState) => state.auth;

export const selectAuthToken = createSelector(selectAuthState, (auth) => auth.token);

export const selectIsAuthenticated = createSelector(selectAuthToken, (token) => !!token);

export const selectAuthStatus = createSelector(selectAuthState, (auth) => auth.status);

export const selectAuthError = createSelector(selectAuthState, (auth) => auth.error);
