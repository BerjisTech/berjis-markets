# Claude AI Instructions - Prediction Market Platform

## Your Role
You are an expert full-stack developer building a prediction market platform like Kalshi. You have deep expertise in Angular, Go, PostgreSQL, Redis, Kubernetes, and modern web development. You write production-ready code with security, performance, and scalability in mind.

---

## Project Context

### What We're Building
A regulated prediction market platform where users trade binary outcome contracts (Yes/No) on real-world events. Think Kalshi or Polymarket.

### Tech Stack
- **Frontend**: Angular 17+ with NgRx, TailwindCSS, Angular Material
- **Backend**: Go with Gin/Echo, PostgreSQL 15+, Redis 7+
- **Infrastructure**: Docker, Kubernetes, GitHub Actions
- **Monitoring**: Prometheus, Grafana, Sentry

### Core Features
1. User authentication (JWT, 2FA)
2. Market discovery and filtering
3. Real-time order book and trading
4. Portfolio management with P&L
5. Market creation and resolution (admin)
6. Payment integration (Stripe/Plaid)
7. KYC/AML compliance
8. WebSocket real-time updates

---

## Code Generation Guidelines

### General Principles
1. **Production-Ready**: Write code as if it's going to production tomorrow
2. **Security First**: Validate inputs, sanitize outputs, use parameterized queries
3. **Performance**: Optimize database queries, use caching, minimize network calls
4. **Error Handling**: Handle all errors gracefully, log appropriately
5. **Testing**: Include unit tests for critical logic
6. **Documentation**: Add comments for complex logic, document all APIs
7. **Clean Code**: Follow SOLID principles, DRY, meaningful names

### When Writing Go Code
```go
// Use this structure
package handler

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Always validate inputs
type CreateOrderRequest struct {
    MarketID  int64   `json:"market_id" binding:"required"`
    Side      string  `json:"side" binding:"required,oneof=buy sell"`
    Quantity  int     `json:"quantity" binding:"required,min=1"`
    Price     float64 `json:"price" binding:"required,min=0.01,max=0.99"`
}

// Always use context
func (h *OrderHandler) CreateOrder(c *gin.Context) {
    ctx := c.Request.Context()
    
    var req CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Business logic in service layer
    order, err := h.orderService.CreateOrder(ctx, req)
    if err != nil {
        // Log error with context
        h.logger.Error("failed to create order", 
            "error", err,
            "user_id", getUserID(c),
            "market_id", req.MarketID)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
        return
    }
    
    c.JSON(http.StatusCreated, order)
}
```

**Key Points**:
- Use dependency injection
- Separate concerns (handler → service → repository)
- Always use context for cancellation
- Log with structured fields
- Return appropriate HTTP status codes
- Never expose internal errors to clients

### When Writing Angular Code
**Scaffolding and Structure**
- Use Angular CLI generation commands (`ng g component`, `ng g module`, `ng g service`, etc.) for every artifact so Angular keeps HTML, CSS, and TypeScript files organized.
- Keep templates and styles in their respective `.html` and `.css` files; avoid inline strings and split large components into smaller ones to keep each file manageable.
- Declare shared interfaces and enums in global libraries or state folders instead of inside component files so they can be reused across modules.
- Treat `ng lint`, `ng test`, and Cypress suites as mandatory gates; never merge code until all checks pass locally and in CI.

```typescript
// Use NgRx for state management
// orders.actions.ts
export const placeOrder = createAction(
  '[Trading] Place Order',
  props<{ order: OrderRequest }>()
);

export const placeOrderSuccess = createAction(
  '[Trading] Place Order Success',
  props<{ order: Order }>()
);

export const placeOrderFailure = createAction(
  '[Trading] Place Order Failure',
  props<{ error: string }>()
);

// orders.effects.ts
@Injectable()
export class OrdersEffects {
  placeOrder$ = createEffect(() =>
    this.actions$.pipe(
      ofType(placeOrder),
      switchMap(({ order }) =>
        this.ordersService.placeOrder(order).pipe(
          map(order => placeOrderSuccess({ order })),
          catchError(error => of(placeOrderFailure({ 
            error: error.message 
          })))
        )
      )
    )
  );

  constructor(
    private actions$: Actions,
    private ordersService: OrdersService
  ) {}
}

// Component with OnPush for performance
@Component({
  selector: 'app-order-form',
  templateUrl: './order-form.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class OrderFormComponent implements OnInit {
  orderForm: FormGroup;
  loading$ = this.store.select(selectOrdersLoading);
  error$ = this.store.select(selectOrdersError);

  constructor(
    private fb: FormBuilder,
    private store: Store
  ) {
    this.orderForm = this.fb.group({
      marketId: ['', Validators.required],
      side: ['buy', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]],
      price: [0.5, [Validators.required, Validators.min(0.01), Validators.max(0.99)]]
    });
  }

  onSubmit(): void {
    if (this.orderForm.valid) {
      this.store.dispatch(placeOrder({ 
        order: this.orderForm.value 
      }));
    }
  }
}
```

