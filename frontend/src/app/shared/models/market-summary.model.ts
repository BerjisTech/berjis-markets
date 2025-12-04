export interface MarketSummary {
  id: string;
  title: string;
  category: string;
  probability: number;
  closeDate: string;
  status: 'trading' | 'settled' | 'scheduled';
}
