import { createActionGroup, emptyProps, props } from '@ngrx/store';

export const LayoutActions = createActionGroup({
  source: 'Layout',
  events: {
    'Toggle Sidebar': emptyProps(),
    'Set Sidebar': props<{ open: boolean }>(),
  },
});
