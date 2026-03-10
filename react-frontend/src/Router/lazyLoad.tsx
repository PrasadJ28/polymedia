import React, { Suspense, ComponentType } from 'react';

// 1. Define a type for the dynamic import function
// It must return a Promise that resolves to an object with a default export
type ImportComponent = () => Promise<{ default: ComponentType<any> }>;

export const lazyLoad = (importFunc: ImportComponent) => {
  const LazyComponent = React.lazy(importFunc);

  // 2. Return a component that forwards props
  return (props: any) => (
    <Suspense fallback={<div>Loading app...</div>}>
      <LazyComponent {...props} />
    </Suspense>
  );
};
