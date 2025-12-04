import { createReducer, on } from '@ngrx/store';

import { LayoutActions } from './layout.actions';

export interface LayoutState {
  sidebarOpen: boolean;
}

const initialState: LayoutState = {
  sidebarOpen: false,
};

export const layoutReducer = createReducer(
  initialState,
  on(LayoutActions.toggleSidebar, (state) => ({ ...state, sidebarOpen: !state.sidebarOpen })),
  on(LayoutActions.setSidebar, (state, { open }) => ({ ...state, sidebarOpen: open }))
);
