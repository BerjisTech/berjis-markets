import { NgFor } from '@angular/common';
import { ChangeDetectionStrategy, Component, computed, signal } from '@angular/core';
import { MatCardModule } from '@angular/material/card';

import { ValueProp } from '../../shared/models/value-prop.model';
import { MarketListComponent } from './components/market-list/market-list.component';

@Component({
  selector: 'app-market-shell',
  standalone: true,
  imports: [NgFor, MatCardModule, MarketListComponent],
  templateUrl: './market-shell.component.html',
  styleUrl: './market-shell.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MarketShellComponent {
  private readonly highlightsSignal = signal<ValueProp[]>([
    {
      title: 'Liquidity uptime',
      description: 'Targeting <5¢ spreads on top markets with automated maker incentives.',
    },
    {
      title: 'Compliance automation',
      description: 'KYC/AML checkpoints modeled directly after the product checklist.',
    },
    {
      title: 'Realtime telemetry',
      description: 'WebSocket fanout budget of 50ms to every connected trader.',
    },
  ]);

  readonly highlights = computed(() => this.highlightsSignal());
}
