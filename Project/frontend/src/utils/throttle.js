export function throttle(callback, interval) {
    let enableCall = true;
  
    return function(...args) {
      if (!enableCall) return;
  
      enableCall = false;
      console.log(args)
      callback.apply(this, args);
      setTimeout(() => enableCall = true, interval);
    }
  }