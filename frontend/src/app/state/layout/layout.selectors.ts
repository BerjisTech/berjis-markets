import { createSelector } from '@ngrx/store';

import { AppState } from '../app.state';

const selectLayoutState = (state: AppState) => state.layout;

export const selectSidebarOpen = createSelector(selectLayoutState, (layout) => layout.sidebarOpen);
