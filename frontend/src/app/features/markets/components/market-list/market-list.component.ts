import { AsyncPipe, DatePipe, NgFor, PercentPipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { map, startWith } from 'rxjs';

import { MarketSummary } from '../../../../shared/models/market-summary.model';

@Component({
  selector: 'app-market-list',
  standalone: true,
  imports: [ReactiveFormsModule, MatFormFieldModule, MatInputModule, MatCardModule, MatChipsModule, NgFor, DatePipe, PercentPipe, AsyncPipe],
  templateUrl: './market-list.component.html',
  styleUrl: './market-list.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MarketListComponent {
  private readonly markets = signal<MarketSummary[]>([
    { id: '1', title: 'Will CPI fall below 2.5% by Q3?', category: 'Economics', probability: 0.42, closeDate: '2025-07-01', status: 'trading' },
    { id: '2', title: 'Will a major hurricane make US landfall?', category: 'Weather', probability: 0.31, closeDate: '2025-10-15', status: 'scheduled' },
    { id: '3', title: 'Will Fed cut rates by 50 bps in 2025?', category: 'Economics', probability: 0.58, closeDate: '2025-03-20', status: 'trading' },
  ]);

  readonly filterControl = new FormControl('', { nonNullable: true });

  readonly filtered$ = this.filterControl.valueChanges.pipe(
    startWith(''),
    map((query) => this.filterMarkets(query ?? ''))
  );

  private filterMarkets(query: string): MarketSummary[] {
    const needle = query.toLowerCase();
    return this.markets().filter((market) =>
      [market.title, market.category].some((field) => field.toLowerCase().includes(needle))
    );
  }
}
