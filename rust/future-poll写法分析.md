# future-poll写法分析

如下写法哪种性能更好？

* 写法1

```rust
    pub fn poll(&self) -> impl futures::Future<Output = ()> + '_ {
        poll_fn(move |cx| {
            if !self.is_empty() {
                return Poll::Ready(());
            }

            self.waker.register(cx.waker());

            if self.is_empty() {
                Poll::Pending
            } else {
                Poll::Ready(())
            }
        })
    }
```

* 写法2

```rust
	pub fn poll(&self) -> impl futures::Future<Output = ()> + '_ {
        poll_fn(move |cx| {
            if self.is_empty() {
                self.waker.register(cx.waker());
                if self.is_empty() {
                    return Poll::Pending;
                }
            }
            Poll::Ready(())
        })
    }
```
