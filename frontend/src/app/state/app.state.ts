import { AuthState } from './auth/auth.reducer';
import { LayoutState } from './layout/layout.reducer';

export interface AppState {
  auth: AuthState;
  layout: LayoutState;
}