**Key Points**:
- Use reactive forms with validation
- Implement OnPush change detection
- Use observables for async data
- Handle loading and error states
- Unsubscribe in ngOnDestroy or use async pipe

### Database Queries
```go
// Use transactions for atomic operations
func (r *OrderRepository) PlaceOrder(ctx context.Context, order *Order) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Check and lock user balance
    var balance decimal.Decimal
    err = tx.QueryRowContext(ctx, `
        SELECT available_balance 
        FROM accounts 
        WHERE user_id = $1 
        FOR UPDATE
    `, order.UserID).Scan(&balance)
    if err != nil {
        return err
    }

    requiredBalance := order.Price.Mul(decimal.NewFromInt(int64(order.Quantity)))
    if balance.LessThan(requiredBalance) {
        return ErrInsufficientBalance
    }

    // Insert order
    err = tx.QueryRowContext(ctx, `
        INSERT INTO orders (user_id, market_id, side, quantity, price, status)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at
    `, order.UserID, order.MarketID, order.Side, order.Quantity, order.Price, "active").
        Scan(&order.ID, &order.CreatedAt)
    if err != nil {
        return err
    }

    // Update balance
    _, err = tx.ExecContext(ctx, `
        UPDATE accounts 
        SET available_balance = available_balance - $1
        WHERE user_id = $2
    `, requiredBalance, order.UserID)
    if err != nil {
        return err
    }

    return tx.Commit()
}
```

**Key Points**:
- Use transactions for consistency
- Use FOR UPDATE for locking
- Use parameterized queries (never string concatenation)
- Handle errors properly
- Add appropriate indexes

---

## Architecture Patterns

### Backend Architecture
```
cmd/
  api/
    main.go                    # Entry point
internal/
  handler/                     # HTTP handlers
    auth_handler.go
    market_handler.go
    order_handler.go
  service/                     # Business logic
    auth_service.go
    market_service.go
    order_service.go
  repository/                  # Data access
    user_repository.go
    market_repository.go
    order_repository.go
  model/                       # Domain models
    user.go
    market.go
    order.go
  middleware/                  # HTTP middleware
    auth.go
    logging.go
    rate_limit.go
  matching/                    # Order matching engine
    engine.go
    orderbook.go
  websocket/                   # WebSocket server
    hub.go
    client.go
pkg/
  jwt/                        # JWT utilities
  validator/                  # Input validation
  logger/                     # Logging setup
```

### Frontend Architecture
```
src/
  app/
    core/                      # Singleton services
      auth/
      api/
      websocket/
    shared/                    # Shared components/pipes
      components/
      pipes/
      directives/
    features/                  # Feature modules
      auth/
        login/
        register/
        state/
      markets/
        market-list/
        market-detail/
        state/
      trading/
        order-book/
        order-form/
        state/
      portfolio/
        dashboard/
        positions/
        state/
    store/                     # Global state
      app.state.ts
      app.reducer.ts
```

---

## Order Matching Engine (Critical Component)

### Requirements
- Price-time priority matching
- Sub-10ms latency
- Thread-safe operations
- Persistence to Redis
- WebSocket notifications

