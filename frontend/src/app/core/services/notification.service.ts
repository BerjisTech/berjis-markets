import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({ providedIn: 'root' })
export class NotificationService {
  constructor(private readonly snackBar: MatSnackBar) {}

  success(message: string): void {
    this.open(message, 'success');
  }

  error(message: string): void {
    this.open(message, 'error');
  }

  private open(message: string, panelClass: string): void {
    this.snackBar.open(message, 'Close', {
      panelClass,
      duration: 4000,
      horizontalPosition: 'right',
      verticalPosition: 'top'
    });
  }
}
