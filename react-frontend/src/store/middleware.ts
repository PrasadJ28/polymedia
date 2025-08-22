import { createActionCreatorInvariantMiddleware, Middleware } from "@reduxjs/toolkit";

export const appMiddleware = (getDefaultMiddleware: any) => {
    const isDev = "development" === "development";

    const isActionCreator = (
        action: unknown,
    ): action is Function & { type: unknown } =>
    typeof action === 'function' && 'type' in action

    const actionCreatorMiddleware = createActionCreatorInvariantMiddleware({
        isActionCreator,
    })

  const base = getDefaultMiddleware({
    immutableCheck: isDev,
    serializableCheck: isDev,
    thunk: true,
    isActionCreator,
  });

  return isDev ? base.concat(actionCreatorMiddleware) : base;
};