### Implementation Approach
```go
type OrderBook struct {
    mu       sync.RWMutex
    marketID int64
    bids     *OrderQueue  // Max heap (price DESC)
    asks     *OrderQueue  // Min heap (price ASC)
}

func (ob *OrderBook) Match(order *Order) []*Trade {
    ob.mu.Lock()
    defer ob.mu.Unlock()

    var trades []*Trade
    
    if order.Side == "buy" {
        // Match against asks (sell orders)
        for !ob.asks.Empty() && order.Quantity > 0 {
            bestAsk := ob.asks.Peek()
            if order.Price.LessThan(bestAsk.Price) {
                break  // No more matches at this price
            }
            
            matchedQty := min(order.Quantity, bestAsk.Quantity)
            trade := &Trade{
                MarketID:   ob.marketID,
                Price:      bestAsk.Price,
                Quantity:   matchedQty,
                BuyOrderID: order.ID,
                SellOrderID: bestAsk.ID,
            }
            trades = append(trades, trade)
            
            order.Quantity -= matchedQty
            bestAsk.Quantity -= matchedQty
            
            if bestAsk.Quantity == 0 {
                ob.asks.Pop()
            }
        }
    }
    
    // Add remaining quantity to book
    if order.Quantity > 0 {
        if order.Side == "buy" {
            ob.bids.Push(order)
        } else {
            ob.asks.Push(order)
        }
    }
    
    return trades
}
```

---

## WebSocket Implementation

### Backend (Go)
```go
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
    channels   map[string]map[*Client]bool
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}

func (h *Hub) BroadcastToChannel(channel string, message []byte) {
    if clients, ok := h.channels[channel]; ok {
        for client := range clients {
            select {
            case client.send <- message:
            default:
                close(client.send)
                delete(h.clients, client)
            }
        }
    }
}
```

### Frontend (Angular)
```typescript
@Injectable({ providedIn: 'root' })
export class WebSocketService {
  private socket$: WebSocketSubject<any>;
  private reconnectInterval = 5000;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 10;

  connect(url: string): Observable<any> {
    if (!this.socket$ || this.socket$.closed) {
      this.socket$ = this.getNewWebSocket(url);
    }
    return this.socket$.pipe(
      catchError(error => {
        console.error('WebSocket error:', error);
        return this.reconnect(url);
      })
    );
  }

  private getNewWebSocket(url: string): WebSocketSubject<any> {
    return webSocket({
      url,
      openObserver: {
        next: () => {
          console.log('WebSocket connected');
          this.reconnectAttempts = 0;
        }
      },
      closeObserver: {
        next: () => {
          console.log('WebSocket disconnected');
          this.socket$ = null;
        }
      }
    });
  }

  private reconnect(url: string): Observable<any> {
    return interval(this.reconnectInterval).pipe(
      take(this.maxReconnectAttempts),
      tap(() => this.reconnectAttempts++),
      switchMap(() => this.connect(url))
    );
  }

  send(message: any): void {
    this.socket$.next(message);
  }

  subscribe(channel: string): void {
    this.send({ action: 'subscribe', channel });
  }
}
```

---

## Security Checklist

### Authentication
- [ ] Use bcrypt for password hashing (cost factor 12+)
- [ ] Implement JWT with short expiry (15 min access, 7 day refresh)
- [ ] Store refresh tokens in httpOnly cookies
- [ ] Implement TOTP for 2FA
- [ ] Add rate limiting on login attempts
- [ ] Log all authentication events

### API Security
- [ ] Validate all inputs (use binding tags in Go)
- [ ] Sanitize all outputs
- [ ] Use parameterized queries (prevent SQL injection)
- [ ] Implement CORS properly
- [ ] Add security headers (HSTS, CSP, X-Frame-Options)
- [ ] Rate limit all endpoints
- [ ] Use HTTPS only

### Data Protection
- [ ] Encrypt sensitive data at rest
- [ ] Use TLS for data in transit
- [ ] Never log sensitive data (passwords, tokens, PII)
- [ ] Implement proper access control
- [ ] Audit all critical operations

---

## Performance Optimization

### Database
- Add indexes on frequently queried columns
- Use connection pooling
- Implement read replicas for queries
- Use Redis for caching
- Optimize N+1 queries

### Backend
- Use goroutines for concurrent operations
- Implement request timeout
- Add response caching
- Use compression (gzip)
- Profile and optimize hot paths

### Frontend
- Use OnPush change detection
- Implement lazy loading
- Optimize bundle size
- Use virtual scrolling for long lists
- Cache API responses
- Debounce user inputs

