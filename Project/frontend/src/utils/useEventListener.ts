import React, { useRef, useEffect, EventHandler } from "react";

export function useEventListener(eventName: any, handler: any, element = window){
    const savedHandler = useRef<any>(null);
    useEffect(() => {
      savedHandler.current = handler;
    }, [handler]);
  
    useEffect(
      () => {
        const isSupported = element && element.addEventListener;
        if (!isSupported) return;
        const eventListener = (event: Event) => savedHandler.current(event);
        element.addEventListener(eventName, eventListener);
        return () => {
          element.removeEventListener(eventName, eventListener);
        };
      },
      [eventName, element]
    );
  };