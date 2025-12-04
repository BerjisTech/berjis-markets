import { ApplicationConfig, importProvidersFrom, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideHttpClient, withFetch, withInterceptors } from '@angular/common/http';
import { provideRouter } from '@angular/router';
import { provideEffects } from '@ngrx/effects';
import { provideStore } from '@ngrx/store';
import { provideRouterStore } from '@ngrx/router-store';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import { MatSnackBarModule } from '@angular/material/snack-bar';

import { routes } from './app.routes';
import { environment } from '../environments/environment';
import { authReducer } from './state/auth/auth.reducer';
import { layoutReducer } from './state/layout/layout.reducer';
import { AuthEffects } from './state/auth/auth.effects';
import { authTokenInterceptor } from './core/interceptors/auth-token.interceptor';
import { errorInterceptor } from './core/interceptors/error.interceptor';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideAnimations(),
    provideHttpClient(withFetch(), withInterceptors([authTokenInterceptor, errorInterceptor])),
    provideStore({ auth: authReducer, layout: layoutReducer }),
    provideRouterStore(),
    provideEffects([AuthEffects]),
    provideStoreDevtools({ maxAge: 25, logOnly: environment.production }),
    importProvidersFrom(MatSnackBarModule),
  ],
};