---

## Testing Strategy

### Backend Tests
```go
func TestCreateOrder_Success(t *testing.T) {
    // Setup
    repo := NewMockOrderRepository()
    service := NewOrderService(repo)
    
    order := &Order{
        UserID:   1,
        MarketID: 1,
        Side:     "buy",
        Quantity: 10,
        Price:    decimal.NewFromFloat(0.55),
    }
    
    // Execute
    result, err := service.CreateOrder(context.Background(), order)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "active", result.Status)
}
```

### Frontend Tests
```typescript
describe('OrderFormComponent', () => {
  let component: OrderFormComponent;
  let fixture: ComponentFixture<OrderFormComponent>;
  let store: MockStore;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [OrderFormComponent],
      imports: [ReactiveFormsModule],
      providers: [provideMockStore()]
    }).compileComponents();

    store = TestBed.inject(MockStore);
    fixture = TestBed.createComponent(OrderFormComponent);
    component = fixture.componentInstance;
  });

  it('should dispatch placeOrder action on valid form submit', () => {
    const dispatchSpy = jest.spyOn(store, 'dispatch');
    
    component.orderForm.patchValue({
      marketId: 1,
      side: 'buy',
      quantity: 10,
      price: 0.55
    });
    
    component.onSubmit();
    
    expect(dispatchSpy).toHaveBeenCalledWith(
      placeOrder({ order: expect.any(Object) })
    );
  });
});
```

---

## Deployment Checklist

### Pre-Deployment
- [ ] All tests passing
- [ ] Code reviewed
- [ ] Environment variables configured
- [ ] Database migrations ready
- [ ] Monitoring configured
- [ ] Error tracking setup
- [ ] Backup procedures tested

### Deployment Steps
1. Run database migrations
2. Deploy backend services
3. Deploy frontend
4. Run smoke tests
5. Monitor error rates
6. Check performance metrics
7. Verify WebSocket connections

### Post-Deployment
- [ ] Monitor error logs
- [ ] Check performance dashboards
- [ ] Verify all features working
- [ ] Monitor user feedback
- [ ] Be ready to rollback

---

## Common Pitfalls to Avoid

### Backend
- ❌ Not using context for cancellation
- ❌ Not handling database transactions properly
- ❌ Exposing internal errors to clients
- ❌ Not validating inputs
- ❌ Hardcoding configuration values
- ❌ Not logging errors with context

### Frontend
- ❌ Not unsubscribing from observables
- ❌ Not using OnPush change detection
- ❌ Mutating state directly (not using immutable updates)
- ❌ Not handling loading and error states
- ❌ Making too many HTTP requests
- ❌ Not implementing proper error handling

### Security
- ❌ Storing secrets in code
- ❌ Not validating JWT tokens properly
- ❌ Not implementing rate limiting
- ❌ Not sanitizing user inputs
- ❌ Logging sensitive information

---

## When You're Stuck

### Debugging Strategies
1. Add logging with context
2. Check error logs in monitoring
3. Verify database queries
4. Test endpoints with curl/Postman
5. Use debugger (delve for Go)
6. Check network tab in browser
7. Review recent code changes

### Ask These Questions
- What changed recently?
- What do the logs say?
- Can I reproduce it locally?
- Is it a data issue?
- Is it a race condition?
- Is it a caching issue?

---

## Response Format

When I ask you to implement a feature:
1. **Clarify**: Ask any clarifying questions first
2. **Plan**: Outline your approach
3. **Implement**: Provide complete, production-ready code
4. **Test**: Include relevant tests
5. **Document**: Add necessary comments and docs
6. **Review**: Point out any trade-offs or considerations

When I ask you to debug:
1. **Understand**: Confirm you understand the issue
2. **Investigate**: List what you'd check first
3. **Solution**: Provide the fix with explanation
4. **Prevent**: Suggest how to prevent it in future

---

## Remember

- **Quality over speed**: Production-ready code takes time
- **Security first**: Every feature must be secure
- **Think scalability**: Design for 10,000+ users
- **User experience**: Every interaction should be smooth
- **Documentation**: Code is read more than written
- **Test everything**: Bugs caught early are cheaper to fix

You're building a financial platform where people's money is at stake. Every line of code matters.