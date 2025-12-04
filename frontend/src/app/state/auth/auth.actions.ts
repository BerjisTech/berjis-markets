import { createActionGroup, emptyProps, props } from '@ngrx/store';

export const AuthActions = createActionGroup({
  source: 'Auth',
  events: {
    'Login Requested': props<{ email: string; password: string }>(),
    'Login Succeeded': props<{ token: string }>(),
    'Login Failed': props<{ error: string }>(),
    'Logout': emptyProps(),
    'Set Token': props<{ token: string | null }>(),
  },
});
