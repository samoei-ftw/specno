import React, { forwardRef } from 'react';

const DropWrapper = forwardRef<HTMLDivElement, { children: React.ReactNode }>(
  ({ children }, ref) => {
    return <div ref={ref}>{children}</div>;
  }
);

export default DropWrapper;