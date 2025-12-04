import { NgFor } from '@angular/common';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ValueProp } from './shared/models/value-prop.model';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, NgFor],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent {
  readonly valueProps: readonly ValueProp[] = [
    {
      title: 'Real-time trading',
      description: 'WebSocket-native Angular shell designed to stream order books under 50 ms.'
    },
    {
      title: 'Security and compliance',
      description: 'Built for KYC, AML, and audit-first flows with strict linting plus testing gates.'
    },
    {
      title: 'Scalable architecture',
      description: 'NgRx-ready modules, lazy routes, and Dockerized builds tuned for rapid iteration.'
    }
  ];

  trackByTitle(_: number, prop: ValueProp): string {
    return prop.title;
  }
}
